package batch_common

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type EnvMongoDB struct {
	Project     string
	Environment string
	RsName      string
	LocalPort   int64
}

//mongodb+srv://ryan:tngkr485@cluster0.ke0pv.mongodb.net/?retryWrites=true&w=majority

/*
ProjectEnv

   PROJECT: 'medical_web'
   ENV: 'dev'
   REGION: 'ap-northeast-2'
*/
var MongoDBEnv EnvMongoDB

func InitMongoDB() error {
	var err error
	connUriDB := fmt.Sprintf("%s_%s", Env.Project, Env.Environment)
	//TODO MongoDB 초기화 함수
	//TODO 1. 몽고디비Env 초기화
	InitMongoDBEnv()
	//TODO 2. connURI 가져오기
	connUri := MongoDBEnv.MakeConnURI(false)
	//TODO 3. 몽고디비 연결
	MongoClient, err = MongoDBEnv.ConnectMongo(connUri)
	if err != nil {
		return err
	}
	//TODO 4. 핑 테스트
	if err = MongoDBEnv.PingMongo(MongoClient); err != nil {
		return err
	}
	//TODO 5. 컬렉션 초기화
	MongoDB = MongoClient.Database(connUriDB)
	//MongoDBEnv.InitCollection()
	return nil
}

func InitMongoDBEnv() {
	MongoDBEnv.Environment = Env.Environment
	MongoDBEnv.Project = Env.Project
	fmt.Println(MongoDBEnv)
}

func (a *EnvMongoDB) MakeConnURI(isLocal bool) string {
	var connUri string

	if Env.Environment == "prd" {
		//TODO PRD connection URI
	} else {
		fmt.Println("여기 들어가야되는데")
		/*
			회사에서 사용하는 디비 정보
			connInfos, _ := a.GetSsmMongoInfo()
			connUriID := connInfos[0]
			connUriPW := connInfos[1]
			additionalOpt := ""
			connUriDomain := connInfos[2]
			connUri = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&replicaSet=%s&w=majority&readPreference=primary&retryWrites=true&ssl=false%s", connUriID, connUriPW, connUriDomain, a.RsName, additionalOpt)
			if isLocal == true {
				connUriDomain = fmt.Sprintf("localhost:%d", a.LocalPort)
				additionalOpt = "&directConnection=true"
				connUri = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&replicaSet=%s&w=majority&readPreference=primary&retryWrites=true&ssl=false%s", connUriID, connUriPW, connUriDomain, a.RsName, additionalOpt)
			}
		*/
		// 개인 계정으로 임시로 테스트 진행
		connInfos, err := AwsGetParam("mongodb_medical_web_dev_tmp")
		fmt.Println(err)
		connUri = connInfos
	}
	fmt.Println(connUri)
	return connUri
}

func (a *EnvMongoDB) ConnectMongo(connUri string) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI(connUri)
	clientOptions.SetMaxPoolSize(1)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetMaxConnIdleTime(20 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return &mongo.Client{}, err
	}
	return mongoClient, nil
}

func CloseMongo() {
	_ = MongoClient.Disconnect(context.TODO())
}
