package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// RequestType ...
type RequestType struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	RequestName     string        `json:"RequestName" bson:"RequestName"`
	RequestCategory string        `json:"RequestCategory" bson:"RequestCategory"`
}

func (r RequestType) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
