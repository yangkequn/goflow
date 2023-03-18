package data

import (
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack/v5"
)

func (db *Ctx) ZAdd(members ...redis.Z) (err error) {
	MarshalRedisZ(members...)
	status := db.Rds.ZAdd(db.Ctx, db.Key, members...)
	return status.Err()
}
func (db *Ctx) ZRem(members ...interface{}) (err error) {
	//msgpack marshal members
	var memberBytes [][]byte
	if memberBytes, err = MarshalSlice(members...); err != nil {
		return err
	}
	status := db.Rds.ZRem(db.Ctx, db.Key, memberBytes)
	return status.Err()
}
func (db *Ctx) ZRange(start, stop int64, outSlice interface{}) (err error) {
	var cmd *redis.StringSliceCmd

	if cmd = db.Rds.ZRange(db.Ctx, db.Key, start, stop); cmd.Err() != nil && cmd.Err() != redis.Nil {
		return cmd.Err()
	}
	return UnmarshalStrings(cmd.Val(), outSlice)
}
func (db *Ctx) ZRangeWithScores(start, stop int64, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZRangeWithScores(db.Ctx, db.Key, start, stop)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZRevRangeWithScores(start, stop int64, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZRevRangeWithScores(db.Ctx, db.Key, start, stop)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZRank(member string) (rank int64, err error) {
	var (
		memberBytes []byte
	)
	//marshal member using msgpack
	if memberBytes, err = msgpack.Marshal(member); err != nil {
		return 0, err
	}
	cmd := db.Rds.ZRank(db.Ctx, db.Key, string(memberBytes))
	return cmd.Val(), cmd.Err()
}
func (db *Ctx) ZRevRank(member string) (rank int64, err error) {
	var (
		memberBytes []byte
	)
	//marshal member using msgpack
	if memberBytes, err = msgpack.Marshal(member); err != nil {
		return 0, err
	}
	cmd := db.Rds.ZRevRank(db.Ctx, db.Key, string(memberBytes))
	return cmd.Val(), cmd.Err()
}
func (db *Ctx) ZScore(member string) (score float64, err error) {
	var (
		memberBytes []byte
		cmd         *redis.FloatCmd
	)
	//marshal member using msgpack
	if memberBytes, err = msgpack.Marshal(member); err != nil {
		return 0, err
	}
	if cmd = db.Rds.ZScore(db.Ctx, db.Key, string(memberBytes)); cmd.Err() != nil && cmd.Err() != redis.Nil {
		return 0, err
	} else if cmd.Err() == redis.Nil {
		return 0, nil
	}
	return cmd.Result()
}
func (db *Ctx) ZCard() (length int64, err error) {
	cmd := db.Rds.ZCard(db.Ctx, db.Key)
	return cmd.Result()
}
func (db *Ctx) ZCount(min, max string) (length int64, err error) {
	cmd := db.Rds.ZCount(db.Ctx, db.Key, min, max)
	return cmd.Result()
}
func (db *Ctx) ZRangeByScore(opt *redis.ZRangeBy, outSlice interface{}) (err error) {
	cmd := db.Rds.ZRangeByScore(db.Ctx, db.Key, opt)
	return UnmarshalStrings(cmd.Val(), outSlice)
}
func (db *Ctx) ZRangeByScoreWithScores(opt *redis.ZRangeBy, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZRangeByScoreWithScores(db.Ctx, db.Key, opt)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZRevRangeByScore(opt *redis.ZRangeBy, outSlice interface{}) (err error) {
	cmd := db.Rds.ZRevRangeByScore(db.Ctx, db.Key, opt)
	return UnmarshalStrings(cmd.Val(), outSlice)
}
func (db *Ctx) ZRevRange(start, stop int64, outSlice interface{}) (err error) {
	var cmd *redis.StringSliceCmd

	if cmd = db.Rds.ZRevRange(db.Ctx, db.Key, start, stop); cmd.Err() != nil && cmd.Err() != redis.Nil {
		return cmd.Err()
	}
	return UnmarshalStrings(cmd.Val(), outSlice)
}
func (db *Ctx) ZRevRangeByScoreWithScores(opt *redis.ZRangeBy, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZRevRangeByScoreWithScores(db.Ctx, db.Key, opt)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZRemRangeByRank(start, stop int64) (err error) {
	status := db.Rds.ZRemRangeByRank(db.Ctx, db.Key, start, stop)
	return status.Err()
}
func (db *Ctx) ZRemRangeByScore(min, max string) (err error) {
	status := db.Rds.ZRemRangeByScore(db.Ctx, db.Key, min, max)
	return status.Err()
}
func (db *Ctx) ZIncrBy(increment float64, member interface{}) (err error) {
	var (
		memberBytes []byte
	)
	//marshal member using msgpack
	if memberBytes, err = msgpack.Marshal(member); err != nil {
		return err
	}
	status := db.Rds.ZIncrBy(db.Ctx, db.Key, increment, string(memberBytes))
	return status.Err()
}
func (db *Ctx) ZPopMax(count int64, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZPopMax(db.Ctx, db.Key, count)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZPopMin(count int64, outSlice interface{}) (scores []float64, err error) {
	cmd := db.Rds.ZPopMin(db.Ctx, db.Key, count)
	return UnmarshalRedisZ(cmd.Val(), outSlice)
}
func (db *Ctx) ZLexCount(min, max string) (length int64) {
	cmd := db.Rds.ZLexCount(db.Ctx, db.Key, min, max)
	return cmd.Val()
}
func (db *Ctx) ZScan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	cmd := db.Rds.ZScan(db.Ctx, db.Key, cursor, match, count)
	return cmd.Result()
}
