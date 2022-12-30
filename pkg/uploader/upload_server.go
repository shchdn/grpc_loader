package uploader

import (
	"context"
	"fmt"
	"grpc_loader/pkg/api"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UploadServer struct {
	api.UploaderServer
	meta api.Metadata
}

type FileMeta struct {
	File         *os.File
	Filepath     string
	Filename     string
	UploadedPath string
	ClientId     string
	ChunksCount  int
}

func NewFileMeta(req *api.UploadRequest, serverMeta *api.Metadata) (FileMeta, error) {
	uploadedPath, err := getTargetDirectory()
	if err != nil {
		log.Error(err)
		return FileMeta{}, err
	}
	filename := req.Filename
	filepath := fmt.Sprintf("%sclient_%s/%s", uploadedPath, req.ClientId, filename)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Error(err)
		return FileMeta{}, status.Errorf(codes.Internal, "Server error")
	}
	clientId := req.ClientId
	chunksCount := serverMeta.GetChunksCount(clientId, req.Filename)
	return FileMeta{file, filepath, filename, uploadedPath, clientId, chunksCount}, nil
}

func NewUploadServer() *UploadServer {
	return &UploadServer{meta: api.NewMetadata()}
}

func (s *UploadServer) InitUpload(ctx context.Context, req *api.UploadInitRequest) (*api.UploadInitResponse, error) {
	log.Infof("[%s] Init upload", req.ClientId)
	userFileExists, err := checkUserFileExists(req)
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "Server error")
	}
	if userFileExists {
		requestedFilePath, err := getTargetFilepath(req.ClientId, req.Filename)
		if err != nil {
			return nil, err
		}
		// Compare meta from storage
		modifiedEqual := s.meta.CompareModified(req.ClientId, req.Filename, req.ModifiedAt)
		// Delete if not equal
		if !modifiedEqual {
			log.Infof("[%s] Files are not equal. Deleting file from server...", req.ClientId)
			err = deleteUserFile(requestedFilePath)
			if err != nil {
				return nil, err
			}
		} else {
			currentChunk := s.meta.GetCurrentChunk(req.ClientId, req.Filename)
			log.Infof("[%s] Files are equal. Current chunk is - %d", req.ClientId, currentChunk)
			return &api.UploadInitResponse{ChunkStartFrom: int32(currentChunk)}, nil
		}
	}
	log.Infof("[%s] File %s doesnt exists on server. Creating ...", req.ClientId, req.Filename)
	_, err = createNewFile(req.Filename, req.ClientId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	s.meta.AddFile(req.ClientId, req.Filename, req.ModifiedAt, int(req.Size), int(req.ChunksCount))
	return &api.UploadInitResponse{ChunkStartFrom: int32(0)}, nil
}

func (s *UploadServer) Upload(stream api.Uploader_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	fileMeta, err := NewFileMeta(req, &s.meta)
	if err != nil {
		return err
	}
	for {
		req, err = stream.Recv()
		if err != nil {
			return err
		}
		onNewChunk(req, stream, fileMeta, &s.meta)
		if int(req.ChunkId) == fileMeta.ChunksCount-1 {
			break
		}
	}
	err = stream.SendMsg(&api.UploadResponse{Status: 1})
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("[%s] %s received. Sending done status", fileMeta.ClientId, fileMeta.Filename)
	return nil
}

func onNewChunk(req *api.UploadRequest, stream api.Uploader_UploadServer, fileMeta FileMeta, serverMeta *api.Metadata) error {
	log.Infof("[%s] Received %d chunk.", req.ClientId, req.ChunkId+1)
	err := writeChunkToFile(fileMeta.File, req.Chunk)
	if err != nil {
		log.Error(err)
		return status.Errorf(codes.Internal, "Server error")
	}
	serverMeta.IncreaseCurrentChunk(req.ClientId, req.Filename, int(req.ChunkId))
	return nil
}

func createNewFile(filename, uuid string) (*os.File, error) {
	uploadedPath, err := getTargetDirectory()
	if err != nil {
		return nil, err
	}
	newFile, err := os.Create(fmt.Sprintf("%sclient_%s/%s", uploadedPath, uuid, filename))
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

func writeChunkToFile(file *os.File, chunk []byte) error {
	if _, err := file.Write(chunk); err != nil {
		return err
	}
	return nil
}

func checkUserFileExists(req *api.UploadInitRequest) (bool, error) {
	uploadedPath, err := getTargetDirectory()
	if err != nil {
		return false, err
	}
	clientPath := fmt.Sprintf("%sclient_%s/", uploadedPath, req.ClientId)
	// Check clientDirectory exists. Create if not
	if _, err := os.Stat(clientPath); err != nil {
		log.Infof("[%s] Client directory not found. Creating...", req.ClientId)
		err := createClientDirectory(req.ClientId)
		if err != nil {
			return false, err
		}
	}
	requestedFilePath, err := getTargetFilepath(req.ClientId, req.Filename)
	if err != nil {
		return false, err
	}
	// Check request file exists
	if _, err := os.Stat(requestedFilePath); err == nil {
		log.Infof("[%s] File with the same name is already in client directory", req.ClientId)
		return true, nil
	}
	return false, nil
}

func getTargetFilepath(clientId, filename string) (string, error) {
	targetDirectory, err := getTargetDirectory()
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%sclient_%s/%s", targetDirectory, clientId, filename), nil
}

func deleteUserFile(filepath string) error {
	if err := os.Remove(filepath); err != nil {
		return err
	}
	return nil
}

func getTargetDirectory() (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	uploadedPath := fmt.Sprintf("%s/uploaded/", rootPath)
	return uploadedPath, nil
}

func createClientDirectory(clientId string) error {
	uploadedPath, err := getTargetDirectory()
	if err != nil {
		return err
	}
	if err := os.Mkdir(fmt.Sprintf("%sclient_%s/", uploadedPath, clientId), os.ModePerm); err != nil {
		return err
	}
	return nil
}
