package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type ContractType struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	ContractType string        `json:"ContractType" bson:"ContractType"`
}

func (c ContractType) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
