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
	msg := &pb.Token{
		Full:  serializeText(token.Full()),
		Stem:  serializeText(token.Stem()),
		Label: token.Label(),
	}

	rawTag := token.Tag()
	if tag, ok := serializeTag(rawTag); ok {
		msg.PosTag = &pb.Token_Tag{Tag: tag}
	} else {
		msg.PosTag = &pb.Token_Other{Other: rawTag}
	}

	return msg;
}

func serializeText(text core.Text) *pb.Text {
	return &pb.Text{
		Raw:     text.Raw(),
		Cleaned: text.Cleaned(),
	}
}

func serializeEntity(entity core.Entity) *pb.Entity {
	return &pb.Entity{
		Text:  entity.Text(),
		Label: entity.Label(),
	}
}

func serializeTag(tag string) (pb.Tag, bool) {
	switch tag {
	case "(":
		return pb.Tag_L_PAREN, true
	case ")":
		return pb.Tag_R_PAREN, true
	case ",":
		return pb.Tag_COMMA, true
	case ":":
		return pb.Tag_COLON, true
	case ".":
		return pb.Tag_PERIOD, true
	case "''":
		return pb.Tag_CLOSING_QUOTE, true
	case "``":
		return pb.Tag_OPENING_QUOTE, true
	case "#":
		return pb.Tag_NUMBER_SIGN, true
	case "$":
		return pb.Tag_CURRENCY, true
	case "CC":
		return pb.Tag_CC, true
	case "CD":
		return pb.Tag_CD, true
	case "DT":
		return pb.Tag_DT, true
	case "EX":
		return pb.Tag_EX, true
	case "FW":
		return pb.Tag_FW, true
	case "IN":
		return pb.Tag_IN, true
	case "JJ":
		return pb.Tag_JJ, true
	case "JJR":
		return pb.Tag_JJR, true
	case "JJS":
		return pb.Tag_JJS, true
	case "LS":
		return pb.Tag_LS, true
	case "MD":
		return pb.Tag_MD, true
	case "NN":
		return pb.Tag_NN, true
	case "NNP":
		return pb.Tag_NNP, true
	case "NNPS":
		return pb.Tag_NNPS, true
	case "NNS":
		return pb.Tag_NNS, true
	case "PDT":
		return pb.Tag_PDT, true
	case "POS":
		return pb.Tag_POS, true
	case "PRP":
		return pb.Tag_PRP, true
	case "PRP$":
		return pb.Tag_PRPS, true
	case "RB":
		return pb.Tag_RB, true
	case "RBR":
		return pb.Tag_RBR, true
	case "RBS":
		return pb.Tag_RBS, true
	case "RP":
		return pb.Tag_RP, true
	case "SYM":
		return pb.Tag_SYM, true
	case "TO":
		return pb.Tag_TO, true
	case "UH":
		return pb.Tag_UH, true
	case "VB":
		return pb.Tag_VB, true
	case "VBD":
		return pb.Tag_VBD, true
	case "VBG":
		return pb.Tag_VBG, true
	case "VBN":
		return pb.Tag_VBN, true
	case "VBP":
		return pb.Tag_VBP, true
	case "VBZ":
		return pb.Tag_VBZ, true
	case "WDT":
		return pb.Tag_WDT, true
	case "WP":
		return pb.Tag_WP, true
	case "WP$":
		return pb.Tag_WPS, true
	case "WRB":
		return pb.Tag_WRB, true
	default:
		return 0, false
	}
}
