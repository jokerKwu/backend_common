package batch_common

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConf interface {
	GetSsmMongoInfo() ([]string, error)
	ConnectMongo(connInfos []string, isLocal bool) (*mongo.Client, error)
	InitCollection() error
	PingMongo(mongoClient *mongo.Client) error
}

func (a *EnvMongoDB) InitCollection() error {
	databaseName := a.Project + "_" + Env.Environment
	mongo := MongoClient.Database(databaseName)
	UserCollection = mongo.Collection("user")
	SubscriptionCollection = mongo.Collection("subscription")
	DeliveryCollection = mongo.Collection("delivery")
	PaymentCollection = mongo.Collection("payment")
	DeliveryHistoryCollection = mongo.Collection("deliveryHistory")

	return nil
}

func (a *EnvMongoDB) PingMongo(mongoClient *mongo.Client) error {
	err := mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("db 핑 통과")
	return nil
}

func (a *EnvMongoDB) GetSsmMongoInfo() ([]string, error) {
	var connInfos []string

	connInfos, err := AwsGetParams([]string{
		fmt.Sprintf("mongodb_%s_%s_id", a.Environment, a.Project),
		fmt.Sprintf("mongodb_%s_%s_pw", a.Environment, a.Project),
		fmt.Sprintf("mongodb_%s_%s_domain", a.Environment, a.Project),
	})
	if err != nil {
		return nil, err
	}

	return connInfos, nil
}

var (
	UserCollection             *mongo.Collection
	SubscriptionCollection     *mongo.Collection
	SubscriptionPlanCollection *mongo.Collection
	AddressBookCollection      *mongo.Collection
	DeliveryCollection         *mongo.Collection
	TermsCollection            *mongo.Collection
	ProductInfoCollection      *mongo.Collection
	QnACollection              *mongo.Collection
	PaymentCollection          *mongo.Collection
	EmailAuthCollection        *mongo.Collection
	DeliveryHistoryCollection  *mongo.Collection
	PaymentHistoryCollection   *mongo.Collection

	AppUserAuthCollection *mongo.Collection
	AppUserCollection     *mongo.Collection
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database
