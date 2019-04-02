package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type AdminRequest struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	NumberOfPaper   int           `json:"NumberOfPaper" bson:"NumberOfPaper"`
	RequestReason   string        `json:"RequestReason" bson:"RequestReason"`
	ApplicationDate string        `json:"ApplicationDate" bson:"ApplicationDate"`
	RequestStatus   bool          `json:"RequestStatus" bson:"RequestStatus"`
	// RequestType     bson.ObjectId `json:"RequestType" bson:"RequestType"`
	// ApliedBy        bson.ObjectId `json:"ApliedBy" bson:"ApliedBy"`
}

func (a AdminRequest) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
