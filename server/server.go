package server

import (
	"context"
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
)

type NLPeweeServer struct {}

func (server NLPeweeServer) Tokenize(ctx context.Context, req *pb.TokenizeRequest) (*pb.TokenizeResponse, error) {
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
