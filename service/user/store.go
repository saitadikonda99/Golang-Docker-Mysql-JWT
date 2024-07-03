package user

import (
	"database/sql"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/types"
)


type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

func(s *store) GetUserByMail(mail string) (*types.User, error) {
	
	rows, err := s.db.Query("SELECT * FROM users WHERE mail = ?", mail)

	if err != nil {
		return nil, err
	}
	
	user := new (types.User)

	for rows.Next() {
		user, err = ScanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, nil
	}

	return user, nil
}

func ScanRow(row *sql.Rows) (*types.User, error) {

	user := new(types.User)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,	
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}