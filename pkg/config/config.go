package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	params map[string]interface{}
}

func New() *Config {
	c := new(Config)
	c.params = make(map[string]interface{})
	return c
}

func (c *Config) SetEnv(key string, defaultValue interface{}) {

	value := env(key)
	if value != "" {
		c.params[key] = value
		return
	}
	c.params[key] = defaultValue
}

//Returns the value of env variable, if not found returns the default value
func (c *Config) Get(key string) string {
	val := c.getValue(key, reflect.String)
	if val != nil {
		return val.(string)
	}
	return ""
}

//Returns env. variable as a slice of strings
func (c *Config) SetEnvVars(key, sep, defaultValue string) {
	val := env(key)
	if val != "" {
		vals := strings.Split(val, sep)
		c.params[key] = vals
	} else {
		c.params[key] = strings.Split(defaultValue, sep)
	}
}

func (c *Config) GetVars(key string) []string {
	val := c.getValue(key, reflect.Slice)
	if val == nil {
		return []string{}
	}
	return val.([]string)
}

//Gets the int value from config if present, else looks up from env
func (c *Config) SetEnvInt(key string, defaultValue int) {
	value := env(key)
	ret, err := strconv.Atoi(value)
	if err != nil {
		ret = defaultValue
	}
	c.params[key] = ret
}

func (c *Config) GetInt(key string) int {
	val := c.getValue(key, reflect.Int)
	if val != nil {
		return val.(int)
	}
	return 0
}

func (c *Config) SetBool(key string, defaultValue bool) {
	value := env(key)
	val, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Invalid bool type set in env. for key:%s Defaulting to %b. Error: %v", key, defaultValue, err)
		val = defaultValue
	}
	c.params[key] = val
}

func (c *Config) GetBool(key string) bool {
	val := c.getValue(key, reflect.Bool)
	if val != nil {
		return val.(bool)
	}
	return false
}

//Returns the value if the requested type matches else returns nil
func (c *Config) getValue(key string, valueType reflect.Kind) interface{} {
	if val, ok := c.params[key]; ok {
		if getType(val) == valueType {
			return val
		}
	}
	return nil
}

//Gets the type of the val using reflection
func getType(val interface{}) reflect.Kind {
	return reflect.TypeOf(val).Kind()
}

func env(key string) string {
	return os.Getenv(key)
}

func (c *Config) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "charset=utf-8")
	keys := make([]string, len(c.params))
	i := 0
	for key := range c.params {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(w, "%s : %v\n", key, c.params[key])
	}
}
