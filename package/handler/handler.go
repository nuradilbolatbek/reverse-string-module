package handler

import "reverseString/package/service"

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) ReverseString(s string) string {
	return h.service.ReverseString(s)
}

func (h *Handler) CountSymbols(str string) int {
	return h.service.CountSymbols(str)
}
