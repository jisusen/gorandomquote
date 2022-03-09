package quote

import (
	"fmt"
)

type Quote struct {
	Text   string
	Author string
}

func (q Quote) handleQuote() string {
	q.Text = "test"
	q.Author = "author"
	return fmt.Sprintf("%s - %s", q.Text, q.Author)
}
