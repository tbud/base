package config

import (
	// "fmt"
	"path/filepath"
	// "reflect"
	"runtime"
	"strings"
)

type Config struct {
	options map[string]interface{}
}

func Load(fileName string) (config Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	config = Config{map[string]interface{}{}}
	scan := fileScanner{}

	if !filepath.IsAbs(fileName) {
		fileName, err = filepath.Abs(fileName)
		if err != nil {
			return
		}
	}

	err = scan.checkValid(fileName)
	if err != nil {
		return
	}

	err = scan.setOptions(config.options)
	return
}

func (c *Config) Int(key string) (result int, found bool) {
	result, found = 0, false
	value := c.getValue(key)
	if value == nil {
		return
	}

	if retFloat, ok := value.(float64); ok {
		return int(retFloat), ok
	}
	return
}

func (c *Config) IntDefault(key string, defaultValue int) int {
	result, found := c.Int(key)
	if !found {
		result = defaultValue
	}
	return result
}

func (c *Config) String(key string) (result string, found bool) {
	result, found = "", false
	value := c.getValue(key)
	if value == nil {
		return
	}

	result, found = value.(string)
	return
}

func (c *Config) StringDefault(key, defaultValue string) string {
	result, found := c.String(key)
	if !found {
		result = defaultValue
	}
	return result
}

func (c *Config) Bool(key string) (result, found bool) {
	result, found = false, false
	value := c.getValue(key)
	if value == nil {
		return
	}

	result, found = value.(bool)
	return
}

func (c *Config) BoolDefault(key string, defaultValue bool) bool {
	result, found := c.Bool(key)
	if !found {
		result = defaultValue
	}
	return result
}

func (c *Config) SubOptions(key string) *Config {
	return nil
}

func (c *Config) getValue(key string) interface{} {
	if len(key) == 0 {
		return nil
	}
	ops := c.options

	keys := strings.Split(key, ".")
	lastkeyIndex := len(keys) - 1
	for i := range keys {
		if value, ok := ops[keys[i]]; ok {
			if i == lastkeyIndex {
				return value
			} else {
				if ops, ok = value.(map[string]interface{}); !ok {
					return nil
				}
			}
		} else {
			return nil
		}
	}
	return nil
}
