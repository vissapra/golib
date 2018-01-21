package config

import (
	"testing"

	"github.com/vissapra/golib/pkg/assert"
)

func TestConfig_GetEnvInt(t *testing.T) {
	config := New()
	config.SetEnvInt("int_val", 1)
	assert.AssertEquals(t, config.GetInt("int_val"), 1, "Should be 1")
	config.SetEnvVars("string_slice", ",", "a,b,c")
	config.SetEnvInt("int_val", 2)
	assert.AssertEquals(t, config.GetVars("string_slice"), []string{"a", "b", "c"}, "Should be {a,b,c}")
}

func TestConfig_GetEnv(t *testing.T) {
	config := New()
	config.params["int_val0"] = 0
	config.params["int_val1"] = "1"
	config.params["int_val2"] = 2
	//types don't match
	assert.AssertEquals(t, config.Get("int_val0"), "", "Should be default to \"\" as the type doesn't match")
	assert.AssertEquals(t, config.GetInt("int_val0"), 0, "Should be 0")
	assert.AssertEquals(t, config.Get("int_val1"), "1", "Should be 1")
	assert.AssertEquals(t, config.GetInt("int_val2"), 2, "Should be 2, as types match")
	assert.AssertEquals(t, config.Get("non_existent"), "", "Should be empty")
}

func TestConfig_GetEnvVars(t *testing.T) {
	config := New()
	config.params["int_val"] = 1
	val := []string{"a", "b", "c"}
	config.params["string_slice"] = val
	assert.AssertEquals(t, config.GetVars("string_slice"), val, "Shouldn't be the default")
	config.SetEnvVars("string_slice", ",", "1,2,3")
	assert.AssertEquals(t, config.GetVars("string_slice"), []string{"1", "2", "3"}, "Should be {\"1\", \"2\", \"3\"}")
}
