package request

type CreateAdminRequest struct {
    Account  string `json:"account" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type AdminLoginRequest struct {
    Account  string `json:"account" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type GetAdminsByIDRequest struct {
    FilterAdminsId string `form:"filterAdminsId" json:"filterAdminsId" binding:"required"`
}
