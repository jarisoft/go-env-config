package config

type Config interface {
	GetVariables(variables map[string]*string) error
	GetVariable(s string) (string, error)
}
