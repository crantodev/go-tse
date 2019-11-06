package person

type Person struct {
	ID         string `db:"id" json:"id"`
	Location   string `db:"location" json:"location"`
	Gender     string `db:"gender" json:"gender"`
	DueDate    string `db:"duedate" json:"duedate"`
	Vote       string `db:"vote" json:"vote"`
	FirstName  string `db:"first_name" json:"first_name"`
	MiddleName string `db:"middle_name" json:"middle_name"`
	LastName   string `db:"last_name" json:"last_name"`
}
