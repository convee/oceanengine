package thirdsite

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// DeleteRequest 删除第三方落地页站点 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// DeleteResponse 删除第三方落地页站点 API Response
type DeleteResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *DeleteResponseData `json:"data,omitempty"`
}

// DeleteResponseData json返回值
type DeleteResponseData struct {
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}
