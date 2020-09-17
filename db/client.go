package db

import (
	"context"
	"fmt"
	"github.com/bpsaunders/user-api/config"
	"github.com/bpsaunders/user-api/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

// Client provides an interface by which to interact with a database
type Client interface {
	CreateUser(entity *models.UserDao) error
	GetUser(id string) (*models.UserDao, error)
	GetAllUsers() (*[]*models.UserDao, error)
	UserExistsWithEmail(email string) (bool, error)
	Shutdown()
}

// DatabaseClient is a concrete implementation of the Client interface
type DatabaseClient struct {
	db MongoDatabaseInterface
}

// NewDatabaseClient returns a new implementation of the Client interface
func NewDatabaseClient(cfg *config.Config) Client {
	return &DatabaseClient{
		db: getMongoDatabase(cfg.MongoDBURL, cfg.MongoDBDatabase),
	}
}

var mgoClient *mongo.Client

func getMongoClient(mongoDBURL string) *mongo.Client {

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(ctx, clientOptions)

	// the program must bail out here if failing to establish a connection to the db, as this will run on application start-up
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// cache mongo client here, in preparation for disconnect on application shutdown
	mgoClient = client

	// check we can connect to the mongodb instance - again, bail out on failure
	pingContext, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	err = client.Ping(pingContext, nil)
	if err != nil {
		log.Error("ping to mongodb timed out. please check the connection to mongodb and that it is running")
		os.Exit(1)
	}

	log.Info("connected to mongodb successfully")

	return client
}

func getMongoDatabase(mongoDBURL, databaseName string) MongoDatabaseInterface {
	return getMongoClient(mongoDBURL).Database(databaseName)
}

// MongoDatabaseInterface is an interface that describes the mongodb driver
type MongoDatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

// CreateUser creates a user entity in the database
func (c *DatabaseClient) CreateUser(entity *models.UserDao) error {

	collection := c.db.Collection("users")
	_, err := collection.InsertOne(context.Background(), entity)

	return err
}

// GetUser fetches a user from the db according to an id
func (c *DatabaseClient) GetUser(id string) (*models.UserDao, error) {

	var entity models.UserDao

	collection := c.db.Collection("users")
	dbResource := collection.FindOne(context.Background(), bson.M{"_id": id})

	err := dbResource.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	err = dbResource.Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// GetAllUsers returns an array of all users in the database
func (c *DatabaseClient) GetAllUsers() (*[]*models.UserDao, error) {

	entities := make([]*models.UserDao, 0)

	collection := c.db.Collection("users")
	cur, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {

		var entity models.UserDao
		err = cur.Decode(&entity)

		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return &entities, nil
}

// UserExistsWithEmail determines whether a user already exists in the database according to an email
func (c *DatabaseClient) UserExistsWithEmail(email string) (bool, error) {

	collection := c.db.Collection("users")
	dbResource := collection.FindOne(context.Background(), bson.M{"email": email})

	err := dbResource.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Shutdown is a hook that can be used to clean up db resources
func (c *DatabaseClient) Shutdown() {
	log.Info("Attempting to close the db connection thread pool")
	if mgoClient != nil {
		err := mgoClient.Disconnect(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("Failed to disconnect from the mongodb: %s", err))
			return
		}
		log.Info("disconnected from mongodb successfully")
	}
}
