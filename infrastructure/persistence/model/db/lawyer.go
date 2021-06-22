package db

// Lawyer ...
type Lawyer struct {
	BaseModel
	ID             int
	BarAssociation string
	Name           string
	NameKana       string
	Sex            string
	OfficeName     string
	PostCode       string
	Address        string
	PhoneNumber    string
	FaxNumber      string
	OfficeKanji    string
	LawyerKanji    string
}
