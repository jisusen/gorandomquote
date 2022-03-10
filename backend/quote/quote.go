package quote

import (
	"fmt"
)

type Quote struct {
	Text   string
	Author string
}

func (q Quote) getQuote() string {
	q.Text = "test"
	q.Author = "author"
	return fmt.Sprintf("%s - %s", q.Text, q.Author)
}
