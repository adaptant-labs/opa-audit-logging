package main

import (
	"context"
	opalogs "github.com/open-policy-agent/opa/plugins/logs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
)

var (
	dbName = "audit-logs"
)

func NewMongoClient() *mongo.Client {
	var opts *options.ClientOptions

	// If a scheme has already been defined, use it outright
	if strings.Contains(mongoServer, "://") {
		opts = options.Client().ApplyURI(mongoServer)
	} else {
		opts = options.Client().ApplyURI("mongodb://" + mongoServer)
	}
	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Couldn't connect to MongoDB", err)
	} else {
		log.Info("Connected to MongoDB instance")
	}

	return client
}

type MongoDataService struct {
	mgo			*mongo.Client
	decisions	*mongo.Collection
}

type MongoDecisionModel struct {
	ID        string          `json:"_id,omitempty" bson:"_id"`
	Partition string          `json:"partition,omitempty" bson:"partition,omitempty"`
	Decision  opalogs.EventV1 `json:"decision" bson:"decision"`
}

func (mds MongoDataService) AddDecisionToPartition(partition string, decision opalogs.EventV1) error {
	var mdm MongoDecisionModel

	// Use the Decision ID for the Mongo Object ID
	mdm.ID = decision.DecisionID
	mdm.Decision = decision

	if len(partition) > 0 {
		mdm.Partition = partition
	}

	_, err := mds.decisions.InsertOne(context.Background(), mdm)
	if err != nil {
		return err
	}

	return nil
}

func (mds MongoDataService) AddDecision(decision opalogs.EventV1) error {
	return mds.AddDecisionToPartition("", decision)
}

func NewMongoDataService() AuditLoggingDataService {
	mds := &MongoDataService{mgo: NewMongoClient()}
	mds.decisions = mds.mgo.Database(dbName).Collection("decisions")
	return mds
}