package core

import (
	"github.com/bbalet/stopwords"
	"github.com/jdkato/prose/v2"
	"github.com/kljensen/snowball"
	"strings"
)

func Tokenize(text string, lang Language) ([]Sentence, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	if docSentences := doc.Sentences(); len(docSentences) > 1 {
		var sentences []Sentence
		for _, docSentence := range docSentences {
			recursiveResult, err := Tokenize(docSentence.Text, lang)
			if err != nil {
				return nil, err
			}
			sentences = append(sentences, recursiveResult...)
		}
		return sentences, nil
	}

	stopwordsLangCode := lang.toStopwordsCode()
	snowballLangCode := lang.toSnowballCode()

	var sentence Sentence

	for _, docToken := range doc.Tokens() {
		token := Token{
			full:  Text{raw: docToken.Text},
			stem:  Text{},
			tag:   docToken.Tag,
			label: docToken.Label,
		}

		if snowballLangCode != nil {
			var err error
			token.stem.raw, err = stemToken(token.full.raw, *snowballLangCode)
			if err != nil {
				return nil, err
			}
		}

		if stopwordsLangCode != nil {
			token.full.cleaned = cleanToken(token.full.raw, *stopwordsLangCode)

			if len(token.stem.raw) > 0 {
				token.stem.cleaned = cleanToken(token.stem.raw, *stopwordsLangCode)
			}
		}

		sentence.tokens = append(sentence.tokens, token)
	}

	for _, docEntity := range doc.Entities() {
		sentence.entities = append(sentence.entities, Entity{
			text:  docEntity.Text,
			label: docEntity.Label,
		})
	}

	return []Sentence{sentence}, nil
}

func cleanToken(token string, langCode string) string {
	return strings.ReplaceAll(stopwords.CleanString(token, langCode, false), " ", "")
}

func stemToken(token string, langCode string) (string, error) {
	return snowball.Stem(token, langCode, true)
}
