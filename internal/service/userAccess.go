package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"user_accerss_server/api/userAccess/params"
	"user_accerss_server/internal/database/mongoDb"
	"user_accerss_server/internal/model/mongoModel"
	"user_accerss_server/internal/tools"
)

type UserAccessService struct {
}

func (u UserAccessService) GetRecordBySessionId(id, domain string, date time.Time) (*mongoModel.UserAccessLog, error) {
	var result mongoModel.UserAccessLog

	filter := bson.M{"date": date, "domain": domain, "session_id": id}
	m := mongoModel.UserAccessLog{}
	collection := m.GetCollection()
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}

func (u UserAccessService) AddRecord(param params.AccessLogPostParam) {
	var err error
	currentTime := time.Now()
	city := tools.ParseIp(param.ClientIP)
	today := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	_, err = u.GetRecordBySessionId(param.SessionID, param.Domain, today)
	domainLog := mongoModel.UserAccessDomainLog{
		Date:   today,
		Domain: param.Domain,
	}

	userLog := mongoModel.UserAccessLog{
		Date:            today,
		Domain:          param.Domain,
		FirstAccessTime: currentTime,
		LastAccessTime:  currentTime,
		ClientIP:        param.ClientIP,
		SessionID:       param.SessionID,
		UserAgent:       param.UserAgent,
		Referer:         param.Referer,
		ClientCountry:   city.CountryName,
	}
	if err != nil {
		domainLog.Upsert(true, param.PageType)
		userLog.Upsert(true)
	} else {
		domainLog.Upsert(false, param.PageType)
		userLog.Upsert(false)
	}

	detailData := mongoModel.UserAccessDetail{
		CreateTime:    currentTime,
		Domain:        param.Domain,
		Uri:           param.Uri,
		ClientIP:      param.ClientIP,
		SessionID:     param.SessionID,
		UserAgent:     param.UserAgent,
		Referer:       param.Referer,
		PageType:      param.PageType,
		ClientCountry: city.CountryName,
	}
	detailData.Add()

}

func (u UserAccessService) GetList(param params.AccessUserLogGetParam) (*[]mongoModel.UserAccessLog, int64) {
	var cursor *mongo.Cursor
	var err error
	var accessLogList []mongoModel.UserAccessLog
	var accessLog mongoModel.UserAccessLog
	model := mongoModel.UserAccessLog{}
	collection := model.GetCollection()
	dateStart := param.GetStartLocalTime()
	dateEnd := param.GetEndLocalTime()
	filter := map[string]interface{}{
		"date": bson.M{"$gte": dateStart, "$lte": dateEnd},
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
