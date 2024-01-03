package repo

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrorDuplicateRecord = fmt.Errorf("duplicate record")
	ErrorRecordNotFound  = fmt.Errorf("record not found")
)

type RepoServices struct {
	User UserServices
}

func New(db *gorm.DB) *RepoServices {
	return &RepoServices{
		User: UserServices{db: db},
	}
}
