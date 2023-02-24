package database

import (
	"github.com/opensource-cloud/sycorax/infrastructure/config"
	"github.com/tidwall/buntdb"
	"log"
	"path"
)

type (
	JsonDB struct {
		db *buntdb.DB
	}
)

// NewJsonDB Returns a new instance of json db
func NewJsonDB(app *config.App, fileName string) *JsonDB {
	log.Print("Initializing json db")

	log.Printf("Trying to open to open the db file, path: %s", app.Paths.Database)
	dbPath := path.Join(app.Paths.Database, fileName)
	db, err := buntdb.Open(dbPath)
	if err != nil {
		log.Fatal("can not open json db: ", err)
	}

	log.Print("Queues DB loaded, creating JsonDB struct")

	return &JsonDB{
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

func (jdb *JsonDB) Get(key string) string {
	var value = ""
	log.Printf("Getting %s on database", key)
	err := jdb.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			log.Printf("Error getting %s on database, error: %s", key, err)
			return err
		}
		value = val
		return nil
	})
	if err != nil {
		log.Print(err)
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
