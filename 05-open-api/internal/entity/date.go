package entity

import (
	"strings"
	"time"
)

const Dateformat = "2006-01-02"

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(Dateformat, value) //parse time
	if err != nil {
		return err
	}

	*d = Date(t) //set result using the pointer
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.ToTime().Format(Dateformat) + `"`), nil
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}
