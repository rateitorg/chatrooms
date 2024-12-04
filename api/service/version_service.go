package service

import (
	"fmt"
	"time"

	"github.com/rateitorg/chatrooms/api/domain/model"
	"github.com/rateitorg/chatrooms/api/domain/response"
	"github.com/rateitorg/chatrooms/config"
)

type VersionService struct {
}

func (vs *VersionService) GetVersionService() response.Response {
	ct := time.Now() // Get current time
	cfg := config.GetConfig()

	responseTime := fmt.Sprintf("%d/%d/%d %d:%d:%d", ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute(), ct.Second())

	versionData := model.VersionModel{
		Name:          cfg.APIName,
		BuildDateTime: responseTime,
		Version:       cfg.APIVersion,
	}

	response := response.Response{
		Code: 200,
		Data: versionData,
	}

	return response
}
