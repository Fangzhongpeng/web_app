package redis

import (
	"web_app/models"
)

func GetPostIdInOrder(p *models.ParamPostList) ([]string, error) {
	//1 ，从redis获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2 ,确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3 ZRevRange查询,哪找分数从大到小顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}
