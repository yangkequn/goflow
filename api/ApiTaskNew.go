package api

import (
	"context"
	"errors"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/yangkequn/saavuu/config"
	"github.com/yangkequn/saavuu/data"
	"github.com/yangkequn/saavuu/logger"
)

type fn func(dc *data.Ctx, pc *Ctx, paramIn map[string]interface{}) (out map[string]interface{}, err error)

var ErrBackTo = errors.New("param[\"backTo\"] is not a string")

func NewApi(serviceName string, f fn) {
	serviceName = "api:" + serviceName
	//check configureation is loaded
	if config.DataRds == nil {
		logger.Lshortfile.Panic("config.DataRedis is nil. you should call config.LoadConfigFromRedis first")
	}
	if config.ParamRds == nil {
		logger.Lshortfile.Panic("config.ParamRedis is nil. you should call config.LoadConfigFromRedis first")
	}
	ProcessOneJob := func(BackToID string, s []byte) (err error) {
		var (
			out            interface{}
			marshaledBytes []byte
			param          map[string]interface{} = map[string]interface{}{}
		)
		if err = msgpack.Unmarshal(s, &param); err != nil {
			return err
		}
		//process one job
		dc := &data.Ctx{Ctx: context.Background(), Rds: config.DataRds}
		pc := &Ctx{Ctx: context.Background(), Rds: config.ParamRds}
		if out, err = f(dc, pc, param); err != nil {
			return err
		}
		//Post Back
		if marshaledBytes, err = msgpack.Marshal(out); err != nil {
			return err
		}
		ctx := context.Background()
		pipline := config.ParamRds.Pipeline()
		pipline.RPush(ctx, BackToID, marshaledBytes)
		pipline.Expire(ctx, BackToID, time.Second*6)
		_, err = pipline.Exec(ctx)
		return err
	}
	apiServices[serviceName] = &ApiInfo{
		ApiName: serviceName,
		ApiFunc: ProcessOneJob,
	}
}