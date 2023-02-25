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
	ID    primitive.ObjectID
	PName string `json:"pname"`
	DOB   string `json:"dob"`
	Owner string `json:"owner"`
	Kind  Kind   `json:"kind"`
	Size  Size   `json:"size"`
	Toy   string `json:"toy"`
}

type Size struct {
	Height string `json:"toy"`
	Weight string `json:"weight"`
}

type Kind struct {
	Animal string `json:"animal"`
	Breed  string `json:"breed"`
}
