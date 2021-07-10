package handlers

import (
	"encoding/json"
	"fmt"
	"inventory/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

var Session *gocql.Session

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

func GetallProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM Products").Iter()
	for iter.MapScan(m) {
		products = append(products, models.Product{
			ID:          m["id"].(int),
			Name:        m["name"].(string),
			Description: m["description"].(string),
			Category:    m["category"].(string),
			Price:       m["price"].(float32),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(products, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var NewProduct models.Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "invalid data")
	}
	json.Unmarshal(reqBody, &NewProduct)
	if err := Session.Query("INSERT INTO Products(id, name, description, category,price) VALUES(?, ?, ?, ?, ?)",
		NewProduct.ID, NewProduct.Name, NewProduct.Description, NewProduct.Category, NewProduct.Price).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewProduct, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	var products []models.Product
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM Products WHERE id=?", ProductID).Iter()
	for iter.MapScan(m) {
		products = append(products, models.Product{
			ID:          m["id"].(int),
			Name:        m["name"].(string),
			Description: m["description"].(string),
			Category:    m["category"].(string),
			Price:       m["price"].(float32)})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(products, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	var UpdateProduct models.Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, " enter data")
	}
	json.Unmarshal(reqBody, &UpdateProduct)
	if err := Session.Query("UPDATE Products SET name = ?, description = ?, category = ?, price = ? WHERE id = ?",
		UpdateProduct.Name, UpdateProduct.Description, UpdateProduct.Category, UpdateProduct.Price, ProductID).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "updated successfully")

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM Products WHERE id = ?", ProductID).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted  the product num %s ", ProductID)
}
