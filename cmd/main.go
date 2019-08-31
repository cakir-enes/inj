package main

import (
	_ "gogui/injector"
	"gogui/params"
	"io/ioutil"
	"log"

	_ "google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func main() {
	var cfg Config
	cfgFile, err := ioutil.ReadFile("/home/ecakir/Projects/gogui/gogui/cmd/conf.yaml")
	if err != nil {
		log.Panicf("ErrCantReadCfg: %s\n", err.Error())
	}
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		log.Panicf("ErrCantParseCfg: %s\n", err.Error())
	}
	mods := params.NewModules()
	for _, m := range cfg.Mods {
		acc, err := params.New(m.Address, m.Name)
		if err != nil {
			log.Printf("ErrCantConnect to %s -> %s\n", m.Name, err.Error())
			break
		}
		mods.AddModule(m.Name, acc)
		defer acc.Close()
	}
	log.Println(mods.GetAllParamInfo())

}

type Config struct {
	Mods []struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	} `yaml:"modules"`
}
