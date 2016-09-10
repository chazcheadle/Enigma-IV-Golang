package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type config struct {
	Wheels map[string]map[string]string `yaml:"wheels"`
}

// Get configuration of cameras from yaml file.
func getConfig() *config {

	confFile := "/home/chaz/projects/golang/src/github.com/chazcheadle/enigma-iv-golang/encoder_wheels.yaml"

	if len(os.Args[1:]) > 0 {
		confFile = os.Args[1]
	}

	conf := &config{}

	f, err := os.Open(confFile)
	defer f.Close()
	if err != nil {
		log.Info("Could not open %s.", confFile)
		log.Info("Error: ", err)
	}

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Info("Error reading %s.", confFile)
		log.Info("Error: ", err)
	}

	err = yaml.Unmarshal(d, conf)
	if err != nil {
		log.Info("Error: ", err)
		log.Fatal("Could not parse ", confFile)
	}

	return conf
}
