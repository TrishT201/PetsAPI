package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pet struct {
	ID    primitive.ObjectID `json:"id"`
	PName string             `json:"pname"`
	DOB   float64            `json:"dob"`
	Owner string             `json:"owner"`
	Kind  string             `json:"kind"`
	Size  int                `json:"size"`
	Toy   string             `json:"toy"`
}
