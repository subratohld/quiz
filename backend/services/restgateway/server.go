package main

import (
	"github.com/subratohld/quiz/restgateway/internal/svc"
)

func main() {
	service := new(svc.Service)
	service.Init()
	service.Start()
}
