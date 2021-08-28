// Provides logic to read and parse env files
package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

// EnvConfig handles configuration from env files.
// It implements the config.Config interface.
type EnvConfig struct {
	Environment     *string // the environment that relates to the file to be parsed
	EnvironmentPath *string // the location where the environment file is located
}

func (c EnvConfig) GetVariable(s string) (string, error) {
	env:= ".env"
	if c.Environment != nil {
		env += "." + *c.Environment
	}

	path := "./"
	if c.EnvironmentPath != nil {
		if !strings.HasSuffix(*c.EnvironmentPath, "/") {
			path = *c.EnvironmentPath + "/"
		} else {
			path = *c.EnvironmentPath
		}
	}

	err := godotenv.Load(fmt.Sprintf("%s%s", path, env))

	if err != nil {
		return "", err
	}

	return os.Getenv(s), nil
}

// GetVariables reads each key of the given map, searches the env file for this
// key and assign the value to the maps variable.
func (c EnvConfig) GetVariables(variables map[string]*string) error {
	for idx := range variables {
		val, err := c.GetVariable(idx)
		if err != nil {
			return fmt.Errorf("error while retrieving value for %s in enviroment file %s in path %s: %v", idx, *c.Environment, *c.EnvironmentPath, err)
		}
		// Reset address to
		*variables[idx] = val
	}

	return nil
}