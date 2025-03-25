package tdu_test

import (
	"testing"

	"github.com/josephpaul0/tdu"
	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	total_size, total_items, err := tdu.Stat("/media/10.0.0.90/ZimaOS-HD/Downloads")
	assert.NoError(t, err)
	assert.True(t, total_size > 0)
	assert.True(t, total_items > 0)
}
