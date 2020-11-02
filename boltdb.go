package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func errLog(msg string) {
	db, err := bolt.Open("errorMsg.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ErrMsgLog"))
		if b == nil {
			new, err := tx.CreateBucket([]byte("ErrMsgLog"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			if err := new.Put([]byte("errMsg"), []byte(msg)); err != nil {
				return err
			}
			if err := new.Put([]byte("date"), []byte(time.Now().Format("2006-01-02 15:04:05"))); err != nil {
				return err
			}
		} else {
			if err := b.Put([]byte("errMsg"), []byte(msg)); err != nil {
				return err
			}
			if err := b.Put([]byte("date"), []byte(time.Now().Format("2006-01-02 15:04:05"))); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
