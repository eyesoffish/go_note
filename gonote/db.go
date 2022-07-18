package note

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/syndtr/goleveldb/leveldb"
	leveldbUtil "github.com/syndtr/goleveldb/leveldb/util"
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

func LevelDBIterate() {
	db, err := leveldb.OpenFile("leveldb.db", nil)
	if err != nil {
		fmt.Println("数据库打开失败")
		panic(err)
	} else {
		fmt.Println("数据库打开成功")
	}
	defer db.Close()
	batch := new(leveldb.Batch)
	for i := 1; i < 11; i++ {
		batch.Put([]byte(fmt.Sprintf("user-%v", i)), []byte("{\"username\":\"三角形\",\"age\":13}"))
	}
	db.Write(batch, nil)
	iter := db.NewIterator(&leveldbUtil.Range{Start: []byte("user-3"), Limit: []byte("user-9")}, nil)
	for iter.Next() {
		fmt.Printf("%v = %v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	fmt.Println("iter err = ", err)
	iter = db.NewIterator(leveldbUtil.BytesPrefix([]byte("user-")), nil)
	for iter.Next() {
		fmt.Printf("%v = %v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
}

// leveldb快照
func LeveldbTransactionAndSnapshot() {
	db, err := leveldb.OpenFile("leveldb.db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ss, err := db.GetSnapshot()
	if err != nil {
		panic(err)
	}
	defer ss.Release()
	t, err := db.OpenTransaction()
	if err != nil {
		panic(err)
	}
	batch := new(leveldb.Batch)
	for i := 1; i < 11; i++ {
		batch.Put([]byte(fmt.Sprintf("cat-%v", i)), []byte("{\"username\":\"猫\",\"age\":1}"))
	}
	t.Write(batch, nil)
	// t.Discard()
	t.Commit()
	ok, _ :=db.Has([]byte("cat-1"), nil)
	fmt.Println("db hat cat - 1 ? ", ok)
	ok, _ =ss.Has([]byte("cat-1"), nil)
	fmt.Println("ss hat cat - 1 ? ", ok)
}

// go redis
func RedisBasic() {
	opt := redis.Options{
		Addr: "localhost:6379",
	}
	db := redis.NewClient(&opt)
	ctc:=context.Background()
	db.Do(ctc, "set", "k1", "v1")
	res, err:=db.Do(ctc, "get","k1").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("该Key不存在")
		}
		fmt.Println("err = ", err)
	} else {
		fmt.Println("get redis res = ", res.(string))
	}

	// 
	db.Do(ctc, "set", "b1", true)
	db.Do(ctc, "set", "b2", 0)
	b, err := db.Do(ctc, "get", "b1").Bool()
	if err == nil {
		fmt.Println("b = ", b)
	}
	b2, err1 := db.Do(ctc, "mget","b1", "b2").BoolSlice()
	if err1 == nil {
		fmt.Println("b=", b2)
	}

	// time = 0 表示永不过期
	db.Set(ctc, "t1", time.Now(), 10 * time.Second)
	t1 := db.Get(ctc, "t1").Val()
	fmt.Printf("t1 = %v", t1)
}