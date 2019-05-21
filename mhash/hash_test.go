package mhash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T)  {
	fmt.Println(hash("11",0))
}

func TestHashTable(t *testing.T) {
	hashTable := InitHashTable(100)
	for i:= 0 ; i<100000;i ++ {
		hashTable.SetKey(strconv.Itoa(i),i)
	}
	hashTable.Show()
	fmt.Println(hashTable.GetKey("10"))
	fmt.Println(hashTable.GetKey("100"))
	fmt.Println(hashTable.GetKey("1000"))
	fmt.Println(hashTable.GetKey("10000"))

	hashTable.DelKey("10")
	hashTable.DelKey("100")
	hashTable.DelKey("100")
	hashTable.DelKey("1000")
	hashTable.DelKey("10000")

	fmt.Println(hashTable.GetKey("10"))
	fmt.Println(hashTable.GetKey("100"))
	fmt.Println(hashTable.GetKey("1000"))
	fmt.Println(hashTable.GetKey("10000"))

}