package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(1)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(1)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(ctx *gin.Context) {
	cntx := ctx.Request.Context()
	req := request{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	cntx = context.WithValue(cntx, "limit", req.Pagination.Page)
	cntx = context.WithValue(cntx, "offset", req.Pagination.MaxPageSize)
	interests := req.Filter.ToSliceInterest()
	userID := uuid.MustParse(req.Filter.UserID)
	blogID := uuid.MustParse(req.Filter.BlogID)

	subs, err := h.svc.Get(cntx, userID, blogID, interests)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, subs)
}
