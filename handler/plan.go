package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/service"
	"github.com/train-do/Golang-Generics/utils"
)

type HandlerTourPlan struct {
	Service *service.ServiceTourPlan
}

func NewRepoTourPlan(service *service.ServiceTourPlan) *HandlerTourPlan {
	return &HandlerTourPlan{service}
}
func (h *HandlerTourPlan) GetByDestinationId(w http.ResponseWriter, r *http.Request) {
	destinationId := utils.ToInt(chi.URLParam(r, "destinationId"))
	var data []model.TourPlan
	err := h.Service.GetByDestinationId(&data, destinationId)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: data}, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
