package signatures

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type Signature struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	Email     string `json:"email,omitempty bson:"email"`
	Age       int    `json:"age,omitempty" bson:"age"`
	Message   string `json:"message,omitempty" bson:"message"`
}

func fetchAllSignatures(r *http.Request, db *DB) ([]Signature, error) {
	signatures := make([]Signature, 0)
	cur, err := db.MongoDB.Collection("signatures").Find(r.Context(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(r.Context(), &signatures)
	if err != nil {
		return nil, err
	}
	return signatures, nil
}
