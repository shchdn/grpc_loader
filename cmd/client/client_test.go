package client

import (
	"github.com/stretchr/testify/assert"
	"grpc_loader/cmd/client/test_utils"
	"grpc_loader/pkg/api"
	"testing"
)

const myId = "tester"
const sendDataSize = 1024*1024*3 + 10 // slightly more than 15Mb

func TestUploadChunks(t *testing.T) {
	stream := test_utils.TestStream{}
	var q api.Uploader_UploadClient
	q = &stream
	err := uploadChunks(q, test_utils.GetTestReader(sendDataSize), "test.file", 0, myId)
	assert.NoErrorf(t, err, "failed to upload chunks")
	wantChunksSent := (sendDataSize + api.OneChunkSize - 1) / api.OneChunkSize
	assert.Equal(t, wantChunksSent, stream.SentMsgCount, "invalid chunks count")
}
