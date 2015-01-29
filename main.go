// random project main.go
package main

//package stat

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"time"
)

type TestData struct {
	_id   bson.ObjectId `bson:"_id,omitempty"`
	Value int           `bson:"Value"`
	DT    time.Time     `bson:"DT"`
}

var (
	IsDrop = true
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func main() {
	var response int
	fmt.Println("Please enter the number of records you want to generate\n")

	_, err := fmt.Scanf("%d", &response)

	if response < 1 {
		fmt.Println("not a valid number, will default to 50")
		response = 50
	} else {
		fmt.Println(response, " records will be generated")

	}

	howMany := response

	session, err := mgo.Dial("localhost")
	failOnError(err, "Connection MongoDB failed")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection test_data
	testCollection := session.DB("test").C("test_data")

	// Insert Datas

	//	sliceNumbers := make([]int64, howMany)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < howMany; i++ {
		var num int
		if rand.Float64() < 0.1 {
			num = 1
		} else {
			num = 0
		}
		doc := TestData{Value: num, DT: time.Now()}
		err = testCollection.Insert(doc)
		failOnError(err, "MongoDB insert failed")

	}

}
