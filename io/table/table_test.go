package table

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestFromObject(t *testing.T) {
	obj := map[string]interface{}{
		"key":  "value",
		"key2": 2,
		"key3": true,
	}
	builder := FromObject(obj)
	assert.NotNil(t, builder)

	act := builder.Build()
	exp := `
key  	key2	key3
-----	----	----
value	2   	true
`
	act = strings.Trim(act, "\n")
	exp = strings.Trim(exp, "\n")
	assert.Equal(t, exp, act)

}

func TestFromInvalidObject(t *testing.T) {
	obj := time.Now()
	builder := FromObject(obj)
	assert.NotNil(t, builder)

	act := builder.Build()
	exp := `
error                                                                      
---------------------------------------------------------------------------
json: cannot unmarshal string into Go value of type map[string]interface {}
`
	act = strings.Trim(act, "\n")
	exp = strings.Trim(exp, "\n")
	assert.Equal(t, exp, act)

}

func TestDefaultFormatter(t *testing.T) {
	obj := map[string]interface{}{
		"key":  "value",
		"key2": 2,
		"key3": true,
	}
	builder := FromObject(obj).DefaultFormatter()
	assert.NotNil(t, builder)

	act := builder.Build()
	exp := `
key                            	key2                           	key3                           
-------------------------------	-------------------------------	-------------------------------
map[key:value key2:2 key3:true]	map[key:value key2:2 key3:true]	map[key:value key2:2 key3:true]
`
	act = strings.Trim(act, "\n")
	exp = strings.Trim(exp, "\n")
	assert.Equal(t, exp, act)

}
