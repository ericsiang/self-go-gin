package service

import (
	"errors"
	"fmt"
	"self_go_gin/common/common_const"
	"self_go_gin/gin_application/handler"
	"self_go_gin/util/bcryptEncap"
	"self_go_gin/util/jwt_secret"
	"self_go_gin/domains/admin/entity/model"
	"self_go_gin/domains/admin/repository"
	"self_go_gin/gin_application/api/v1/admin/request"
	"gorm.io/gorm"
)

type AdminService struct {
	repo repository.AdminRepositoryInterface
}

func NewAdminService() *AdminService {
	return &AdminService{
		repo: repository.NewAdminRepository(),
	}
}

func (s *AdminService) CreateAdmin(req request.CreateAdminRequest) (*model.Admins, error) {
	log_data := map[string]interface{}{
		"req": req,
	}
	_, err := s.repo.GetAdminByAccount(req.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("AdminService CreateAdmin() data: %s \n %w", log_data, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//密碼加密
		bcryptPassword, err := bcryptEncap.GenerateFromPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("AdminService CreateAdmin() bcrypt fail data: %s \n %w", log_data, err)
		}

		newAdmins := model.Admins{
			Account:  req.Account,
			Password: string(bcryptPassword),
		}

		user, err := s.repo.CreateAdmin(newAdmins)
		if err != nil {
			return nil, fmt.Errorf("AdminService AdminService() data: %s \n %w", log_data, err)
		}

		return user, nil
	}

	return nil, fmt.Errorf("AdminService AdminService() resource exist data: %s \n %w", log_data, handler.ErrResourceExist)
}

func (s *AdminService) CheckLogin(req request.AdminLoginRequest) (*string, error) {
	log_data := map[string]interface{}{
		"req": req,
	}
	admin, err := s.repo.GetAdminByAccount(req.Account)
	if err != nil {
		return nil, fmt.Errorf("AdminService CheckLogin() data: %s \n %w", log_data, err)
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("AdminService CheckLogin() CompareHashAndPassword() data : %+v \n %w", log_data, err) // 密碼錯誤
	}

	jwtToken, err := jwt_secret.GenerateToken(common_const.LoginUser, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("UserService AdminService() GenerateToken() data : %+v \n %w", log_data, err) // 密碼錯誤
	}

	return &jwtToken, nil

}
