package server

import (
	"grpc_loader/pkg/api"
	"grpc_loader/pkg/uploader"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const maxMsgSize = 1024 * 1024 * 2 // 2 MB

func RunServer() {
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxMsgSize))
	srv := uploader.NewUploadServer()
	api.RegisterUploaderServer(s, srv)
	l, err := net.Listen("tcp", api.ServerPort)

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
