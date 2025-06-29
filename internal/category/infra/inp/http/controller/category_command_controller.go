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
		utils.EncodeJSON(w, http.StatusMethodNotAllowed, utils.DefaultErrorResult())
		return
	}

	var req request.CategoryCreateRequest
	if err := utils.DecodeJSON(r, &req); err != nil { 
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	category, err := req.ToDomainEntity()

	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	} 


	ctx := r.Context() 
	err = h.CategoryCommandService.HandleCreate(ctx, *category)

	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	utils.EncodeJSON(w, http.StatusOK, utils.DefaultSuccessResult())
}