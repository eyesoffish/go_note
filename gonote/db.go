package note

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// levelDB 基本使用
func LevelDBBasic() {
	db,err := leveldb.OpenFile("leveldb.db", nil)
	if err != nil {
		fmt.Println("数据打开失败", err)
		return
	}
	defer db.Close()

	db.Put([]byte("user-1"), []byte("{\"username\":\"三角形\",\"age\":10}"), nil)
	// db.Put([]byte("user-2"), []byte("{\"username\":\"方块\",\"age\":11}"), nil)
	db.Delete([]byte("user-3"), nil)
	db.Has([]byte("user-3"), nil)
	data, _ := db.Get([]byte("user-1"), nil)
	fmt.Println("data = ", string(data))
	batch := new(leveldb.Batch)
	batch.Delete([]byte("user-1"))
	// batch.Put([]byte("user-1"), []byte("{\"username\":\"正方形\",\"age\":10}"))
	batch.Put([]byte("user-2"), []byte("{\"username\":\"三角形\",\"age\":11}"))
	batch.Put([]byte("user-3"), []byte("{\"username\":\"三角形\",\"age\":12}"))
	batch.Put([]byte("user-4"), []byte("{\"username\":\"三角形\",\"age\":13}"))
	
	db.Write(batch, nil)

	data1, _ := db.Get([]byte("user-3"), nil)
	fmt.Println("data = ", string(data1))
}