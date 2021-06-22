package db

// Office ...
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
