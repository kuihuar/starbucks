package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

//var Conf = &conf
//
//type conf struct {
//	Server server  `yaml:server`
//
//}
type BasicConf struct {
	Env string `yaml:"local.env"`
	//Log zlog.LogConfig
	//Pprof PprofConfig
}
var (
	Basic BasicConf
)

func InitConf() {
	LoadConf("config.yaml", "", &Basic)
}

func LoadConf(filename, subConf string, s interface{}){
	var path string
	path = filepath.Join("conf",subConf, filename)
	if yamlFile, err := ioutil.ReadFile(path); err != nil {
		panic(filename + "get error: %v"+ err.Error())
	}else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " Umarshal error: %v" + err.Error())
	}
}