package client

import (
	"bufio"
	"context"
	"fmt"
	"grpc_loader/pkg/api"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SendTimeoutSec = 5

func RunClient() {
	var filepath string
	myId := generateClientId()
	client, err := createClient()
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Pass me a filepath:")
		fmt.Scanln(&filepath)
		for {
			err := uploadFileToServer(client, myId, filepath)
			if err == nil {
				break
			}
		}
	}
}

func generateClientId() api.ClientId {
	return uuid.New().String()
}

func getChunkStartFrom(client api.UploaderClient, filename, modifiedAt, myId string, size int) (int, error) {
	chunksCount := setChunksCount(size)
	req := &api.UploadInitRequest{Filename: filename, ModifiedAt: modifiedAt, Size: int32(size), ChunksCount: int32(chunksCount), ClientId: myId}
	res, err := client.InitUpload(context.Background(), req)
	if err != nil {
		return 0, err
	}
	return int(res.ChunkStartFrom), nil
}

func setChunksCount(size int) int {
	answer := size / api.OneChunkSize
	if size%api.OneChunkSize != 0 {
		return answer + 1
	}
	return answer
}

func createClient() (api.UploaderClient, error) {
	conn, err := grpc.Dial(api.ServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := api.NewUploaderClient(conn)
	return c, nil
}

func uploadFileToServer(client api.UploaderClient, myId string, filepath string) error {
	file, chunksCount, chunkStartFrom := openFile(filepath, myId, client)
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	if chunkStartFrom == chunksCount-1 {
		log.Info("File is already on server")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), SendTimeoutSec*time.Second)
	defer cancel()
	stream, err := client.Upload(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = uploadChunks(stream, ctx, reader, fileInfo.Name(), chunkStartFrom, myId)
	if err != nil {
		return err
	}
	return nil
}

func uploadChunks(stream api.Uploader_UploadClient, ctx context.Context, reader *bufio.Reader, filename string, chunkStartFrom int, myId string) error {
	var err error
	buf := make([]byte, api.OneChunkSize)
	for {
		n, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				// if N is not zero, we have to send last partially-filled chunk
				if n == 0 {
					break
				}
			} else {
				log.Fatal(err)
			}
		}
		err = stream.SendMsg(&api.UploadRequest{Chunk: buf, ChunkId: int32(chunkStartFrom), ClientId: myId, Filename: filename})
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("%d chunk sent", chunkStartFrom+1)
		chunkStartFrom++
	}
	var res = &api.UploadResponse{}
	err = stream.RecvMsg(res)
	if err != nil {
		log.Error(err)
		return err
	}
	if res.Status == 1 {
		log.Info("Server response: file successfully uploaded")
		return nil
	}
	return fmt.Errorf("server responsed bad")
}

func openFile(filepath, myId string, client api.UploaderClient) (file *os.File, chunksCount int, chunkStartFrom int) {
	pathSlice := strings.Split(filepath, "/")
	filename := pathSlice[len(pathSlice)-1]
	log.Infof("%s got", filename)
	var err error
	file, err = os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	fileModifiedAt := fileInfo.ModTime().String()
	if err != nil {
		log.Fatal(err)
	}
	chunksCount = setChunksCount(int(fileSize))
	chunkStartFrom, err = getChunkStartFrom(client, filename, fileModifiedAt, myId, int(fileSize))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Seek(int64(chunkStartFrom)*api.OneChunkSize, 0)
	if err != nil {
		log.Fatal(err)
	}
	return
}
