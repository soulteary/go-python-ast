package main

import (
	"log"

	python3 "github.com/datadog/go-python3"
	"github.com/soulteary/go-python-ast/internal/rpc"
	"github.com/soulteary/go-python-ast/internal/web"
)

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		log.Fatalln("Failed to initialize Python environment")
	}

	go rpc.Launch()
	web.Launch()
}
