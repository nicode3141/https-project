package main

import (
	"https-project/internal/server"
	"log"
)

func main(){
	srv := server.New()

	log.Println("Server listening on port 8080 localhost")

	err := srv.ListenAndServe();

	if(err != nil) {
		log.Fatal(err);
	}
}