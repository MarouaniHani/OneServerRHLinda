package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type LeaveRequest struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	LeaveReason     string        `json:"LeaveReason" bson:"LeaveReason"`
	LeaveStartDate  string        `json:"LeaveStartDate" bson:"LeaveStartDate"`
	LeaveEndDate    string        `json:"LeaveEndDate" bson:"LeaveEndDate"`
	ApplicationDate string        `json:"ApplicationDate" bson:"ApplicationDate"`
	RequestStatus   bool          `json:"RequestStatus" bson:"RequestStatus"`
	// ApprouvedBy     bson.ObjectId `json:"RequestStatus" bson:"RequestStatus"`
	// RequestType     bson.ObjectId `json:"RequestType" bson:"RequestType"`
	// ApliedBy        bson.ObjectId `json:"ApliedBy" bson:"ApliedBy"`
}

func (l LeaveRequest) String() string {
	b, err := json.Marshal(l)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
