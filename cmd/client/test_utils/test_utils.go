package test_utils

import (
	"bufio"
	"context"
	"google.golang.org/grpc/metadata"
	"grpc_loader/pkg/api"
	"io"
)

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

func GetTestReader(maxBytes int) *bufio.Reader {
	return bufio.NewReader(&TestReader{
		maxBytes: maxBytes,
	})
}

// TestStream satisfies grpc stream interface allows us to test client without server.
type TestStream struct {
	SentMsgCount int
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
	c.SentMsgCount++
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
