package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"momssi-apig-app/config"
	"momssi-apig-app/internal/utils"
	"strings"
)

type MySqlClient struct {
	cfg config.Mysql
	db  *sqlx.DB
}

func NewMysqlClient(cfg config.Mysql) (*MySqlClient, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	db, err := sqlx.Connect(cfg.Driver, url)
	if err != nil {
		return nil, fmt.Errorf("failed connect mysql, err : %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping mysql, err : %w", err)
	}

	client := &MySqlClient{
		cfg: cfg,
		db:  db,
	}

	if err := client.checkDefaultTable(); err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MySqlClient) checkDefaultTable() error {

	query := checkExistChatQuery()

	var count int
	if err := m.db.QueryRow(query).Scan(&count); err != nil {
		log.Fatal("Error checking table existence: ", err)
	}

	if count > 0 {
		return nil
	}

	content, err := utils.ReadFileContent("/internal/database/schema/database.sql")
	if err != nil {
		log.Fatalf("error opening SQL file: %v", err)
	}

	queries := strings.Split(content, ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err = m.db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query : %s, err : %w", query, err)
		}
	}

	return nil
}

func (m *MySqlClient) ExecSingleResultQuery(src interface{}, query string, args ...interface{}) error {
	if src == nil {
		return errors.New("result is empty")
	}
	if err := m.db.Get(src, query, args...); err != nil {
		return fmt.Errorf("failed to execute single result query: %w", err)
	}
	return nil
}

func (m *MySqlClient) ExecQuery(query string, args ...interface{}) (int64, error) {
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return insertID, nil
}

func checkExistChatQuery() string {
	return `
    SELECT COUNT(*)
    FROM information_schema.tables
    WHERE table_schema = 'momssi' AND table_name = 'member';
    `
}
