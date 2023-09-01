package sluggfy

import (
	"fmt"

	"github.com/gosimple/slug"
)

func NewUID(text string, uid string) string {
	fmt.Print(string(uid[len(uid)-5:]))
	text = fmt.Sprintf("%s-%s", text, string(uid[len(uid)-5:]))
	return New(text)
}

func New(text string) string {
	return slug.Make(text)
}
