package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type State struct {
	Id      primitive.ObjectID `json: "id, omitempty"`
	Name    string             `json: "state, omitempty" validate: "required"`
	Capital string             `json: "capital, omitempty" validate: "required"`
}
