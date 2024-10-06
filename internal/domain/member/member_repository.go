package member

import (
	"errors"
	"fmt"
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

	count, err := mr.db.ExecCountQuery(qs, email)
	if err != nil {
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
		"AND m.email = ?",
		"AND m.delete_yn = 'N'",
	})

	result, err := mr.db.ExecSingleResultQuery(qs, email)
	if err != nil {
		return MemberInfo{}, err
	}

	memberInfo, ok := result.(MemberInfo)
	if !ok {
		return MemberInfo{}, fmt.Errorf("failed to convert result to MemberInfo, email : %s", email)
	}

	return memberInfo, nil
}

func (mr *MemberRepository) Save(data *MemberInfo) (int64, error) {
	qs := query([]string{
		"INSERT INTO",
		"momssi.member(`email`, `password`, `name`, `admin_yn`, `status`)",
		"VALUES (?, ?, ?, ?, ?)",
	})

	if err := mr.db.ExecQuery(qs, data.Email, data.Password, data.Name, data.AdminYn, data.Status); err != nil {
		return 0, err
	}

	mqs := query([]string{
		"SELECT `id`",
		"FROM momssi.member m",
		"WHERE 1=1",
		"AND m.email = ?",
	})

	result, err := mr.db.ExecSingleResultQuery(mqs, data.Email)
	if err != nil {
		return 0, err
	}

	id, ok := result.(int64)
	if !ok {
		return 0, errors.New("failed convert data type")
	}

	return id, nil
}

func (mr *MemberRepository) UpdateLoginInfo(loginIp, email, refreshToken string) error {
	qs := query([]string{
		"UPDATE momssi.member m",
		"SET",
		"last_login_at = NOW()",
		"login_fail_count = 0",
		"last_login_ip = ?",
		"refresh_token = ?",
		"WHERE 1=1",
		"email = ?",
		"delete_yn = 'N'",
	})

	if err := mr.db.ExecQuery(qs, loginIp, refreshToken, email); err != nil {
		return fmt.Errorf("failed exec query, err : %w", err)
	}

	return nil
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
