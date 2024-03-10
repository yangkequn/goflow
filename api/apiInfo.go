package api

import (
	"context"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/yangkequn/goflow/config"
)

type ApiInfo struct {
	// Name is the name of the service
	Name       string
	DataSource string
	WithHeader bool
	Ctx        context.Context
	// ApiFuncWithMsgpackedParam is the function of the service
	ApiFuncWithMsgpackedParam func(s []byte) (ret interface{}, err error)
}

var ApiServices cmap.ConcurrentMap[string, *ApiInfo] = cmap.New[*ApiInfo]()

func apiServiceNames() (serviceNames []string) {
	for _, serviceInfo := range ApiServices.Items() {
		serviceNames = append(serviceNames, serviceInfo.Name)
	}
	return serviceNames
}
func GetServiceDB(serviceName string) (db *redis.Client) {
	var (
		ok bool
	)
	serviceInfo, _ := ApiServices.Get(serviceName)
	DataSource := serviceInfo.DataSource
	if db, ok = config.Rds[DataSource]; !ok {
		log.Panic().Str("DataSource not defined in enviroment. Please check the configuration", DataSource).Send()
	}
	return db
}

var fun2ApiInfoMap = &sync.Map{}
var APIGroupByDataSource = cmap.New[[]string]()