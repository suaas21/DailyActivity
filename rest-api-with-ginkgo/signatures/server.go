package signatures

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func NewServer(db *DB) *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/signatures", func(w http.ResponseWriter, r *http.Request) {
		data, err := fetchAllSignatures(r, db)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(data)
	})

	m.Post("/signature", func(w http.ResponseWriter, r *http.Request) {
		sigData := Signature{}
		err := json.NewDecoder(r.Body).Decode(&sigData)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(w)
			return
		}

		_, err = db.MongoDB.Collection("signatures").InsertOne(r.Context(), sigData)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(w)
			return
		}

		w.WriteHeader(http.StatusOK) // HTTP 200
		_ = json.NewEncoder(w).Encode(sigData)
	})
	return m
}
