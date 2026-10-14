package service

import (
	"fmt"
	"go-admin/common/config"
	"go-admin/common/result"
	"go-admin/common/util"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUploadService interface {
	Upload(c *gin.Context)
}

type UploadServiceImpl struct{}

func (u *UploadServiceImpl) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.UPLOADERROR), result.ApiCode.GetMessage(result.ApiCode.UPLOADERROR))
		return
	}

	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf(
		"%s%s%s%s",
		config.Config.ImgSetting.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()),
	)
	util.CreateDir(filePath)
	fullPath := filePath + "/" + fileName
	c.SaveUploadedFile(file, fullPath)
	result.Success(c, config.Config.ImgSetting.ImgHost+fullPath)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}
