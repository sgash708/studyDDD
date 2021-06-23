package repository

import (
	"github.com/sgash708/studyDDD/infrastructure/persistence/model/db"
)

// IOfficeRepository is ...
type IOfficeRepository interface {
	Insert(ofice *db.Office) error
}
