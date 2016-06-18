package main

import (
	"github.com/markbates/goth"
	"log"
)

func main_() {
	log.Println("fff")
	providers := goth.GetProviders()
	log.Printf("%v", providers)
	for _, provider := range providers {
		log.Println(provider.Name())
	}
}
