package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	FromDate primitive.DateTime `json:"fromDate" query:"fromDate" bson:"fromDate"`
	ToDate   primitive.DateTime `json:"toDate" query:"toDate" bson:"toDate"`
	UserID   primitive.ObjectID `json:"userID" query:"userID" bson:"userID"`
	Sort     string             `json:"sort" bson:"sort"`
	Page     uint16             `json:"page"`
	Limit    uint16             `json:"limit"`
}

//TagFilter filter
type TagFilter struct {
	BaseFilter ",inline"
	Key        string  `json:"key" query:"key" bson:"key"`
	Value      string  `json:"value" query:"value" bson:"value"`
	Min        float64 `json:"min" query:"min" bson:"min"`
	Max        float64 `json:"max" query:"max" bson:"max"`
}

//AdvertiseFilter
type AdvertiseFilter struct {
	BaseFilter ",inline"
	Tags       []*Tag
}
