package signatures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DailyActivity/rest-api-with-ginkgo/signatures"
	"github.com/modocache/gory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Convert JSON data into a slice.
*/
func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

var _ = Describe("Server", func() {
	var db *signatures.DB
	var apiClient http.Client
	var client *mongo.Client
	var baseURI string
	var req *http.Request
	var err error

	BeforeEach(func() {
		// Initialize api client with timeout
		apiClient = http.Client{Timeout: time.Minute * 2}
		// Base URI where the api server running
		baseURI = "http://localhost:3000"
		// Initialize mongoDB and  client
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err = mongo.Connect(context.TODO(), clientOptions)
		Expect(err).NotTo(HaveOccurred())
		db = &signatures.DB{}
		db.DBClient = client
		db.MongoDB = client.Database("test_signatures")
		db.Ctx = context.Background()
	})

	AfterEach(func() {
		// Clear the database after each ginkgo_test_book.
		err = db.MongoDB.Drop(db.Ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("GET /signatures", func() {

		// Set up a new GET req before every ginkgo_test_book
		// in this describe block.
		BeforeEach(func() {
			req, err = http.NewRequest("GET", baseURI+"/signatures", nil)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when no signatures exist", func() {
			It("returns a status code of 200", func() {
				res, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))
			})

			It("returns a null body", func() {
				resp, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				err = resp.Body.Close()
				Expect(err).NotTo(HaveOccurred())

				var sigData signatures.Signature
				err = json.Unmarshal(body, &sigData)
				Expect(sigData).To(Equal(signatures.Signature{}))
			})
		})

		Context("when signatures exist", func() {
			// Insert two valid signatures into the database
			// before each ginkgo_test_book in this context.
			BeforeEach(func() {
				_, err := db.MongoDB.Collection("signatures").InsertOne(db.Ctx, gory.Build("signature"))
				Expect(err).NotTo(HaveOccurred())
				_, err = db.MongoDB.Collection("signatures").InsertOne(db.Ctx, gory.Build("signature"))
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns a status code of 200", func() {
				res, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))
			})

			It("returns those signatures in the body", func() {
				resp, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				err = resp.Body.Close()
				Expect(err).NotTo(HaveOccurred())

				peopleJSON := sliceFromJSON(body)
				Expect(len(peopleJSON)).To(Equal(2))

				personJSON := peopleJSON[0].(map[string]interface{})
				Expect(personJSON["first_name"]).To(Equal("Jane"))
				Expect(personJSON["last_name"]).To(Equal("Doe"))
				Expect(personJSON["age"]).To(Equal(float64(27)))
				Expect(personJSON["message"]).To(Equal("I agree!"))
				Expect(personJSON["email"]).To(
					ContainSubstring("jane-doe"))
			})
		})
	})

	Describe("POST /signature", func() {
		Context("with valid JSON", func() {
			// Create a POST req with valid JSON from
			// our factory before each ginkgo_test_book in this context.
			BeforeEach(func() {
				body, err := json.Marshal(
					gory.Build("signature"))
				Expect(err).NotTo(HaveOccurred())
				req, err = http.NewRequest(
					"POST", baseURI+"/signature", bytes.NewReader(body))
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns a status code of 200", func() {
				res, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))
			})

			It("returns the inserted signature", func() {
				resp, err := apiClient.Do(req)
				Expect(err).NotTo(HaveOccurred())

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				err = resp.Body.Close()
				Expect(err).NotTo(HaveOccurred())

				personJSON := mapFromJSON(body)
				Expect(personJSON["first_name"]).To(Equal("Jane"))
				Expect(personJSON["last_name"]).To(Equal("Doe"))
				Expect(personJSON["age"]).To(Equal(float64(27)))
				Expect(personJSON["message"]).To(Equal("I agree!"))
				Expect(personJSON["email"]).To(
					ContainSubstring("jane-doe"))
			})
		})
	})
})
