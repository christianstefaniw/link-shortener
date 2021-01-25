package main

import (
	"./server"
	"./storage"
)


func main() {
	storage.LoadEnv()
	storage.Collection = storage.LoadStorage()
	server.New()
}
