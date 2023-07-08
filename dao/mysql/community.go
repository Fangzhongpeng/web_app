package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

//	func GetCommunityList() (communityList []*models.Community, err error) {
//		sqlStr := "select community_id,community_name from community"
//		if err := db.Select(&communityList, sqlStr); err != nil {
//			if err == sql.ErrNoRows {
//				zap.L().Warn("there is no community in db")
//				err = nil
//			}
//		}
//		return
//	}
func GetCommunityList() (communityList []*models.Community, err error) {
	//communityList是一个切片类型，切片是引用类型，不需要使用new函数初始化分配内存
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	//社区详情可能很大，返回值使用指针传递，需要使用new函数分配一个指针地址
	community = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name ,introduction, create_time from community where community_id = ?"
	//查询单挑数据用db.Get
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no communitydatail in db")
			err = ErrorInvalidID
		}
	}
	return
}
