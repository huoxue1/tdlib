package impl

import (
	"fmt"
	"github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/utils/help"
	"github.com/nutsdb/nutsdb"
	"time"
)

type NustdbClient struct {
	db *nutsdb.DB
}

func (n *NustdbClient) SetAdd(key, value string) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.SAdd(BucketSet, help.StringToByteSlice(key), help.StringToByteSlice(value))
	})
}

func (n *NustdbClient) SetIsMem(key string, member string) bool {
	tx, err := n.db.Begin(false)
	if err != nil {
		return false
	}
	defer tx.Commit()
	isMember, err := tx.SIsMember(BucketSet, help.StringToByteSlice(key), help.StringToByteSlice(member))
	if err != nil {
		return false
	}
	return isMember
}

func (n *NustdbClient) SetDel(key string, member string) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.SRem(BucketSet, help.StringToByteSlice(key), help.StringToByteSlice(member))
	})
}

const (
	Bucket    = "tdlib"
	BucketSet = "tdlib-set"
)

func InitNustdb(config *conf.Config) (*NustdbClient, error) {
	db, err := nutsdb.Open(nutsdb.DefaultOptions, nutsdb.WithDir(config.Cache.Nustdb.Dir))
	if err != nil {
		return nil, err
	}
	return &NustdbClient{
		db: db,
	}, nil
}

func (n *NustdbClient) Set(key, value string) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(Bucket, help.StringToByteSlice(key), help.StringToByteSlice(value), 0)
	})
}

func (n *NustdbClient) SetTtl(key, value string, ttl time.Duration) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(Bucket, help.StringToByteSlice(key), help.StringToByteSlice(value), uint32(ttl))
	})
}

func (n *NustdbClient) Get(key string) string {
	tx, err := n.db.Begin(false)
	if err != nil {
		return ""
	}
	defer tx.Commit()
	entry, err := tx.Get(Bucket, help.StringToByteSlice(key))
	if err != nil {
		return ""
	}
	return help.ByteSliceToString(entry.Value)
}

func (n *NustdbClient) GetDefault(key string, defaultValue string) string {
	tx, err := n.db.Begin(false)
	if err != nil {
		return defaultValue
	}
	defer tx.Commit()
	entry, err := tx.Get(Bucket, help.StringToByteSlice(key))
	if err != nil {
		return defaultValue
	}
	return help.ByteSliceToString(entry.Value)
}

func (n *NustdbClient) Delete(key string) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Delete(Bucket, help.StringToByteSlice(key))
	})
}

func (n *NustdbClient) ForEach(f func(key string, value string) bool) {
	n.Set("test", "test")
	get := n.Get("test")
	fmt.Println(get)
	_ = n.db.View(func(tx *nutsdb.Tx) error {
		entries, err := tx.GetAll(Bucket)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if !f(help.ByteSliceToString(entry.Key), help.ByteSliceToString(entry.Value)) {
				break
			}
		}
		return nil
	})
}
