package main

import (
	"github.com/subratohld/quiz/questionbank/internal/svc"
)

func main() {
	service := new(svc.Service)
	service.Init()
	service.Start()
}
