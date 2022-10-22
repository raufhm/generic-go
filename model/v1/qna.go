package model

type ContactForm struct {
	Name  string
	Email string
}

type Address struct {
	PropertyType string
	BlockNo      string
}

type SingleAnswer string
type MultipleAnswer []string

type IAnswer interface {
	SingleAnswer | MultipleAnswer | ICustom
}

type ICustom interface {
	ContactForm | Address
}

type Qna[ans IAnswer] struct {
	Question string `json:"Question"`
	Answer   ans    `json:"Answer"`
}
