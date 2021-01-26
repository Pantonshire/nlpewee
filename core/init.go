package core

import "github.com/bbalet/stopwords"

func Init() {
	stopwords.DontStripDigits()
}
