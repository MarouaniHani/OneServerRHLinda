package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// EmployeeRole ...
type EmployeeRole struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Role string        `json:"Role" bson:"Role"`
}

func (e EmployeeRole) String() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
