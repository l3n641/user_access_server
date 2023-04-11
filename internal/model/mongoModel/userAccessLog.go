package mongoModel

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"user_accerss_server/internal/database/mongoDb"
)

type UserAccessLog struct {
	Id              primitive.ObjectID `bson:"_id"`               //记录的日期
	Date            time.Time          `bson:"date"`              //记录的日期
	FirstAccessTime time.Time          `bson:"first_access_time"` //第一次访问的时间
	LastAccessTime  time.Time          `bson:"last_access_time"`  //最后一次访问的时间
	Domain          string             `bson:"domain"`            //域名
	SessionID       string             `bson:"session_id"`        // 会话id
	ClientIP        string             `bson:"client_ip"`         // 客户ip
	ClientCountry   string             `bson:"client_country"`    // 客户所属的国家
	UserAgent       string             `bson:"user_agent"`        //客户浏览器ua
	Referer         string             `bson:"referer"`           //客户referer
	PageViews       uint               `bson:"page_views"`        //pv 次数
}

func (p *UserAccessLog) GetCollection() *mongo.Collection {
	db := mongoDb.GetMongoDb()
	collection := db.Collection("user_access_log")
	return collection
}

func (p *UserAccessLog) Upsert(isNew bool) {

	collection := p.GetCollection()
	options.Update().SetUpsert(true)
	filter := bson.M{"date": p.Date, "domain": p.Domain, "session_id": p.SessionID}
	opts := options.Update().SetUpsert(true)

	update := bson.D{
		{
			"$inc", bson.D{{ // $inc 代表增加或减少
				"page_views", 1, // 在原值的基础上 +1
			}},
		},
		{"$set", bson.D{{"last_access_time", p.LastAccessTime}}},
		{"$set", bson.D{{"client_ip", p.ClientIP}}},
		{"$set", bson.D{{"client_country", p.ClientCountry}}},
		{"$set", bson.D{{"user_agent", p.UserAgent}}},
		{"$set", bson.D{{"referer", p.Referer}}},
	}
	if isNew {
		update = append(update, bson.E{"$set", bson.D{{"first_access_time", p.FirstAccessTime}}})
	}
	// 执行 upsert 操作
	collection.UpdateOne(context.TODO(), filter, update, opts)

}

func (p *UserAccessLog) Add() (primitive.ObjectID, error) {

	collection := p.GetCollection()
	id := primitive.NewObjectID()
	p.Id = id
	iResult, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		return id, err
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (p *UserAccessLog) Update(accessTime time.Time) (primitive.ObjectID, error) {
	var prototype UserAccessLog

	filter := bson.M{"_id": p.Id} //查询id
	after := options.After        // 设置返回的结果为修改后的

	// 如果要查找修改前的值用options.Before

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{
		{
			"$inc", bson.D{{ // $inc 代表增加或减少
				"page_views", 1, // 在原值的基础上 +1
			}},
		},
		{"$set", bson.D{{"last_access_time", accessTime}}},
	}

	// collection 为 某一个具体的Collection.
	collection := p.GetCollection()
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&prototype)
	return p.Id, err

}
