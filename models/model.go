/*
models: a struct that reflects the data obj(s) serialized
	and deserialized to/from the db layer
*/

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*


`json:" "` are known as as struct tag, allow us to attach

meta-information to correspongding struct properties.
We use them to reformat the JSON response returned by the API
*/

type Pet struct {
	// primitive.ObjectID inform the MongoDB server that it's a MongoDB ObjectIds
	ID    primitive.ObjectID `json:"id"`
	PName string             `json:"pname"`
	DOB   float64            `json:"dob"`
	Owner string             `json:"owner"`
	Kind  string             `json:"kind"`
	Size  int                `json:"size"`
	Toy   string             `json:"toy"`
}
