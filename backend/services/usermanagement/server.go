package main

import (
	"github.com/subratohld/quiz/usermanagement/internal/svc"
)

func main() {
	service := new(svc.Service)
	service.Init()
	service.Start()
}
