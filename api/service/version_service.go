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

	responseTime := fmt.Sprintf("%d/%d/%d %d:%d:%d", ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute(), ct.Second())

	versionData := model.VersionResponse{
		Name:          config.API_NAME,
		BuildDateTime: responseTime,
		Version:       config.API_VERSION,
	}

	response := response.Response{
		Code: 200,
		Data: versionData,
	}

	return response
}
