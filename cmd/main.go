package main

import (
	"MiniProject/internal/app/db"
	"MiniProject/internal/app/handler"
	"MiniProject/internal/app/utils"
	"MiniProject/internal/server"
	"flag"
	"fmt"
	"log"
)

func main() {
	migr := flag.Bool("migrate", false, fmt.Sprint("Migrating Entity"))
	rnddata := flag.Bool("rnddata", false, fmt.Sprint("Adding random data"))
	start := flag.Bool("start", false, fmt.Sprint("Starting server"))
	flag.Parse()

	if *migr {
		session := db.Connect()
		log.Println("Migrating")
		db.Migrate(session)
	}
	if *rnddata {
		session := db.Connect()
		log.Println("Creating Random Data")
		utils.Generate(session)

	}
	if *start {
		srv := new(server.Server)
		log.Println("Starting server")
		defer log.Println("End of Program")
		log.Println("Open the server")
		fmt.Println("Running and Serving on: http://127.0.0.1:5000")
		session := db.Connect()
		mHandler := handler.Handler{DB: session}
		if err := srv.Run(":5000", mHandler.InitRoutes()); err != nil {
			log.Fatalln(err)
		}
	}

}
