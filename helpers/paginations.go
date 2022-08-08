package helpers

import (
	"strconv"

	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/gin-gonic/gin"
)

func CreatePagination(context *gin.Context) *dto.PaginationDTO {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	return &dto.PaginationDTO{Limit: limit, Page: page}
}
