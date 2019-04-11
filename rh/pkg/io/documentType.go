package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type DocumentType struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Type string        `json:"Type" bson:"Type"`
}

func (d DocumentType) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
