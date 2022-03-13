package main

import (
	"context"
	"flag"
	"grpc-server/config"
	"grpc-server/pkg/api"
	"grpc-server/src/app"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var configPath = ""
	var cfg = new(config.Config)
	// ищем youtube api key
	cfgPath := flag.String("c", configPath, "path to your config file")
	flag.Parse()
	// если не указан флаг
	if (*cfgPath) == "" {
		// проверяем глобальную переменную
		if apiKey := os.Getenv("YTAPIKEY"); apiKey == "" {
			log.Fatal("you must choose youtube api key: set YTAPIKEY enviroment or pass a config using -c flag")
		} else {
			cfg.YoutubeApiKey = apiKey
		}
	} else {
		cfg.YoutubeApiKey = (*cfgPath)
	}

	ctx := context.TODO()
	// router
	grpcServer := grpc.NewServer()
	// handler
	ytService := new(app.YoutubeGRPCServer)
	// setup our handler
	if err := ytService.Setup(ctx, cfg); err != nil {
		log.Fatal(err)
	}
	// register handler in router
	api.RegisterYoutubePlaylistServer(grpcServer, ytService)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	// run the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
