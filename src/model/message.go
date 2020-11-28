package model

import "github.com/volatiletech/null/v8"

type Message struct {
	Author string    `json:"author"`
	Body   null.JSON `json:"body"`
}
