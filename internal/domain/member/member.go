package member

import (
	"golang.org/x/crypto/bcrypt"
	"momssi-apig-app/internal/domain/member/types"
	"time"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	AdminYn  string `json:"admin_yn"`
}

type SignUpRes struct {
	MemberId int64 `json:"member_id"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateRequest struct {
}

type MemberInfo struct {
	ID        int64              `json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Name      string             `json:"name"`
	AdminYn   string             `json:"admin_yn"`
	Status    types.MemberStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func NewMemberInfo(req SignUpRequest) *MemberInfo {
	return &MemberInfo{
		Email:     req.Email,
		Password:  req.Password,
		Name:      req.Name,
		AdminYn:   req.AdminYn,
		Status:    types.ACTIVE,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *MemberInfo) hashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashedPassword)
	return nil
}

func (m *MemberInfo) checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type Service interface {
	SignUp(request SignUpRequest) (int64, error)
	Login(id, password string) MemberInfo
	LoginSuccess(email, refreshToken string) error
	isDuplicatedId(username string) error
	getUserInfo(username string) MemberInfo
	updateUserInfo(request *UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
	isExistByUsername(username string) (bool, error)
	Save(data *MemberInfo) (int64, error)
}
