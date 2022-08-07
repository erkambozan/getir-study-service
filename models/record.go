package models

type Record struct {
    Key          string  `bson:"key"`
    CreatedAt    string  `bson:"createdAt"`
	TotalCount   int     `bson:"totalCount"`
}