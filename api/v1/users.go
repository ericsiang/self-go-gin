package v1

import (
	"api/common/common_msg_id"
	"api/handler"
	"api/initialize"
	"api/model"
	"api/util/bcryptEncap"
	"api/util/gin_response"
	"api/util/jwt_secret"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func validCheckAndTrans(context *gin.Context, err error) (ok bool) {
	validTrans, ok := initialize.ErrorValidateCheckAndTrans(err)
	if ok {
		var validTransMsgData = make(map[int]string)
		var i int
		for _, v := range validTrans  {
			validTransMsgData[i] = v
			i++
		}

		gin_response.ErrorResponse(context, http.StatusBadRequest, "validate fail", common_msg_id.Fail, validTransMsgData)
		return ok
	}
	return false
}

// @Summary  Create Users
// @Description Create Users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body swagger_docs.DocUsersCreate true "request body"
// @Success 200 {string} json "{"msg": {"success": "success"},"data": {}}"
// @Failure 400 {string} json "{"msg": {"fail": "帳密錯誤"},"data": null}"
// @Router /api/v1/auth/users [post]
func CreateUser(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type responseData struct {
		UsersId uint   `json:"id"`
		Account string `json:"account"`
	}

	var data receiveData
	var newUsers model.Users
	var respData responseData

	if err := context.ShouldBindJSON(&data); err != nil {
		check := validCheckAndTrans(context, err)
		if check {
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() BindJSON fail", err, http.StatusBadRequest, nil)
		return
	}

	//密碼加密
	bcryptPassword, err := bcryptEncap.GenerateFromPassword(data.Password)
	if err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() bcrypt fail", err, http.StatusBadRequest, nil)
		return
	}
	newUsers.Account = data.Account
	newUsers.Password = string(bcryptPassword)
	users, err := newUsers.CreateUser()
	if err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() model CreateUser() fail", err, http.StatusBadRequest, nil)
		return
	}
	respData = responseData{
		UsersId: users.ID,
		Account: users.Account,
	}
	gin_response.SuccessResponse(context, http.StatusOK, "", respData, common_msg_id.Success)

}

// @Summary  User Login
// @Description User Login
// @Tags Users
// @Accept json
// @Produce json
// @Param request body swagger_docs.DocUsersLogin true "request body"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/users/login [post]
func UserLogin(context *gin.Context) {
	type ReceiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data ReceiveData
	var users model.Users
	if err := context.ShouldBindJSON(&data); err != nil {
		check := validCheckAndTrans(context, err)
		if check {
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() User 建立失敗(BindJSON fail)", err, http.StatusBadRequest, nil)
		return
	}

	user, err := users.GetUsersByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		handler.HandlerError(context, common_msg_id.Fail, "UserLogin() model GetUsersByAccount fail", err, http.StatusBadRequest, nil)
		return
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "UserLogin() User CompareHashAndPassword fail", err, http.StatusBadRequest, gin_response.CreateMsgData("fail", "帳密錯誤"))
		return
	}

	jwtToken, err := jwt_secret.GenerateToken(user.ID)
	if err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "UserLogin() User jwt GenerateToken fail", err, http.StatusBadRequest, nil)
		return
	}

	gin_response.SuccessResponse(context, http.StatusOK, "", gin_response.CreateMsgData("jwt_token", jwtToken), common_msg_id.Success)

}

// @Summary Get Users By ID
// @Description Get Users By ID
// @Tags Users
// @Accept json
// @Produce json
// @Security JwtTokenAuth
// @Param filterUsersId path string true "filterUsersId"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/auth/users/{filterUsersId} [get]
func GetUsersById(context *gin.Context) {
	type receiveData struct {
		FilterUsersId string `form:"filterUsersId" json:"filterUsersId" binding:"required"`
	}

	var data receiveData
	usersId, _ := context.Get("usersId")
	zap.S().Info("usersId :", usersId)
	data.FilterUsersId = context.Param("filterUsersId")
	zap.S().Info("filterUsersId :", data)

	gin_response.SuccessResponse(context, http.StatusOK, "", nil, common_msg_id.Success)

}
