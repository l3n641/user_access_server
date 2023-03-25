package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user_accerss_server/api/params"
	"user_accerss_server/internal/database/mongoDb"
	"user_accerss_server/internal/model/mongoModel"
)

type UserAccessDomainLogService struct {
}

func (u UserAccessDomainLogService) GetList(param params.AccessDomainLogGetParam) (*[]mongoModel.UserAccessDomainLog, int64) {
	var cursor *mongo.Cursor
	var err error
	var data []mongoModel.UserAccessDomainLog
	var accessLog mongoModel.UserAccessDomainLog

	model := mongoModel.UserAccessDomainLog{}
	collection := model.GetCollection()
	dateStart := param.GetStartLocalTime()
	dateEnd := param.GetEndLocalTime()
	filter := map[string]interface{}{
		"date": bson.M{"$gte": dateStart, "$lte": dateEnd},
	}
	if param.Domain != "" {
		filter["domain"] = param.Domain
	}

	pageOption := mongoDb.GetPageSizeOption(param.Page, param.PageSize)

	if cursor, err = collection.Find(context.TODO(), filter, pageOption); err != nil {
		log.Fatal(err)
		log.Fatal(filter)
	}
	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&accessLog); err != nil {
			log.Fatal(err)
		} else {
			data = append(data, accessLog)
		}
	}
	itemCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return &data, itemCount
}