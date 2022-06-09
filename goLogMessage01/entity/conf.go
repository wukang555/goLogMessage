package entity

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	FileName string `yaml:"fileName"`

	Email struct {
		SenMail    string `yaml:"sendMail"`
		Password   string `yaml:"password"`
		AcceptMail string `yaml:"acceptMail"`
		Subject    string `yaml:"subject"`
	}
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
