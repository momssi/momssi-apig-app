package member

import (
	"momssi-apig-app/internal/database"
	"strings"
)

type MemberRepository struct {
	db *database.MySqlClient
}

func NewMemberRepository(db *database.MySqlClient) Repository {
	return &MemberRepository{db: db}
}

func (mr *MemberRepository) IsExistByEmail(email string) (bool, error) {
	qs := query([]string{
		"SELECT",
		"COUNT(1)",
		"FROM momssi.member m",
		"WHERE 1=1",
		"AND m.email = ?",
		"AND m.status = 'active'",
	})

	var count int
	if err := mr.db.ExecSingleResultQuery(&count, qs, email); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (mr *MemberRepository) FindMemberByEmail(email string) (MemberInfo, error) {
	qs := query([]string{
		"SELECT",
		"*",
		"FROM momssi.member m",
		"WHERE 1=1",
		"AND m.status = 'ACTIVE'",
		"AND m.email = ?",
		"AND m.delete_yn = 'N'",
	})

	memberInfo := &MemberInfo{}
	if err := mr.db.ExecSingleResultQuery(memberInfo, qs, email); err != nil {
		return MemberInfo{}, err
	}

	return *memberInfo, nil
}

func (mr *MemberRepository) Save(data *MemberInfo) (int64, error) {
	qs := query([]string{
		"INSERT INTO",
		"momssi.member(`email`, `password`, `name`, `admin_yn`, `status`)",
		"VALUES (?, ?, ?, ?, ?)",
	})

	memberId, err := mr.db.ExecQuery(qs, data.Email, data.Password, data.Name, data.AdminYn, data.Status)
	if err != nil {
		return 0, err
	}

	return memberId, nil
}

func (mr *MemberRepository) UpdateLoginInfo(loginIp, email, refreshToken string) (int64, error) {
	qs := query([]string{
		"UPDATE momssi.member m",
		"SET",
		"last_login_at = NOW(),",
		"login_fail_count = 0,",
		"last_login_ip = ?,",
		"refresh_token = ?",
		"WHERE 1=1",
		"AND email = ?",
		"AND delete_yn = 'N'",
	})

	memberId, err := mr.db.ExecQuery(qs, loginIp, refreshToken, email)
	if err != nil {
		return 0, err
	}

	return memberId, nil
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
