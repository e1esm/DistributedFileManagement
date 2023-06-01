package models

import "github.com/google/uuid"

type File struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Content []byte    `json:"bytes"`
}

func NewFile(content []byte) *File {
	return &File{Id: uuid.New(), Content: content}
}
