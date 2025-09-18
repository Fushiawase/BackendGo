package repo

import (
	"BackendGo/core/document"
	"BackendGo/core/errors"
	"fmt"

	"gorm.io/gorm"
)

type sqliteDocumentRepository struct {
	db *gorm.DB
}

func NewSqliteRepo(db *gorm.DB) document.Repository {
	return &sqliteDocumentRepository{db}
}

var _ document.Repository = &sqliteDocumentRepository{}

func (r *sqliteDocumentRepository) Create(d *document.Document) error {
	dRow := &DocumentRow{Name: d.Name, Description: d.Description}
	if err := r.db.Create(dRow).Error; err != nil {
		return err
	}
	d.Id = dRow.Id
	return nil
}

func (r *sqliteDocumentRepository) FindByID(id int) (*document.Document, error) {
	var dRow DocumentRow
	if err := r.db.First(&dRow, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.RecordNotFoundErr{Message: fmt.Sprintf("Couldn't find document of id: %d", id)}
		}
		return nil, err
	}
	d := &document.Document{Id: dRow.Id, Name: dRow.Name, Description: dRow.Description}
	return d, nil
}

func (r *sqliteDocumentRepository) Delete(id int) error {
	res := r.db.Delete(DocumentRow{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.RecordNotFoundErr{Message: fmt.Sprintf("Couldn't find document of id: %d", id)}
	}
	return nil
}
