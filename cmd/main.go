package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrKrabsmr/test-task-verba-group/configs"
	server "github.com/mrKrabsmr/test-task-verba-group/internal"
)

var (
	version = flag.Int("v", 1, "choice a version")
	debug   = flag.Bool("debug", false, "debug mode")
	initDB  = flag.Bool("init", false, "init db")
)

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := configs.Config{
		Address:   os.Getenv("ADDRESS"),
		DBAddress: os.Getenv("DBADDRESS"),
		Version:   *version,
		Debug:     *debug,
	}

	apiServer := server.NewAPIServer(config)
	apiServer.MustRun(*initDB)
}
