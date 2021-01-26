package core

type Language uint8

const (
	English Language = iota
)

func (lang Language) toStopwordsCode() *string {
	switch lang {
	case English:
		code := new(string)
		*code = "en"
		return code
	default:
		return nil
	}
}

func (lang Language) toSnowballCode() *string {
	switch lang {
	case English:
		code := new(string)
		*code = "english"
		return code
	default:
		return nil
	}
}
