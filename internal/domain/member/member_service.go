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

	if err := us.isDuplicatedId(req.Username); err != nil {
		return 0, err
	}

	memberInfo := NewMemberInfo(req)
	return us.repo.SignUp(memberInfo)
}

func (us *MemberService) isDuplicatedId(username string) error {
	isExist, err := us.repo.isExistByUsername(username)
	if err != nil {
		return fmt.Errorf("failed get username, err : %w", err)
	}
	if isExist {
		return form.GetCustomErr(form.ErrDuplicatedUsername)
	}

	return nil
}

func (us *MemberService) login(id, password string) MemberInfo {
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
