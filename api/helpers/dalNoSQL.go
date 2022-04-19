package helpers

import (
	"context"
	"sync"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/spf13/viper"
)

// MongoHost -MongoHost
type MongoHost struct {
	HostName        string        `json:"hostName"`
	Server          string        `json:"server"`
	Port            int           `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Database        string        `json:"database"`
	IsDefault       bool          `json:"isDefault"`
	MaxIdleConns    int           `json:"maxIdleConns" `
	MaxOpenConns    int           `json:"maxOpenConns"`
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" `
	IsDisabled      bool          `json:"isDisabled" `
}

var instance *mongo.Client
var mutex sync.Mutex
var once sync.Once

// func init() {
// 	config = make(map[string]MongoHost)
// }

var (
	Username  = confighelper.GetConfig("MONGO_USERNAME")
	Password  = confighelper.GetConfig("MONGO_PASSWORD")
	Database  = confighelper.GetConfig("DBNAME")
	PoolLimit = viper.GetInt("PoolLimit")
	Source    = confighelper.GetConfig("MONGO_AUTHENTICATION_DB")
)

// GetMongoConnection - GetMongoConnection
func GetMongoConnection() (*mongo.Client, error) {
	once.Do(func() {
		defer mutex.Unlock()
		clientOption := options.Client()

		clientOption.SetHosts([]string{confighelper.GetConfig("MONGODSN")}).
			SetMaxConnIdleTime(time.Second * time.Duration(viper.GetInt("MONGO_MaxConnIdleTime"))).
			SetConnectTimeout(time.Second * 3).
			SetMaxPoolSize(uint64(PoolLimit)).
			SetReadPreference(readpref.Primary())
			// SetSocketTimeout(1 * time.Minute).
			// SetDirect(true) // important if in cluster, connect to primary only.

		if confighelper.GetConfig("MONGO_USERNAME") != "" {
			cred := options.Credential{}
			cred.Username = confighelper.GetConfig("MONGO_USERNAME")
			cred.Password = confighelper.GetConfig("MONGO_PASSWORD")
			cred.AuthSource = confighelper.GetConfig("DBNAME")
			clientOption.SetAuth(cred)
		}
		client, err := mongo.NewClient(clientOption)
		if err != nil {
			logginghelper.LogError(err)
			return
		}
		err = client.Connect(context.Background())
		if err != nil {
			logginghelper.LogError(err)
			return
		}
		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			logginghelper.LogError("failed to connect to primary - ", err)
			return
		}

		instance = client
	})
	return instance, nil
}

// EnsureIndex added by Mayur Wadekar
// EnsureIndex will create index on collection provided
func EnsureIndex(cd *mongo.Collection, indexQuery []string) error {

	opts := options.CreateIndexes().SetMaxTime(5 * time.Second)

	index := []mongo.IndexModel{}

	for _, val := range indexQuery {
		temp := mongo.IndexModel{}
		temp.Keys = bsonx.Doc{{Key: val, Value: bsonx.Int32(1)}}
		index = append(index, temp)
	}

	_, err := cd.Indexes().CreateMany(context.Background(), index, opts)
	if err != nil {
		logginghelper.LogError("Error while executing index Query", err)
		return err
	}
	return nil
}

// CheckCollectionExists added by Mayur Wadekar
// CheckCollectionExists will check provided collection present or not in db
func CheckCollectionExists(coll string) (bool, error) {
	session, dberr := GetMongoConnection()
	if dberr != nil {
		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
		return false, dberr
	}

	names, err := session.Database(confighelper.GetConfig("DBNAME")).ListCollectionNames(context.Background(), bson.M{"listCollections": 1, "filter": bson.M{"name": coll}})
	if err != nil {
		logginghelper.LogError("ERROR in getting collection names", err)
		return false, err
	}

	for _, name := range names {
		if name == coll {
			return true, nil
		}
	}
	return false, nil
}
