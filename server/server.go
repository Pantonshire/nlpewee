package server

import (
	"context"
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
)

type NLPeweeServer struct {}

func (server NLPeweeServer) Tokenize(ctx context.Context, req *pb.TokenizeRequest) (*pb.TokenizeResponse, error) {
	return server.tokenize(req, false)
}

func (server NLPeweeServer) TokenizeClean(ctx context.Context, req *pb.TokenizeRequest) (*pb.TokenizeResponse, error) {
	return server.tokenize(req, true)
}

func (server NLPeweeServer) tokenize(req *pb.TokenizeRequest, clean bool) (*pb.TokenizeResponse, error) {
	text, lang, err := desTokenizeRequest(req)
	if err != nil {
		return nil, err
	}
	sentences, err := core.Tokenize(text, lang, clean)
	if err != nil {
		return nil, err
	}
	return serializeTokenizeResponse(sentences), nil
}
