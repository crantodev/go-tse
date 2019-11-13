package person

type Person struct {
	ID         string `db:"id" json:"id"`
	IDProvince string `db:"id_province" json:"id_province"`
	IDDistrict string `db:"id_district" json:"id_district"`
	IDCanton   string `db:"id_canton" json:"id_canton"`
	IDLocation string `db:"id_location" json:"id_location"`
	Gender     string `db:"gender" json:"gender"`
	DueDate    string `db:"duedate" json:"duedate"`
	Vote       string `db:"vote" json:"vote"`
	FirstName  string `db:"first_name" json:"first_name"`
	MiddleName string `db:"middle_name" json:"middle_name"`
	LastName   string `db:"last_name" json:"last_name"`
}
