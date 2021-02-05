package ct

import (
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestGetLogsForPeriod(t *testing.T) {
	start := time.Now()
	end := time.Now()
	got := GetLogsForPeriod(start, end)

	assert.IsType(t, got, []string{})
}
