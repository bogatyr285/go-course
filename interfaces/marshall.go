package main

import (
	"encoding/json"
	"log"
	"time"
)

type DataModel struct {
	Created   time.Time `json:"created"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"pw"`
}

func (d *DataModel) MarshalJSON() ([]byte, error) {
	type Alias DataModel
	bys, err := json.Marshal(&struct {
		*Alias
		Password string `json:"pw,omitempty"`
		Created  string `json:"created"`
	}{
		Alias:    (*Alias)(d),
		Password: "*****",
		Created:  d.Created.Format(time.ANSIC),
	})
	if err != nil {
		return nil, err
	}
	return bys, nil
}

func main() {
	d := DataModel{Created: time.Now(), FirstName: "FNAMe", Password: "password"}

	b, err := json.Marshal(&d)
	if err != nil {
		panic(err)
	}
	log.Println(string(b))
}
