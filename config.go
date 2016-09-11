package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type config struct {
	Wheels map[string]map[string]string `yaml:"wheels"`
}

// Get configuration of cameras from yaml file.
func getConfig() *config {

	// Get current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Default config file name.
	confFile := dir + "/encoder_wheels.yaml"
	if len(os.Args[1:]) > 0 {
		confFile = os.Args[1]
	}

	f, err := os.Open(confFile)
	defer f.Close()
	if err != nil {
		log.Warn("Could not open: ", confFile)
		log.Fatal("Error: ", err)
	}

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Warn("Error reading: ", confFile)
		log.Fatal("Error: ", err)
	}

	conf := &config{}

	err = yaml.Unmarshal(d, conf)
	if err != nil {
		log.Warn("Error: ", err)
		log.Fatal("Could not parse ", confFile)
	}

	return conf
}
