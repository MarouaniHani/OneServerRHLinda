package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id               bson.ObjectId `json:"id" bson:"_id"`
	EventName        string        `json:"EventName" bson:"EventName"`
	EventDescription string        `json:"EventDescription" bson:"EventDescription"`
	EventStartDate   string        `json:"EventStartDate" bson:"EventStartDate"`
	// EmployeeId       bson.ObjectId `json:"EmployeeId" bson:"EmployeeId"`
}

func (e Event) String() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
