package controller

import (
	"net/http"
	"strings"

	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/infra/inp/http/request"
	"github.com/bilgehanaygn/urun/internal/category/infra/inp/http/response"
	"github.com/bilgehanaygn/urun/internal/common/utils"
	"github.com/google/uuid"
)

type CategoryController struct {
	CategoryService application.CategoryService
}


func (h *CategoryController) HandleCreate(w http.ResponseWriter, r *http.Request) {

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
	err = h.CategoryService.HandleCreate(ctx, *category)

	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	utils.EncodeJSON(w, http.StatusOK, utils.DefaultSuccessResult())
}

func (h *CategoryController) HandleUpdate(w http.ResponseWriter, r *http.Request) {

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
	err = h.CategoryService.HandleUpdate(ctx, *category)

	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	utils.EncodeJSON(w, http.StatusOK, utils.DefaultSuccessResult())
}

func (h *CategoryController) HandleGetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.EncodeJSON(w, http.StatusMethodNotAllowed, utils.DefaultErrorResult())
		return
	}

	idParam := strings.TrimPrefix(r.URL.Path, "/category/") 
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	ctx := r.Context()
	category, err := h.CategoryService.HandleGetById(ctx, id)
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	utils.EncodeJSON(w, http.StatusOK, response.NewCategoryDetailDTO(category))
}

func (h *CategoryController) HandleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.EncodeJSON(w, http.StatusMethodNotAllowed, utils.DefaultErrorResult())
		return
	}

	ctx := r.Context()
	categories, err := h.CategoryService.HandleList(ctx)
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, utils.DefaultErrorResult())
		return
	}

	dtos := make([]response.CategoryDetailDto, len(categories))
	for i, c := range categories {
		dtos[i] = response.NewCategoryDetailDTO(c)
	}

	utils.EncodeJSON(w, http.StatusOK, dtos)
}