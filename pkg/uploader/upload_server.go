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
	ClientId     api.ClientId
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
	chunksCount := serverMeta.GetParams(req.ClientId, req.Filename).ChunksCount
	return FileMeta{file, filepath, filename, uploadedPath, req.ClientId, chunksCount}, nil
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
		modifiedEqual := s.meta.IsMetaExists(req.ClientId, req.Filename, req.ModifiedAt)
		// Delete if not equal
		if !modifiedEqual {
			log.Infof("[%s] Files are not equal. Deleting file from server...", req.ClientId)
			err = deleteUserFile(requestedFilePath)
			if err != nil {
				return nil, err
			}
		} else {
			currentChunk := s.meta.GetParams(req.ClientId, req.Filename).CurrentChunk
			log.Infof("[%s] Files are equal. Current chunk is - %d", req.ClientId, currentChunk)
			// Return next chunk
			return &api.UploadInitResponse{ChunkStartFrom: int32(currentChunk + 1)}, nil
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
	fileMetaInited := false
	var fileMeta FileMeta
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Error("Failed to receive: ", err)
			return err
		}
		if !fileMetaInited {
			fileMeta, err = NewFileMeta(req, &s.meta)
			if err != nil {
				return err
			}
		}
		err = s.onNewChunk(req, fileMeta, &s.meta)
		if err != nil {
			log.Errorf("[%s] Failed to handle chunk: %s", req.ClientId, err)
		}
		if int(req.ChunkId) == fileMeta.ChunksCount-1 {
			break
		}
	}
	err := stream.SendMsg(&api.UploadResponse{Status: 1})
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("[%s] %s received. Sending done status", fileMeta.ClientId, fileMeta.Filename)
	return nil
}

func (s *UploadServer) onNewChunk(req *api.UploadRequest, fileMeta FileMeta, serverMeta *api.Metadata) error {
	log.Infof("[%s] Received %d chunk.", req.ClientId, req.ChunkId+1)
	err := writeChunkToFile(fileMeta.File, req.Chunk)
	if err != nil {
		log.Error("Write to file failed: ", err)
		// Write to file failed with error.
		// File could be corrupted, we cant just retry writing chunk.
		// So we have to remove file and meta and then wait client to retry request.
		err = deleteUserFile(fileMeta.Filepath)
		s.meta.RemoveMeta(req.ClientId, req.Filename)
		if err != nil {
			// At this point we can't handle this error proper way.
			log.Errorf("Failed to remove corrupted file: %s", err)
		}
		return status.Errorf(codes.Internal, "Server error")
	}
	serverMeta.SetCurrentChunk(req.ClientId, req.Filename, int(req.ChunkId))
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
	_, err := file.Write(chunk)
	return err
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

func getTargetFilepath(clientId api.ClientId, filename string) (string, error) {
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

func createClientDirectory(clientId api.ClientId) error {
	uploadedPath, err := getTargetDirectory()
	if err != nil {
		return err
	}
	if err := os.Mkdir(fmt.Sprintf("%sclient_%s/", uploadedPath, clientId), os.ModePerm); err != nil {
		return err
	}
	return nil
}
