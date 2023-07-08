package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// func CreatePost(p *models.Post) (err error) {
//
//		// 1. 生成post id
//		p.ID = snowflake.GenID()
//		//2 保存到数据库
//		//post := &models.Post{
//		//	ID:          p.ID,
//		//	AuthorID:    p.AuthorID,
//		//	CommunityID: p.CommunityID,
//		//	Title:       p.Title,
//		//	Content:     p.Content,
//		//}
//		//3 保存到数据库
//
//		return mysql.CreatePost(p)
//	}
func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 3. 返回
}

// GetPostById 根据帖子id查询帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//查询并组合我们接口想用的数据

	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid)", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	//根据作者id查作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByid(post.AuthorID)", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
		return
	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID)", zap.Int64("post.CommunityID", post.CommunityID), zap.Error(err))
		return
	}
	//拼装数据
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	//
	//data.AuthorName = user.Username
	//data.CommunityDetail = community
	return
}

// GetPostList 获取帖子列表
//
//	func GetPostList(uid int64) (data []*models.Post, err error) {
//		return mysql.GetPostList()
//	}
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 2 去redis查询id列表
	ids, err := redis.GetPostIdInOrder(p)
	if err != nil {
		return
	}
	// 3 根据id去mysql数据库查询帖子详情
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdInOrder(p) return 0 data")
		return
	}

	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return
	}
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
