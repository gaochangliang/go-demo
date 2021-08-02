package config

type Server struct {
	System System
	Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
