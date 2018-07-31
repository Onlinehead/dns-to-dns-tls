package main

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// Config structure to read from yaml configuration file
type Config struct {
	Server     serverConf    `yaml:"server"`
	Upstreams  []string      `yaml:"upstreams"`
	ReqTimeout time.Duration `yaml:"requestTimeout"`
	ResTimeout time.Duration `yaml:"responseTimeout"`
}

type serverConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	TCP  bool   `yaml:"tcp"`
	UDP  bool   `yaml:"udp"`
}

func (c *Config) readConfig(file string) *Config {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot read the config file: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Cannot unmarshal the config file: %v", err)
	}
	if !c.Server.TCP && !c.Server.UDP {
		c.Server.TCP = true
		log.Println("TCP and UDP serving disabled in config by default, force TCP")
	}
	return c
}
