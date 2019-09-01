package main

import (
	_ "gogui/injector"
	"gogui/params"
	"io/ioutil"
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	_ "google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func basic() map[string][]params.ParamInfo {
	m := make(map[string][]params.ParamInfo)
	m["ABC"] = []params.ParamInfo{params.ParamInfo{Name: "ASD"}}
	return m
}

// //go:generate sh -c protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/*.proto
func main() {
	var cfg Config
	cfgFile, err := ioutil.ReadFile("/home/ecakir/Projects/gogui/gogui/conf.yaml")
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

	js := mewn.String("../frontend/dist/app.js")
	css := mewn.String("../frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Quotes",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	// basic := func() map[string][]params.ParamInfo {
	// 	return mods.GetAllParamInfo()
	// }
	app.Bind(&mods)
	app.Run()

}

type Config struct {
	Mods []struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	} `yaml:"modules"`
}
