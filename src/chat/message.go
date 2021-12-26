package chat

import "github.com/volatiletech/null/v8"

type Message struct {
	Author   string    `json:"author"`              //
	Body     null.JSON `json:"body"`                //
	FileName string    `json:"file_name" store:"-"` //
	FileType string    `json:"file_type" store:"-"` //
	FileSize int       `json:"file_size" store:"-"` //
	FileData []byte    `json:"file_data" store:"-"` //
}
