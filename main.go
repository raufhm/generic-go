package main

import (
	"fmt"
	"github.com/raufhm/generic/model/v1"
)

func main() {
	single := model.Qna[model.SingleAnswer]{
		Question: "what's your name?",
		Answer:   "single",
	}
	multiple := model.Qna[model.MultipleAnswer]{
		Question: "What are your listing?",
		Answer:   []string{"HDB", "CONDO/APARTMENT"},
	}
	contactForm := model.Qna[model.ContactForm]{
		Question: "Contact details?",
		Answer: model.ContactForm{
			Name:  "rauf",
			Email: "rauf@email.com",
		},
	}
	address := model.Qna[model.Address]{
		Question: "Address details?",
		Answer: model.Address{
			PropertyType: "HDB",
			BlockNo:      "10",
		},
	}

	fmt.Println(single)
	fmt.Println(multiple)
	fmt.Println(contactForm)
	fmt.Println(address)
}
