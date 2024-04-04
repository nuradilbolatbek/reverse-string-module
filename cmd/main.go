package main

import (
	"fmt"
	"reverseString/package/handler"
	"reverseString/package/repository"
	"reverseString/package/service"
)

func main() {
	str := "Hello, world!"
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	reversed := h.ReverseString(str)
	len1 := h.CountSymbols(str)

	fmt.Println("Reversed:", reversed)
	fmt.Println("length:", len1)
}
