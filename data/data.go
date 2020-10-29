// Package data contains everything related to reading data files
package data

import (
	"git.sr.ht/~hjertnes/when/utils"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)



func getFilename() string{
	filename, isSet := os.LookupEnv("WHEN_FILE")
	if !isSet {
		filename = "~/txt/when.yml"
	}

	return utils.ReplaceTilde(filename)
}

func read(filename string) (map[string]time.Time, error){
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := make(map[string]time.Time)

	err = yaml.Unmarshal(content, result)

	return result, err
}

func write(filename string, data map[string]time.Time) error{
	content, err := yaml.Marshal(data)
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(filename, content, os.ModePerm)
	return err
}

func create(filename string) (map[string]time.Time, error){
	conf := make(map[string]time.Time)
	content, err := yaml.Marshal(conf)
	if err != nil{
		return conf, err
	}
	err = ioutil.WriteFile(filename, content, os.ModePerm)
	return conf, err
}

func Read() (map[string]time.Time, error){
	filename := getFilename()

	if utils.Exist(filename){
		return read(filename)
	} else {
		return create(filename)
	}
}

func Write(data map[string]time.Time) error{
	filename := getFilename()

	if utils.Exist(filename){
		return write(filename, data)
	} else {
		_, err := create(filename)
		return err
	}
}
