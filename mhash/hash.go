package mhash

import (
	"fmt"
	"math/rand"
	"time"
)

// 非常简单 随机 hash 函数
func hash(key string , rand int) int   {
	bytes := []byte(key)
	result := 0
	for _,v := range bytes {
		result = (int(v)+result) * 7
	}
	return  result + rand
}

type HashTable struct {
	len int
	rand int
	data []*link
}

func InitHashTable(length ... int)  HashTable {
	l := 17
	if len(length) >0    {
		l = length[0]
	}
	rand.Seed(time.Now().Unix())
	return HashTable{
		len:l,
		rand:rand.Int(),
		data:make([]*link,l),
	}
}

// s设置key
func (table *HashTable) SetKey (key string,values int )  {
	if table == nil {
		panic("table not exits")
		return
	}
	hs := hash(key,table.rand)
	index := hs % table.len
	if table.data[index] == nil {
		table.data[index] = new(link)
	}
	table.data[index].set(key,values)
	return
}

// 获取key
func (table *HashTable) GetKey (key string)  (int ,bool){
	if table == nil {
		return 0,false
	}
	hs := hash(key,table.rand)
	index := hs % table.len
	return  table.data[index].get(key)
}

// 删除key
func (table *HashTable) DelKey (key string){
	if table == nil {
		return
	}
	hs := hash(key,table.rand)
	index := hs % table.len
	table.data[index].del(key)
	return
}

func (table *HashTable)Show()  {
	total := 0
	for i := range table.data {
		l := table.data[i].len()
		fmt.Println("长度:",i,":",l)
		total +=l
	}
	fmt.Println("总长度:",total)
}