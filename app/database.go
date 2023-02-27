package app

import (
	"errors"
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"path"
)

type (
	JsonDB struct {
		db *buntdb.DB
	}
)

// InitJsonDB Returns a new instance of json db
func (app *App) InitJsonDB() {
	// TODO: Rethink the name of the db
	fileName := "sycorax.db"

	log.Print("Initializing json db")

	log.Printf("Trying to open to open the db file, path: %s", app.Paths.Database)
	dbPath := path.Join(app.Paths.Database, fileName)
	db, err := buntdb.Open(dbPath)
	if err != nil {
		log.Fatalf("Can not open json db: %s", err)
	}

	// Creating indexes
	err = db.CreateIndex("queues", "queues:*", buntdb.IndexString)
	if err != nil {
		log.Fatalf("Error creating queues db index, err: %s", err)
		return
	}

	app.DB = &JsonDB{
		db,
	}
}

func (jdb *JsonDB) Close() {
	defer func(db *buntdb.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing db: ", err)
		}
	}(jdb.db)
}

func (jdb *JsonDB) Has(key string) bool {
	object := jdb.Get(key)
	exists := object != ""
	log.Printf("Object %s exists? %t", key, exists)
	return exists
}

func (jdb *JsonDB) Get(key string) string {
	var value = ""
	log.Printf("Getting %s on database", key)
	err := jdb.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		value = val
		return nil
	})
	if err != nil {
		log.Printf("Error getting %s on database, error: %s", key, err)
		return ""
	}
	return value
}

func (jdb *JsonDB) Set(key string, value string) error {
	log.Printf("Setting %s on database", key)
	err := jdb.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		if err != nil {
			return err
		}
		log.Printf("Object %s successfully setted on database", key)
		return nil
	})
	if err != nil {
		log.Printf("Error setting %s on database, error: %s", key, err)
		return err
	}
	return nil
}

func (jdb *JsonDB) Delete(key string) error {
	log.Printf("Deletting %s on database", key)
	err := jdb.db.Update(func(tx *buntdb.Tx) error {
		val, err := tx.Delete(key)
		if err != nil {
			return err
		}
		if val != "" {
			return errors.New(fmt.Sprintf("Key %s and Value %s was not deletted from database", key, val))
		}
		log.Printf("Object %s successfully deleted", key)
		return nil
	})
	if err != nil {
		log.Printf("Error removing key  %s on database, error: %s", key, err)
		return err
	}
	return nil
}

func (jdb *JsonDB) FindManyByIndex(index string) ([]string, error) {
	var list []string
	err := jdb.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(index, func(key, value string) bool {
			list = append(list, value)
			return true
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("Error finding index  %s on db, err: %s", index, err)
		return nil, err
	}
	return list, nil
}

func (jdb *JsonDB) MakeQueueKey(id string) string {
	return fmt.Sprintf("queues:%s", id)
}
