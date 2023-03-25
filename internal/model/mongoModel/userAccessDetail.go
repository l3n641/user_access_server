package mongoModel

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user_accerss_server/internal/database/mongoDb"
)

type UserAccessDetail struct {
	CreateTime    time.Time `bson:"create_time"`    //创建记录的时间
	Domain        string    `bson:"domain"`         //域名
	Uri           string    `bson:"uri"`            //网页链接
	PageType      string    `bson:"page_type"`      //链接类型
	SessionID     string    `bson:"session_id"`     // 会话id
	ClientIP      string    `bson:"client_ip"`      // 客户ip
	ClientCountry string    `bson:"client_country"` // 客户所属的国家
	UserAgent     string    `bson:"user_agent"`     //客户浏览器ua
	Referer       string    `bson:"referer"`        //客户referer
}

func (p *UserAccessDetail) GetCollection() *mongo.Collection {
	db := mongoDb.GetMongoDb()
	collection := db.Collection("user_access_detail")
	return collection
}

func (p *UserAccessDetail) Add() (interface{}, error) {

	collection := p.GetCollection()
	iResult, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	//_id:默认生成一个全局唯一ID
	id := iResult.InsertedID.(primitive.ObjectID)

	return id, nil
}
