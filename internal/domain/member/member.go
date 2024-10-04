package member

import (
	"github.com/golang-jwt/jwt/v5"
	"momssi-apig-app/internal/domain/member/types"
	"time"
)

var JWTKey = []byte("your_secret_key")

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	AdminYn  string `json:"admin_yn"`
}

type SignUpRes struct {
	MemberId int64 `json:"member_id"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UpdateRequest struct {
}

type MemberInfo struct {
	ID        int64              `json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Nickname  string             `json:"nickname"`
	AdminYn   string             `json:"admin_yn"`
	Status    types.MemberStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func NewMemberInfo(req SignUpRequest) *MemberInfo {
	return &MemberInfo{
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  req.Nickname,
		AdminYn:   req.AdminYn,
		Status:    types.ACTIVE,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Service interface {
	SignUp(request SignUpRequest) (int64, error)
	login(id, password string) MemberInfo
	isDuplicatedId(username string) error
	getUserInfo(username string) MemberInfo
	updateUserInfo(request *UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
	isExistByUsername(username string) (bool, error)
	SignUp(data *MemberInfo) (int64, error)
}
