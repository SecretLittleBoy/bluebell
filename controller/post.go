package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bluebell/logic"
	"bluebell/models"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	userId, _, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userId
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	post, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("logic.GetPostById() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

func GetPostListHandler(c *gin.Context) {
	pageNum, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.ParseInt(c.Query("size"), 10, 64)
	if err != nil {
		pageSize = 10
	}
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /v2/post [get]
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderByTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetCommunityPostListHandler(c *gin.Context) {
	communityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := &models.ParamCommunityPostList{
		ParamPostList: models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderByTime,
		},
		CommunityID: communityID,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
