package core

type Sentence struct {
	tokens   []Token
	entities []Entity
}

func (sentence *Sentence) Tokens() []Token {
	if sentence == nil {
		return nil
	}
	return sentence.tokens
}

func (sentence *Sentence) Entities() []Entity {
	if sentence == nil {
		return nil
	}
	return sentence.entities
}

type Token struct {
	full, stem Text
	tag, label string
}

func (token *Token) Full() Text {
	if token == nil {
		return Text{}
	}
	return token.full
}

func (token *Token) Stem() Text {
	if token == nil {
		return Text{}
	}
	return token.stem
}

func (token *Token) Tag() string {
	if token == nil {
		return ""
	}
	return token.tag
}

func (token *Token) Label() string {
	if token == nil {
		return ""
	}
	return token.label
}

type Text struct {
	raw, cleaned string
}

func (text *Text) Raw() string {
	if text == nil {
		return ""
	}
	return text.raw
}

func (text *Text) Cleaned() string {
	if text == nil {
		return ""
	}
	return text.cleaned
}

type Entity struct {
	text, label string
}

func (entity *Entity) Text() string {
	if entity == nil {
		return ""
	}
	return entity.text
}

func (entity *Entity) Label() string {
	if entity == nil {
		return ""
	}
	return entity.label
}
