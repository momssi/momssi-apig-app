package member

import (
	"fmt"
	"momssi-apig-app/api/form"
)

type MemberService struct {
	repo Repository
}

func NewMemberService(repo Repository) Service {
	return &MemberService{
		repo: repo,
	}
}

func (us *MemberService) SignUp(req SignUpRequest) (int64, error) {

	if err := us.isDuplicatedId(req.Email); err != nil {
		return 0, err
	}

	memberInfo := NewMemberInfo(req)

	if err := memberInfo.hashPassword(); err != nil {
		return 0, fmt.Errorf("failed encoding password, err : %w", err)
	}

	return us.repo.Save(memberInfo)
}

func (us *MemberService) isDuplicatedId(email string) error {
	isExist, err := us.repo.isExistByUsername(email)
	if err != nil {
		return fmt.Errorf("failed get email, err : %w", err)
	}
	if isExist {
		return form.GetCustomErr(form.ErrDuplicatedEmail)
	}

	return nil
}

func (us *MemberService) Login(id, password string) MemberInfo {
	//TODO implement me
	panic("implement me")
}

func (us *MemberService) LoginSuccess(email, refreshToken string) error {
	//TODO implement me
	panic("implement me")
}

func (us *MemberService) getUserInfo(username string) MemberInfo {
	//TODO implement me
	panic("implement me")
}

func (us *MemberService) updateUserInfo(request *UpdateRequest) int64 {
	//TODO implement me
	panic("implement me")
}

func (us *MemberService) deleteByUsername(username string) int64 {
	//TODO implement me
	panic("implement me")
}
