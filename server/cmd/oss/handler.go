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
	"time"
)

// OSSServiceImpl implements the last service interface defined in the IDL.
type OSSServiceImpl struct{}

// CreateOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) CreateOSS(ctx context.Context, req *oss.CreateOSSRequest) (resp *oss.CreateOSSResponse, err error) {
	var ossRecord model.OSSRecord
	sf, err := snowflake.NewNode(consts.OSSSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	ossRecord.Path = fmt.Sprintf("%s/%s", req.Path, sf.Generate().String())
	err = config.DB.Model(&model.OSSRecord{}).Create(&ossRecord).Error
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
	resp.Id = ossRecord.ID
	resp.BaseResp = tools.BuildBaseResp(200, "Create url returns success")
	return
}

// GetOSS implements the OSSServiceImpl interface.
func (s *OSSServiceImpl) GetOSS(ctx context.Context, req *oss.GetOSSRequest) (resp *oss.GetOSSResponse, err error) {
	// TODO: Your code here...
	return
}
