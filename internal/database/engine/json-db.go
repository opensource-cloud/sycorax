package jsondb

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
	"path"
)

type (
	JsonDB struct {
		db *buntdb.DB
	}
	JsonDBConfig struct {
		Path   string
		DBName string
	}
)

// NewJsonDB Returns a new instance of json db
func NewJsonDB(config *JsonDBConfig) (*JsonDB, error) {
	log.Println("--------------- [DB] ---------------")
	log.Print("Initializing json db")

	log.Printf("Trying to open to open the db file, path: %s", config.Path)
	dbPath := path.Join(config.Path, config.DBName)
	db, err := buntdb.Open(dbPath)
	if err != nil {
		log.Fatalf("Can not open json db: %s", err)
	}

	log.Print("Setting db config")
	// Config.SyncPolicy = Always - fsync after every write, very durable, slower
	err = db.SetConfig(buntdb.Config{
		SyncPolicy: 2,
	})
	if err != nil {
		log.Fatalf("Error setting db config: %v", err)
		return nil, err
	}

	// Creating indexes
	log.Println("Creating queues db index")
	err = db.CreateIndex("queues", "queues:*", buntdb.IndexJSON("refId"))
	if err != nil {
		log.Fatalf("Error creating queues db index, err: %s", err)
		return nil, err
	}
	log.Println("Queues index db created")

	log.Println("DB Initialized")
	log.Println("--------------- [DB] ---------------")

	return &JsonDB{
		db: db,
	}, nil
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
		_, err := tx.Delete(key)
		return err
	})
	if err != nil {
		log.Printf("Error deleting key %s, detail: %v", key, err)
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
