package tdlib

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type MyDB struct {
	db   *leveldb.DB
	lock sync.RWMutex
}

var (
	myDB *MyDB
)

func InitDB() (*MyDB, error) {
	if myDB == nil {
		db, err := leveldb.OpenFile("./db", nil)
		if err != nil {
			return nil, err
		}
		myDB = &MyDB{db: db}
		return myDB, err
	} else {
		return myDB, nil
	}

}

func (m *MyDB) Range(r func(key, value string)) {
	iterator := m.db.NewIterator(&util.Range{}, &opt.ReadOptions{
		DontFillCache: true,
		Strict:        0,
	})
	for iterator.Next() {
		r(string(iterator.Key()), string(iterator.Value()))
	}
}

func (m *MyDB) Delete(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	err := m.db.Delete([]byte(key), &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         true,
	})
	if err != nil {
		return err
	}
	return err
}

func (m *MyDB) Load(key string) (string, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	data, err := m.db.Get([]byte(key), &opt.ReadOptions{
		DontFillCache: true,
		Strict:        0,
	})
	if err != nil {
		return "", err
	}
	return string(data), err
}

func (m *MyDB) Store(key, value string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	err := m.db.Put([]byte(key), []byte(value), &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         true,
	})
	if err != nil {
		return err
	}
	return err
}

func (m *MyDB) Close() {
	err := m.db.Close()
	if err != nil {
		log.Errorln("关闭db出现错误" + err.Error())
		return
	}
}
