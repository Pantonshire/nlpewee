package server

import (
	"context"
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type empty struct{}

type Server struct {
	killChan       chan empty
	tokenizeBucket chan empty
}

var serverShutdownError = status.Error(codes.Unavailable, "server shut down while waiting to process request")

func NewServer(bucketSize uint) *Server {
	server := Server{
		killChan:       make(chan empty),
		tokenizeBucket: make(chan empty, bucketSize),
	}
	var i uint
	for i = 0; i < bucketSize; i++ {
		server.tokenizeBucket <- empty{}
	}
	return &server
}

func (server *Server) Kill() {
	close(server.killChan)
}

func (server *Server) p() error {
	select {
	case <-server.tokenizeBucket:
		return nil
	case <-server.killChan:
		return serverShutdownError
	}
}

func (server *Server) v() {
	server.tokenizeBucket <- empty{}
}

func (server *Server) Tokenize(ctx context.Context, req *pb.TokenizeRequest) (*pb.TokenizeResponse, error) {
	if err := server.p(); err != nil {
		return nil, err
	}
	defer server.v()
	text, lang, err := desTokenizeRequest(req)
	if err != nil {
		return nil, err
	}
	sentences, err := core.Tokenize(text, lang)
	if err != nil {
		return nil, err
	}
	return serializeTokenizeResponse(sentences), nil
}
