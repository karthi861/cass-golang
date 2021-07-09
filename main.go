package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func main() {
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
