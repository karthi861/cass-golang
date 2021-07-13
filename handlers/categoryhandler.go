package handlers

import (
	"encoding/json"
	"fmt"
	"inventory/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gocql/gocql"
)

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 3
	cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = "ecommerce"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra initialized")
}

func Getallcategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM Categories").Iter()
	for iter.MapScan(m) {
		categories = append(categories, models.Category{
			ID:   m["id"].(int),
			Name: m["Name"].(string),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(categories, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var NewCategory models.Category
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "invalid data")
	}
	json.Unmarshal(reqBody, &NewCategory)
	if err := Session.Query("INSERT INTO Products(id, Name) VALUES(?, ?)",
		NewCategory.ID, NewCategory.Name).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewCategory, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}
