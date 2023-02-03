package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create(ctx *gin.Context) {
	req := request{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cntx := ctx.Request.Context()
	fmt.Println(req)
	interests := req.Filter.ToSliceInterest()
	if err := h.svc.Create(cntx, interests); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Subscription created successfully")
}
