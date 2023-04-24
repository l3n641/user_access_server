package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user_accerss_server/api/userAccess/params"
	"user_accerss_server/internal/database/mongoDb"
	"user_accerss_server/internal/model/mongoModel"
)

type UserAccessDetailService struct {
}

func (u UserAccessDetailService) GetList(param params.AccessUserLogGetParam) (*[]mongoModel.UserAccessDetail, int64) {
	var cursor *mongo.Cursor
	var err error
	var accessLogList []mongoModel.UserAccessDetail
	var accessLog mongoModel.UserAccessDetail

	model := mongoModel.UserAccessDetail{}
	collection := model.GetCollection()
	dateStart := param.GetStartLocalTime()
	dateEnd := param.GetEndLocalTime()
	filter := map[string]interface{}{
		"create_time": bson.M{"$gte": dateStart, "$lte": dateEnd},
	}
	if param.Domain != "" {
		filter["domain"] = param.Domain
	}

	if param.SessionID != "" {
		filter["session_id"] = param.SessionID
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
			accessLogList = append(accessLogList, accessLog)
		}
	}
	itemCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return &accessLogList, itemCount
}
