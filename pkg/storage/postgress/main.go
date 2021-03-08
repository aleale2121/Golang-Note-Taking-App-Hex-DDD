package postgress

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDatabaseConnection() (*gorm.DB, error) {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:root@localhost/Note?sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return dbconn, nil
}
