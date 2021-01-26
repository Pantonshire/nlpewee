package server

import (
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
)

func serializeTokenizeResponse(sentences []core.Sentence) *pb.TokenizeResponse {
	sentenceMsgs := make([]*pb.Sentence, len(sentences))
	for i, sentence := range sentences {
		sentenceMsgs[i] = serializeSentence(sentence)
	}
	return &pb.TokenizeResponse{
		Sentences: sentenceMsgs,
	}
}

func serializeSentence(sentence core.Sentence) *pb.Sentence {
	tokens := sentence.Tokens()
	tokenMsgs := make([]*pb.Token, len(tokens))
	for i, token := range tokens {
		tokenMsgs[i] = serializeToken(token)
	}
	entities := sentence.Entities()
	entityMsgs := make([]*pb.Entity, len(entities))
	for i, entity := range entities {
		entityMsgs[i] = serializeEntity(entity)
	}
	return &pb.Sentence{
		Tokens:   tokenMsgs,
		Entities: entityMsgs,
	}
}

func serializeToken(token core.Token) *pb.Token {
	return &pb.Token{
		Text:  token.Text(),
		Stem:  token.Stem(),
		Tag:   token.Tag(),
		Label: token.Label(),
	}
}

func serializeEntity(entity core.Entity) *pb.Entity {
	return &pb.Entity{
		Text:  entity.Text(),
		Label: entity.Label(),
	}
}
