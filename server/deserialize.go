package server

import (
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
)

type DeserializationError uint8

const (
	InvalidLanguageError DeserializationError = iota
)

func (err DeserializationError) Error() string {
	switch err {
	case InvalidLanguageError:
		return "Invalid language"
	default:
		return "Unknown deserialization error"
	}
}

func desLanguage(lang pb.Language) (core.Language, error) {
	switch lang {
	case pb.Language_ENGLISH:
		return core.English, nil
	default:
		return 0, InvalidLanguageError
	}
}

func desTokenizeRequest(req *pb.TokenizeRequest) (text string, lang core.Language, err error) {
	text = req.GetText()
	lang, err = desLanguage(req.GetLanguage())
	if err != nil {
		return "", 0, err
	}
	return text, lang, nil
}
