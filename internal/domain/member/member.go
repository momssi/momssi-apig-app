package member

import (
	"golang.org/x/crypto/bcrypt"
	"momssi-apig-app/api/form"
	"momssi-apig-app/internal/domain/member/types"
	"time"
)

type MemberInfo struct {
	ID             int64     `json:"id" db:"id"`                             // MySQL 칼럼: id
	Email          string    `json:"email" db:"email"`                       // MySQL 칼럼: email
	Password       string    `json:"password" db:"password"`                 // MySQL 칼럼: password
	Name           string    `json:"name" db:"name"`                         // MySQL 칼럼: name
	AdminYn        string    `json:"admin_yn" db:"admin_yn"`                 // MySQL 칼럼: admin_yn
	DeleteYn       string    `json:"delete_yn" db:"delete_yn"`               // MySQL 칼럼: delete_yn
	LastLoginIP    string    `json:"last_login_ip" db:"last_login_ip"`       // MySQL 칼럼: last_login_ip
	RefreshToken   string    `json:"refresh_token" db:"refresh_token"`       // MySQL 칼럼: refresh_token
	LoginFailCount int       `json:"login_fail_count" db:"login_fail_count"` // MySQL 칼럼: login_fail_count
	Status         string    `json:"status" db:"status"`                     // MySQL 칼럼: status
	LastLoginAt    time.Time `json:"last_login_at" db:"last_login_at"`       // MySQL 칼럼: last_login_at
	CreatedAt      time.Time `json:"created_at" db:"created_at"`             // MySQL 칼럼: created_at
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`             // MySQL 칼럼: updated_at
}

func NewMemberInfo(req form.SignUpRequest) *MemberInfo {
	return &MemberInfo{
		Email:     req.Email,
		Password:  req.Password,
		Name:      req.Name,
		Status:    types.ACTIVE.String(),
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

func (m *MemberInfo) checkPassword(inputPw string) error {
	return bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(inputPw))
}

type Service interface {
	SignUp(request form.SignUpRequest) (int64, error)
	Login(id, password string) (MemberInfo, error)
	LoginSuccess(loginIP, email, refreshToken string) error
	isDuplicatedId(username string) error
	getUserInfo(username string) MemberInfo
	updateUserInfo(request *form.UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
	IsExistByEmail(email string) (bool, error)
	Save(data *MemberInfo) (int64, error)
	FindMemberByEmail(email string) (MemberInfo, error)
	UpdateLoginInfo(loginIp, email, refreshToken string) error
}
