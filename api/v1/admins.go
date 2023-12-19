package v1

import (
	"api/util/gin_response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Get Admins By ID
// @Description Get Admins By ID
// @Tags Admins
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param filterAdminsId path string true "filterAdminsId"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/auth/admins/{filterAdminsId} [get]
func GetAdminsById(context *gin.Context) {
	type receiveData struct {
		FilterAdminsId string `form:"filterAdminsId" json:"filterAdminsId" binding:"required"`
	}

	var data receiveData
	admin_id, _ := context.Get("filterAdminsId")
	zap.S().Info("admin_id :", admin_id)
	data.FilterAdminsId = context.Param("filterAdminsId")
	zap.S().Info("FilterAdminsId :", data)
	gin_response.SuccessResponse(context, http.StatusOK, gin_response.CreateMsg("GetAdminsById test", "just test"), data)

}
