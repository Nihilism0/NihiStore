package main

import (
	"NihiStore/server/cmd/oss/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/oss"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"time"
)

// OSSServiceImpl implements the last service interface defined in the IDL.
type OSSServiceImpl struct{}

// CreateGoodsOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) CreateGoodsOSS(ctx context.Context, req *oss.CreateGoodsOSSRequest) (resp *oss.CreateGoodsOSSResponse, err error) {
	resp = new(oss.CreateGoodsOSSResponse)
	var ossRecord model.GoodsOSSRecord
	sf, err := snowflake.NewNode(consts.OSSSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
		tools.BuildBaseResp(errx.SnowFlakeError, "SnowFlakeError")
		return resp, nil
	}
	ossRecord.GoodsId = req.GoodsId
	ossRecord.Path = fmt.Sprintf("%s/%s", req, sf.Generate().String())
	err = config.DB.Model(&model.HeadOSSRecord{}).Create(&ossRecord).Error
	if err != nil {
		klog.Error("presigned put object url err", err)
		return resp, nil
	}
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u, err := config.MinioClient.PresignedPutObject(ctx, "nihistore", req.Path, time.Duration(req.TimeoutSec)*time.Second)
	if err != nil {
		klog.Error("presigned put object url err", err)
		resp.BaseResp = tools.BuildBaseResp(errx.PutObjectError, "PutObjectError")
		return resp, nil
	}
	resp.UploadUrl = u.String()
	resp.Id = int64(ossRecord.ID)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Create url returns success")
	return resp, nil
}

// GetGoodsOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) GetGoodsOSS(ctx context.Context, req *oss.GetGoodsOSSRequest) (resp *oss.GetGoodsOSSResponse, err error) {
	resp = new(oss.GetGoodsOSSResponse)
	var ossRecord model.GoodsOSSRecord
	config.DB.Model(&model.GoodsOSSRecord{}).Where("id = ?", req.Id).First(&ossRecord)
	url, err := config.MinioClient.PresignedGetObject(ctx, "nihilism", ossRecord.Path, time.Duration(req.TimeoutSec)*time.Second, nil)
	resp.Url = url.String()
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Get URL success")
	return resp, nil
}

// CreateHeadOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) CreateHeadOSS(ctx context.Context, req *oss.CreateHeadOSSRequest) (resp *oss.CreateHeadOSSResponse, err error) {
	resp = new(oss.CreateHeadOSSResponse)
	var ossRecord model.HeadOSSRecord
	sf, err := snowflake.NewNode(consts.OSSSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
		tools.BuildBaseResp(errx.SnowFlakeError, "SnowFlakeError")
		return resp, nil
	}
	ossRecord.UserId = req.UserId
	ossRecord.Path = fmt.Sprintf("%s/%s", req, sf.Generate().String())
	err = config.DB.Model(&model.HeadOSSRecord{}).Create(&ossRecord).Error
	if err != nil {
		klog.Error("presigned put object url err", err)
		return resp, nil
	}
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u, err := config.MinioClient.PresignedPutObject(ctx, "nihistore", req.Path, time.Duration(req.TimeoutSec)*time.Second)
	if err != nil {
		klog.Error("presigned put object url err", err)
		resp.BaseResp = tools.BuildBaseResp(errx.PutObjectError, "PutObjectError")
		return resp, nil
	}
	resp.UploadUrl = u.String()
	resp.Id = int64(ossRecord.ID)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Create url returns success")
	return resp, nil
}

// GetHeadOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) GetHeadOSS(ctx context.Context, req *oss.GetHeadOSSRequest) (resp *oss.GetHeadOSSResponse, err error) {
	resp = new(oss.GetHeadOSSResponse)
	var ossRecord model.HeadOSSRecord
	config.DB.Model(&model.HeadOSSRecord{}).Where("id = ?", req.Id).First(&ossRecord)
	url, err := config.MinioClient.PresignedGetObject(ctx, "nihilism", ossRecord.Path, time.Duration(req.TimeoutSec)*time.Second, nil)
	resp.Url = url.String()
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Get URL success")
	return resp, nil
}
