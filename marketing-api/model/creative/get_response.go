package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// GetResponse 获取创意列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List Creative数组
	List []Creative `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Creative 创意
type Creative struct {
	// CreativeID 创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Title 创意素材标题
	Title string `json:"title,omitempty"`
	// CreativeWordIDs 动态词包列表
	CreativeWordIDs []uint64 `json:"creative_word_ids,omitempty"`
	// Status 创意素材状态
	Status enum.CreativeStatus `json:"status,omitempty"`
	// OpStatus 创意素材操作状态
	OpStatus enum.CreativeOpStatus `json:"op_status,omitempty"`
	// ImageMode 创意素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageIDs 图片素材，图片ID列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// ImageID 视频素材，封面图片ID
	ImageID string `json:"image_id,omitempty"`
	// VideoID 视频素材，视频ID
	VideoID string `json:"video_id,omitempty"`
	// PlayableInfo 基础试玩素材信息，使用基础试玩素材时返回
	PlayableInfo *PlayableInfo `json:"playable_info,omitempty"`
	// DerivePosterCID 是否将视频的封面和标题同步到图片创意，1为开启，0为不开启。视频素材时返
	DerivePosterCID int `json:"derive_poster_cid,omitempty"`
	// ThirdPartyID 第三方ID
	ThirdPartyID string `json:"third_party_id,omitempty"`
	// DpaDictIDs DPA词包ID列表，针对DPA广告
	DpaDictIDs []uint64 `json:"dpa_dict_ids,omitempty"`
	// TemplateID DPA模板ID，针对DPA广告
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateImageID DPA创意实际显示的图片ID，针对DPA广告
	TemplateImageID string `json:"template_image_id,omitempty"`
	// DpaTemplate 是否使用商品库视频模板，针对DPA广告
	DpaTemplate uint64 `json:"dpa_template,omitempty"`
	// DpaVideoTempateType 商品库视频模板生成类型，针对DPA广告
	DpaVideoTempateType string `json:"dpa_video_template_type,omitempty"`
	// DpaVideoTaskIDs 自定义商品库视频模板ID，针对DPA广告
	DpaVideoTaskIDs []string `json:"dpa_video_task_ids,omitempty"`
	// ComponentInfo 创意组件信息
	ComponentInfo []struct {
		// ComponentID 创意组件id
		ComponentID string `json:"component_id,omitempty"`
	} `json:"component_info,omitempty"`
	// Materials 素材信息列表，标题，图片，视频均是不同素材，注意部分老数据可能此结构为空
	Materials []Material `json:"materials,omitempty"`
	// CreativeCreateTime 广告创意创建时间，格式yyyy-MM-dd HH:mm:ss
	CreativeCreateTime string `json:"creative_create_time,omitempty"`
	// CreativeModifyTime 广告创意更新时间，格式yyyy-MM-dd HH:mm:ss
	CreativeModifyTime string `json:"creative_modity_time,omitempty"`
}
