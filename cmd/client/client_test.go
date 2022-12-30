package client

import (
	"bufio"
	"context"
	"grpc_loader/pkg/api"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

const myId = "tester"
const sendDataSize = 1024*1024*3 + 10 // slightly more than 15Mb

func TestUploadChunks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), SendTimeoutSec*time.Second)
	defer cancel()
	stream := TestStream{}
	var q api.Uploader_UploadClient
	q = &stream
	err := uploadChunks(q, ctx, getTestReader(sendDataSize), "test.file", 0, myId)
	assert.NoErrorf(t, err, "failed to upload chunks")
	wantChunksSent := (sendDataSize + api.OneChunkSize - 1) / api.OneChunkSize
	assert.Equal(t, wantChunksSent, stream.sentMsgCount, "invalid chunks count")
}

// TestReader emulates reading from file, returning identical symbols.
type TestReader struct {
	maxBytes int
	sent     int
}

func (r *TestReader) Read(b []byte) (int, error) {
	if r.sent == r.maxBytes {
		return 0, io.EOF
	}
	b[0] = '1'
	r.sent++
	return 1, nil
}

func getTestReader(maxBytes int) *bufio.Reader {
	return bufio.NewReader(&TestReader{
		maxBytes: maxBytes,
	})
}

// TestStream satisfies grpc stream interface allows us to test client without server.
type TestStream struct {
	sentMsgCount int
}

func (c *TestStream) CloseAndRecv() (*api.UploadResponse, error) {
	resp := &api.UploadResponse{Status: 1}
	return resp, nil
}

func (c *TestStream) RecvMsg(m interface{}) error {
	res := m.(*api.UploadResponse)
	res.Status = 1
	return nil
}

func (c *TestStream) SendMsg(m interface{}) error {
	c.sentMsgCount++
	return nil
}

func (c *TestStream) Header() (metadata.MD, error) {
	panic("should not be called")
}

func (c *TestStream) Trailer() metadata.MD {
	panic("should not be called")
}

func (c *TestStream) CloseSend() error {
	panic("should not be called")
}

func (c *TestStream) Context() context.Context {
	panic("should not be called")
}

func (c *TestStream) Send(*api.UploadRequest) error {
	panic("should not be called")
}
