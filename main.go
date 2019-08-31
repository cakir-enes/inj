package main

import (
	"context"
	pb "gogui/injector"
	"gogui/params"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// func basic() string {
// 	return "World!ddsdsdasd"
// }

func main() {

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Printf("couldnt connect %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	client := pb.NewInjectorClient(conn)
	acc, err := params.New("localhost:8090", "FTE")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer acc.Close()
	mods := params.NewModules()
	mods.AddModule("FTE", acc)
	log.Println(mods.GetAllParamInfo())
	log.Println(acc.GetAllParamInfo())
	log.Println(client.GetAllParameterInfos(context.Background(), &empty.Empty{}))
	// stream := MultiGet([]string{"a", "b", "c"})
	// js := mewn.String("./frontend/build/static/js/main.js")
	// css := mewn.String("./frontend/build/static/css/main.css")

	// app := wails.CreateApp(&wails.AppConfig{
	// 	Width:     1024,
	// 	Height:    768,
	// 	Title:     "gogui",
	// 	JS:        js,
	// 	CSS:       css,
	// 	Colour:    "#131313",
	// 	Resizable: true,
	// })
	// app.Bind(basic)
	// app.Run()
}
