package handler

import (
	"fmt"
	"self_go_gin/gin_application/inter/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func MysqlErrorCheck(ctx *gin.Context, err error) (bool, error) {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 { // Duplicate entry detected
				ctx.JSON(http.StatusBadRequest, response.FailResponse{
					Msg: "duplicate_entry",
				})
				return true, fmt.Errorf("MysqlErrorCheck() \n %w", err)
			}
		}
	}
	return false, nil
}
