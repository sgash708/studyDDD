package db

// Office 法律事務所モデル
type Office struct {
	BaseModel
	ID                   int
	CorporateName        string
	BarAssociation       string
	Name                 string
	OfficeBarAssociation string
	PostCode             string
	Address              string
	PhoneNumber          string
	FaxNumber            string
	CorpKanji            string
	OfficeKanji          string
}
