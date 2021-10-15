package handler

import (
	"datatables_go/helper"
	"datatables_go/users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service        users.Service
	userRepository users.Repository
}

func NewUserHandler(service users.Service, userRepository users.Repository) *userHandler {
	return &userHandler{service, userRepository}
}

func (h *userHandler) GetAll(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	var d users.DTJson

	err = json.Unmarshal(body, &d)

	if err != nil {
		panic(err)
	}

	fetchUser, err := h.service.GetAllUser(d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	totalData, err := h.userRepository.GetTotalUser(d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	response := helper.ApiResponse(d.Draw, totalData, users.FormatUsers(fetchUser))
	c.JSON(http.StatusOK, response)
}
