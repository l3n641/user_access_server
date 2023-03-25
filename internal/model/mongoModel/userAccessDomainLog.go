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

type UserAccessDomainLog struct {
	Id            primitive.ObjectID `bson:"_id"`
	Date          time.Time          `bson:"date"`       //记录的日期
	Domain        string             `bson:"domain"`     //域名
	PageViews     uint               `bson:"page_views"` //pv 次数
	HomePage      uint               `bson:"home_page"`
	UniqueVisitor uint               `bson:"unique_visitor"`
	CategoryPage  uint               `bson:"category_page"`
	SingularPage  uint               `bson:"singular_page"`
	TagPage       uint               `bson:"tag_page"`
	CardPage      uint               `bson:"cart_page"`
	CheckoutPage  uint               `bson:"checkout_page"`
	AccountPage   uint               `bson:"account_page"`
}

func (p *UserAccessDomainLog) GetCollection() *mongo.Collection {
	db := mongoDb.GetMongoDb()
	collection := db.Collection("user_access_domain_log")
	return collection
}

func (p *UserAccessDomainLog) getBsonByPageType(pageType string) bson.E {
	switch pageType {
	case "home":
		return bson.E{Key: "$inc", Value: bson.D{{
			"home_page", 1,
		}}}

	case "category":
		return bson.E{Key: "$inc", Value: bson.D{{
			"category_page", 1,
		}}}

	case "singular":
		return bson.E{Key: "$inc", Value: bson.D{{
			"singular_page", 1,
		}}}

	case "tag":
		return bson.E{Key: "$inc", Value: bson.D{{
			"tag_page", 1,
		}}}

	case "cart":
		return bson.E{Key: "$inc", Value: bson.D{{
			"cart_page", 1,
		}}}

	case "checkout":
		return bson.E{Key: "$inc", Value: bson.D{{
			"checkout_page", 1,
		}}}

	case "account":
		return bson.E{Key: "$inc", Value: bson.D{{
			"account_page", 1,
		}}}

	default:
		return bson.E{}
	}

}

func (p *UserAccessDomainLog) Upsert(isNew bool, pageType string) error {
	var update bson.D
	collection := p.GetCollection()
	updateOptions := options.Update().SetUpsert(true)
	filter := bson.M{"date": p.Date, "domain": p.Domain}

	update = bson.D{
		{
			"$inc", bson.D{{ // $inc 代表增加或减少
				"page_views", 1, // 在原值的基础上 +1
			}},
		},
	}

	if isNew {
		update = append(update, bson.E{Key: "$inc", Value: bson.D{{
			"unique_visitor", 1,
		}}})
	}

	pageInc := p.getBsonByPageType(pageType)
	if pageInc.Key != "" {
		update = append(update, pageInc)
	}

	// 执行 upsert 操作
	_, err := collection.UpdateOne(context.TODO(), filter, update, updateOptions)
	return err

}
