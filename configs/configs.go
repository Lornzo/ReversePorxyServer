package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configs struct {
	UseSSL    bool               `json:"usessl"`
	CertFiles CertFiles          `json:"certfiles"`
	Services  map[string]Service `json:"services"`
}

var confs configs

func Load(path string) (err error) {

	var (
		jsonByte []byte
	)

	if jsonByte, err = ioutil.ReadFile(path); err != nil {
		var returnErr error = fmt.Errorf("config loading : %s", err.Error())
		panic(returnErr)
	}

	json.Unmarshal(jsonByte, &confs)
	return
}

func Get() (confs configs) {
	return confs
}

func GetService(domain string) (s Service, err error) {
	var (
		isOk bool
	)
	if s, isOk = confs.Services[domain]; !isOk {
		err = fmt.Errorf("domain:" + domain + " is not exist")
	}
	return

}
