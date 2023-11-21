package v1

import (
	"api/initialize"
	"api/model"
	"api/util/gin_response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func validCheckAndTrans(context *gin.Context, err error) {
	validTrans, ok := initialize.ErrorValidateCheckAndTrans(err)
	if ok {
		gin_response.ErrorResponse(context, http.StatusBadRequest, validTrans)
		return
	}
}

func CreateUser(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var data receiveData
	var newUsers model.Users

	if err := context.ShouldBindJSON(&data); err != nil {
		validCheckAndTrans(context,err)
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.S().Error("CreateUser() User 建立失敗(BindJSON fail) :" + err.Error())
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail","BindJSON fail"))
		return
	}

	_, err := newUsers.GetUsersByAccount(data.Account)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		var Msg = make(map[string]string)
		Msg["fail"] = "db GetUsersByAccount fail"
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail","db GetUsersByAccount fail"))
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//密碼加密
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			zap.S().Error("CreateUser() User 建立失敗(bcrypt fail) :" + err.Error())
			var Msg = make(map[string]string)
			Msg["fail"] = "bcrypt GenerateFromPassword fail"
			gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
			return
		}
		newUsers.Account = data.Account
		newUsers.Password = string(bcryptPassword)
		users, err := newUsers.CreateUser()
		if err != nil {
			zap.S().Error("CreateUser() User 建立失敗(db fail) :" + err.Error())
			var Msg = make(map[string]string)
			Msg["fail"] = "db CreateUser fail"
			gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
			return
		}

		var Msg = make(map[string]string)
		Msg["success"] = "success"
		gin_response.SuccessResponse(context, http.StatusOK, Msg, users)
	} else {
		var Msg = make(map[string]string)
		Msg["fail"] = "帳號已存在"
		gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
	}

}

func UserLogin(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data receiveData
	var users model.Users
	if err := context.ShouldBindJSON(&data); err != nil {
		validCheckAndTrans(context,err)

		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.S().Error("CreateUser() User 建立失敗(BindJSON fail) :" + err.Error())
		var Msg = make(map[string]string)
		Msg["User 建立失敗"] = "BindJSON fail"
		gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
		return
	}

	user, err := users.GetUsersByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		var Msg = make(map[string]string)
		Msg["User Login 失敗"] = "db GetUsersByAccount fail"
		gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
	}
	zap.S().Info("UserLogin() User :", user.Password)
	zap.S().Info("UserLogin() receiveData :", data.Password)
	//密碼驗證
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		zap.S().Error("UserLogin() User bcrypt:", err.Error())
		var Msg = make(map[string]string)
		Msg["User Login 失敗"] = "帳密錯誤"
		gin_response.ErrorResponse(context, http.StatusBadRequest, Msg)
		return
	}

	var Msg = make(map[string]string)
	Msg["User Login 成功"] = "success"
	gin_response.SuccessResponse(context, http.StatusOK, Msg, user)

}

func GetUsers(context *gin.Context) {
	type receiveData struct {
	}

	// var data receiveData
	// var users model.Users

}

func GetUsersById(context *gin.Context) {
	type receiveData struct {
		FilterUsersId int64 `form:"filterUsersId"`
	}

	var data receiveData

	if err := context.ShouldBindQuery(&data); err != nil {

		return
	} else {
		var users model.Users
		users.GetUsersById(data.FilterUsersId)

	}
}
