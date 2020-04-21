package main

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"log"
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
		log.Printf("%v, Get %v", time.Now().String(), err)
	}

	ops := []*as.Operation{
		//as.AddOp(as.NewBin("visit_times", 0)), // add the value of the bin to the existing value
		//as.ListIncrementOp("visit_times", 0, 1),
		//as.ListAppendOp("list_double",  1.02),
		as.ListIncrementOp("list_double",  -1,1.15),
		//as.MapIncrementOp(as.DefaultMapPolicy(), "map", "a",2),
		//as.MapIncrementOp(as.DefaultMapPolicy(), "nestingMap", "b.bb",3), // fails
		//as.MapIncrementOp(as.DefaultMapPolicy(), "nestingMap", "bb",3), // fails, 不存在则新加了一个列，嵌套类型不能操作!
		//as.MapIncrementOp(as.DefaultMapPolicy(), "nestingMap", "b",3), // fails, 不存在则新加了一个列，嵌套类型不能操作!
		//as.ListIncrementOp("list_int", 1, 2),
		//as.ListIncrementOp("list_int", 2, 3),
		//as.AddOp(as.NewBin("test_str", "1")), // 存在则停止下面执行
		//as.AppendOp(as.NewBin("test_str", "1")), // 必须存在
		//as.AddOp(as.NewBin("list_int", []int{1,2,3})),
		//as.ListInsertOp("list_int", -1, 4), // -1 最后一个
		//as.ListAppendOp("listStr",  "123"),
		//as.ListPopOp("list_int", -1),
		//as.GetOp(), // get the value of the record after all operations are executed, 部分操作不能跟这个，不然不执行!
	}

	rec, err := client.Operate(nil, pkKey, ops...)
	if err != nil {
		log.Printf("%v, Operate %v", time.Now().String(), err)
	}
	fmt.Printf("OldRecords: %#v\n", oldRecords)
	fmt.Printf("Records now: %#v\n", rec)

}
