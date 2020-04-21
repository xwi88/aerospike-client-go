package main

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"log"
	"math/rand"
	"time"
)

func main() {
	//10.14.41.51
	hostname := "10.14.41.51"
	port := 3000
	namespace := "bar"
	setName := "xwi88"

	client, err := as.NewClient(hostname, port)
	if err != nil {
		log.Fatalf("%v, NewClient %v", time.Now().String(), err)
	}
	defer client.Close()

	pkVal := int64(1)
	pkKey, err := as.NewKey(namespace, setName, pkVal)

	oldRecords, err := client.Get(nil, pkKey)
	if err != nil {
		log.Printf("%v, NewKey %v", time.Now().String(), err)
	}

	if oldRecords == nil {
		log.Printf("%v, no record exist with key: %v", time.Now().String(), pkVal)
		//key, err := as.NewKey(namespace, "set",
		//	"key value goes here and can be any supported primitive")
		pkKey, err = as.NewKey(namespace, setName, pkVal)
		if err != nil {
			log.Fatalf("%v, NewKey %v", time.Now().String(), err)
		}
	} else {
		oldKeyVal := oldRecords.Key.Value().GetObject().(int64)
		pkVal = oldKeyVal + 1
		pkKey, err = as.NewKey(namespace, setName, pkVal)
		if err != nil {
			log.Fatalf("%v, NewKey %v", time.Now().String(), err)
		}
	}

	binKey := as.NewBin("pk_int", pkVal)
	binInt := as.NewBin("int", rand.Int63())
	binStr := as.NewBin("str", "test_str")
	binBytes := as.NewBin("bytes", []byte("test_str"))
	binDouble := as.NewBin("double", rand.Float32())

	binListStr := as.NewBin("listStr", []string{"1", "2", "3"})
	binListInt := as.NewBin("listStr", []int{1, 2, 3})
	binMap := as.NewBin("map", map[string]interface{}{"a": 123, "b": "123"})
	binNestingList := as.NewBin("nestingList", [][]int{[]int{1, 2, 3}, []int{4, 5, 6}})
	binNestingMap := as.NewBin("nestingMap", map[string]interface{}{"a": 123,
		"b": map[string]interface{}{"bb": 123, "bc": "bc"},
		"c": []interface{}{1, "34"}})

	// Write a record
	err = client.PutBins(nil, pkKey, binKey, binInt, binStr, binBytes, binDouble, binListStr, binListInt, binMap,
		binNestingList, binNestingMap)
	if err != nil {
		log.Fatalf("%v, PutBins %v", time.Now().String(), err)
	}

	// Read a record
	record, err := client.Get(nil, pkKey)
	if err != nil {
		log.Fatalf("%v, Get %v", time.Now().String(), err)
	}

	fmt.Printf("%+v\n", record.String())

}
