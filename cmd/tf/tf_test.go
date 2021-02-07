package tf

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestParser(t *testing.T) {
	t.Helper()

	t.Run("It parses a terraform file", func(t *testing.T) {
		path := "./fixtures/main.tf"
		want := new(Terraform)
		got := Parser(path)
		assert.IsType(t, want, got)
	})
}
