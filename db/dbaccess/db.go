package dbaccess

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/KaoNinjaratzu/outrun/config"
)

var db *sql.DB
var DatabaseIsBusy = false

// TODO: Reimplement all the database access functions here for MySQL.

func Set(bucket, key string, value []byte) error {
	CheckIfDBSet()
	/*value = Compress(value) // compress the input first
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
	return err*/
	return nil
}

func Get(bucket, key string) ([]byte, error) {
	CheckIfDBSet()
	/*var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = b.Get([]byte(key))
		if value == nil {
			return errors.New("no value named '" + key + "' in bucket '" + bucket + "'")
		}
		return nil
	})
	result, derr := Decompress(value) // decompress the result
	if derr != nil {
		return result, derr
	}
	return result, err*/
	return []byte{}, nil
}

func Delete(bucket, key string) error {
	CheckIfDBSet()
	/*return db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})*/
	return nil
}

func ForEachKey(bucket string, each func(k, v []byte) error) error {
	CheckIfDBSet()
	/*err := db.View(func(tx *sql.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err2 := b.ForEach(each)
		return err2
	})
	return err*/
	return nil
}

func ForEachLogic(each func(tx *sql.Tx) error) error {
	CheckIfDBSet()
	/*err := db.View(each)
	return err*/
	return nil
}

func CheckIfDBSet() {
	if db == nil {
		log.Println("Connecting to MySQL database...")
		sqldb, err := sql.Open("mysql", config.CFile.MySQLUsername+":"+config.CFile.MySQLPassword+"@"+config.CFile.MySQLServerAddress+"/"+config.CFile.MySQLDatabaseName)
		if err != nil {
			panic(err)
		}
		err = sqldb.Ping()
		if err != nil {
			panic(err)
		}
		db = sqldb
		log.Println("Successfully connected to database!")
	}
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return errors.New("cannot close database if it's not set!")
}
