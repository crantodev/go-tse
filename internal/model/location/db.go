package location

import (
	"fmt"
	"regexp"
)

type Location struct {
	ID         string `db:"id_location" json:"id_location"`
	IDProvince string `db:"id_province" json:"id_province"`
	IDDistrict string `db:"id_district" json:"id_district"`
	IDCanton   string `db:"id_canton" json:"id_canton"`
	Province   string `db:"province" json:"province"`
	District   string `db:"district" json:"district"`
	Canton     string `db:"canton" json:"canton"`
}

func (l *Location) Save() error {
	return nil
}

func (l *Location) IDSplitter() error {
	re := regexp.MustCompile(`^(\d{1})(\d{2})(\d{3})$`)
	matched := re.FindStringSubmatch(l.ID)
	fmt.Printf("Whole: %s, Province: %s, District: %s, Canton: %s", matched[0], matched[1], matched[2], matched[3])

	return nil
}
