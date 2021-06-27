package mssql

import (
	"database/sql"
	"fmt"

	irepo "github.com/sgash708/studyDDD/domain/repository"
	"github.com/sgash708/studyDDD/infrastructure/persistence/model/db"
)

// OfficeRepository is ...
type OfficeRepository struct {
	*sql.DB
}

// NewOfficeRepository 構造体初期化
func NewOfficeRepository(Conn *sql.DB) irepo.IOfficeRepository {
	return &OfficeRepository{Conn}
}

// Insert データ挿入
func (repo *OfficeRepository) Insert(office *db.Office) error {
	stmt, err := repo.DB.Prepare("INSERT INTO godb.offices(id, corporate_name, bar_association, name, office_bar_association, postcode, address, phone_number, fax_number, corp_kanji, office_kanji) SELECT %s, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s, %s WHERE NOT EXISTS (SELECT id FROM godb.offices WHERE id = %s);")
	if err != nil {
		return err
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			fmt.Println("ステートメント終了できません。", err)
		}
	}()
	_, err = stmt.Exec(office.ID, office.CorporateName, office.BarAssociation, office.Name, office.OfficeBarAssociation, office.PostCode, office.Address, office.PhoneNumber, office.FaxNumber, office.CorpKanji, office.OfficeKanji, office.ID)
	if err != nil {
		return err
	}
	return nil
}
