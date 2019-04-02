package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Convention struct {
	Id                    bson.ObjectId `json:"id" bson:"_id"`
	ConventionName        string        `json:"ConventionName" bson:"ConventionName"`
	ConventionDescription string        `json:"ConventionDescription" bson:"ConventionDescription"`
	// DocumentId            bson.ObjectId `json:"LeaveEndDate" bson:"LeaveEndDate"`
	// AddedBy               bson.ObjectId `json:"ApplicationDate" bson:"ApplicationDate"`

}

func (c Convention) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
