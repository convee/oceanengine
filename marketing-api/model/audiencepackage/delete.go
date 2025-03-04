package audiencepackage

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// DeleteRequest 删除定向包 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudiencePackageID 删除的定向包id
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// DeleteResponse 删除定向包 API Response
type DeleteResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	} `json:"data,omitempty"`
}
