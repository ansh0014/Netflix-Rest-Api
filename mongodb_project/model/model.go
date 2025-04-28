package model

type Netflix struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Watched bool   `json:"watched,omitempty" bson:"watched,omitempty"`
}
