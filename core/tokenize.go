package core

import (
	"github.com/bbalet/stopwords"
	"github.com/jdkato/prose/v2"
	"github.com/kljensen/snowball"
	"strings"
)

func Tokenize(text string, lang Language, clean bool) ([]Sentence, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	if docSentences := doc.Sentences(); len(docSentences) > 1 {
		var sentences []Sentence
		for _, docSentence := range docSentences {
			recursiveResult, err := Tokenize(docSentence.Text, lang, clean)
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
		var cleaned string

		if clean && stopwordsLangCode != nil {
			cleaned = cleanToken(docToken.Text, *stopwordsLangCode)
			if len(cleaned) == 0 {
				continue
			}
		}

		token := Token{
			tag:   docToken.Tag,
			label: docToken.Label,
		}

		if len(cleaned) > 0 {
			token.text = cleaned
		} else {
			token.text = docToken.Text
		}

		if snowballLangCode != nil {
			stem, err := stemToken(token.text, *snowballLangCode)
			if err != nil {
				return nil, err
			}
			token.stem = stem
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
