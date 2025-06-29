package controller

import (
	"net/http"

	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/infra/inp/http/request"
	"github.com/bilgehanaygn/urun/internal/common/utils"
)

type CategoryCommandController struct {
	CategoryCommandService application.CategoryCommandService
}


func (h *CategoryCommandController) HandleCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.EncodeJSON(w, http.StatusMethodNotAllowed, utils.DefaultErrorResult())
		return
	}

	var req request.CategoryCreateRequest
	if err := utils.DecodeJSON(r, &req); err != nil { 
		w.WriteHeader(http.StatusBadRequest)
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
	}

	category, err := req.ToDomainEntity()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
	} 


	ctx := r.Context() 
	err = h.CategoryCommandService.HandleCreate(ctx, *category)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJSON(w, http.StatusOK, utils.DefaultSuccessResult())
}