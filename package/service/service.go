package service

import (
	"reverseString/func"
	"reverseString/package/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ReverseString(str string) string {
	return _func.Reverse(str)
}

func (s *Service) CountSymbols(str string) int {
	return _func.CountSymbols(str)
}
