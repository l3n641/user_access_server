package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"user_accerss_server/api/params"
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
