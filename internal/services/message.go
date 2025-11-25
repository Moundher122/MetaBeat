package services

import (
	"net/http"
)

func (h *Handler) SendAnonymMessage(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetAllMessages(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) MarkMessageRead(w http.ResponseWriter, r *http.Request)   {}
