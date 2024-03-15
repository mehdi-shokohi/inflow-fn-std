package db

import (
	"context"

	conf "github.com/mehdi-shokohi/inflow-fn-std/config"
	"github.com/mehdi-shokohi/mongoHelper"
)

func GetDefaultDb[T any](ctx context.Context, colName string, model T) mongoHelper.MongoContainer[T] {
	dbId := "dafault"
	uri := conf.Get_env_value("MONGO_DB_URI")
	inflowDatabase := conf.Get_env_value("DB")
	db := mongoHelper.NewMongoById[T](ctx, dbId, uri, inflowDatabase, colName, model)
	return db
}
