package controller

import (
	"golangapi/helper"
	"golangapi/model/web"
	"golangapi/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImplementation struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryservice service.CategoryService) CategoryController {
	return &CategoryControllerImplementation{
		CategoryService: categoryservice,
	}
}

func (controller *CategoryControllerImplementation) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreate := web.CategoryCreate{}
	helper.ReadFromRequestBody(r, &categoryCreate)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreate)
	webResponse := web.WebRespons{ //response web
		Code:   200,
		Status: "OK",
		Data:   categoryResponse, //mengambil datanya
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdate := web.CategoryUpdate{}
	helper.ReadFromRequestBody(r, &categoryUpdate)

	categoryId := p.ByName("categoryId") //konversi dulu
	id, err := strconv.Atoi(categoryId)  //konversion
	helper.PanicIfError(err)

	categoryUpdate.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdate)
	webResponse := web.WebRespons{ //response web
		Code:   200,
		Status: "OK",
		Data:   categoryResponse, //mengambil datanya
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId") //konversi dulu
	id, err := strconv.Atoi(categoryId)  //konversion
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebRespons{ //response web
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId") //konversi dulu
	id, err := strconv.Atoi(categoryId)  //konversion
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebRespons{ //response web
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebRespons{ //response web
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
