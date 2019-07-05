package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	
	"gopkg.in/yaml.v2"
)

// Config is the configuration struct
type Config struct {
	ENV    string `yaml:ENV`
	Debug  bool   `yaml:"debug"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Firebase struct {
		ApiKey  string `yaml:"api_key"`
		TestPsw string `yaml:"test_psw"`
	} `yaml:"firebase"`
}

var cfg Config

// FromFile gets configuration from a file
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	
	return &cfg, nil
}

// Init must be called before server init to read & init config
func Init(path string) *Config {
	p := flag.String("config", path, "path of the yml config file")
	flag.Parse()
	
	_, err := FromFile(*p)
	
	if err != nil {
		log.Fatal(err)
	}
	
	if port := os.Getenv("PORT"); len(port) > 0 {
		cfg.Server.Port = port
	}
	
	return &cfg
}
