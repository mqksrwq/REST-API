package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dataBaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DataBaseURL = dataBaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s "+
				"CASCADE", strings.Join(tables, ", "))); err != nil {
			}
		}
		s.Close()
	}

}
