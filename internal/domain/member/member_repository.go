package member

import (
	"errors"
	"momssi-apig-app/internal/database"
	"strings"
)

type MemberRepository struct {
	db *database.MySqlClient
}

func NewMemberRepository(db *database.MySqlClient) Repository {
	return &MemberRepository{db: db}
}

func (mr *MemberRepository) isExistByUsername(username string) (bool, error) {
	qs := query([]string{
		"SELECT",
		"COUNT(1)",
		"FROM momssi.member m",
		"WHERE 1=1",
		"AND m.username = ?",
		"AND m.status = 'active'",
	})

	count, err := mr.db.ExecCountQuery(qs, username)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (mr *MemberRepository) SignUp(data *MemberInfo) (int64, error) {
	qs := query([]string{
		"INSERT INTO",
		"momssi.member(`username`, `password`, `nickname`, `admin_yn`, `status`)",
		"VALUES (?, ?, ?, ?, ?)",
	})

	if err := mr.db.ExecQuery(qs, data.Username, data.Password, data.Nickname, data.AdminYn, data.Status); err != nil {
		return 0, err
	}

	mqs := query([]string{
		"SELECT `id`",
		"FROM momssi.member m",
		"WHERE 1=1",
		"AND m.username = ?",
	})

	result, err := mr.db.ExecSingleResultQuery(mqs, data.Username)
	if err != nil {
		return 0, err
	}

	id, ok := result.(int64)
	if !ok {
		return 0, errors.New("failed convert data type")
	}

	return id, nil
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
