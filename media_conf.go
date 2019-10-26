package service

import (
	"strings"
	"util/errors"
	"yunGame/web/model"

	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
)

const (
	Open     = 1  // 打开
	Close    = 0  // 关闭
	NotExist = -1 // 表示字段不存在
)

type mediaService struct{}

var MediaService = new(mediaService)

func (*mediaService) Get(model1, version string) (*model.MediaConfStr, error) {
	decodeConf, err := model.GetDecodeConf(model1)
	if err != nil {
		appzaplog.Error("mediaService Get error", zap.Error(err))
		return nil, err
	}

	mediaConf, err := model.GetMediaConf(model1, version)
	if err != nil {
		appzaplog.Error("mediaService Get error", zap.Error(err))
		return nil, err
	}

	switch {
	case decodeConf == nil:
		if mediaConf != nil {
			return &model.MediaConfStr{
				IsYuvRender:        mediaConf.IsYuvRender,
				IsDisableH265Codec: mediaConf.IsDisableH265Codec,
				IsSwDecoder:        Close,
			}, nil
		}

		return nil, errors.ErrMediaConfNotFound
	case decodeConf.IsOpen == Open:
		return &model.MediaConfStr{
			IsYuvRender:        Open,
			IsDisableH265Codec: Open,
			IsSwDecoder:        Open,
		}, nil
	case decodeConf.IsOpen == Close:
		if mediaConf != nil {
			return &model.MediaConfStr{
				IsYuvRender:        mediaConf.IsYuvRender,
				IsDisableH265Codec: mediaConf.IsDisableH265Codec,
				IsSwDecoder:        Close,
			}, nil
		}
		return &model.MediaConfStr{
			IsYuvRender:        NotExist,
			IsDisableH265Codec: NotExist,
			IsSwDecoder:        Close,
		}, nil
	}
	return nil, nil
}

func HandleVersion(version string) string {
	if len(strings.Split(version, ".")) == 1 {
		return version + ".0.0"
	}
	return version
}
