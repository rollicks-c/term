package args

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDefaults(t *testing.T) {

	ac := NewCollector([]string{})
	act, err := ac.GetDateAbs(0)
	assert.Error(t, err)
	assert.True(t, act.IsZero())

	now := time.Now()
	act, err = ac.GetDateAbs(0, WithDefault(now))
	assert.NoError(t, err)
	assert.Equal(t, now, act)

}
