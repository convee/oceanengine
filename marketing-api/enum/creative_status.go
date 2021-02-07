package enum

// 创意状态
type CreativeStatus string

const (
	CREATIVE_STATUS_DELIVERY_OK              CreativeStatus = "CREATIVE_STATUS_DELIVERY_OK"              // 投放中
	CREATIVE_STATUS_NOT_START                CreativeStatus = "CREATIVE_STATUS_NOT_START"                // 未到达投放时间
	CREATIVE_STATUS_NO_SCHEDULE              CreativeStatus = "CREATIVE_STATUS_NO_SCHEDULE"              // 不在投放时段
	CREATIVE_STATUS_DISABLE                  CreativeStatus = "CREATIVE_STATUS_DISABLE"                  // 创意暂停
	CREATIVE_STATUS_CAMPAIGN_DISABLE         CreativeStatus = "CREATIVE_STATUS_CAMPAIGN_DISABLE"         // 已被广告组暂停
	CREATIVE_STATUS_CAMPAIGN_EXCEED          CreativeStatus = "CREATIVE_STATUS_CAMPAIGN_EXCEED"          // 广告组超出预算
	CREATIVE_STATUS_AUDIT                    CreativeStatus = "CREATIVE_STATUS_AUDIT"                    // 新建审核中
	CREATIVE_STATUS_REAUDIT                  CreativeStatus = "CREATIVE_STATUS_REAUDIT"                  // 修改审核中
	CREATIVE_STATUS_DELETE                   CreativeStatus = "CREATIVE_STATUS_DELETE"                   // 已删除
	CREATIVE_STATUS_DONE                     CreativeStatus = "CREATIVE_STATUS_DONE"                     // 已完成（投放达到结束时间）
	CREATIVE_STATUS_AD_DISABLE               CreativeStatus = "CREATIVE_STATUS_AD_DISABLE"               // 广告计划暂停
	CREATIVE_STATUS_AUDIT_DENY               CreativeStatus = "CREATIVE_STATUS_AUDIT_DENY"               // 审核不通过
	CREATIVE_STATUS_BALANCE_EXCEED           CreativeStatus = "CREATIVE_STATUS_BALANCE_EXCEED"           // 账户余额不足
	CREATIVE_STATUS_BUDGET_EXCEED            CreativeStatus = "CREATIVE_STATUS_BUDGET_EXCEED"            // 超出预算
	CREATIVE_STATUS_DATA_ERROR               CreativeStatus = "CREATIVE_STATUS_DATA_ERROR"               // 数据错误（数据错误时返回，极少出现）
	CREATIVE_STATUS_PRE_ONLINE               CreativeStatus = "CREATIVE_STATUS_PRE_ONLINE"               // 预上线
	CREATIVE_STATUS_AD_AUDIT                 CreativeStatus = "CREATIVE_STATUS_AD_AUDIT"                 // 广告计划新建审核中
	CREATIVE_STATUS_AD_REAUDIT               CreativeStatus = "CREATIVE_STATUS_AD_REAUDIT"               // 广告计划修改审核中
	CREATIVE_STATUS_AD_AUDIT_DENY            CreativeStatus = "CREATIVE_STATUS_AD_AUDIT_DENY"            // 广告计划审核不通过
	CREATIVE_STATUS_ALL                      CreativeStatus = "CREATIVE_STATUS_ALL"                      // 所有包含已删除
	CREATIVE_STATUS_NOT_DELETE               CreativeStatus = "CREATIVE_STATUS_NOT_DELETE"               // 所有不包含已删除（状态过滤默认值）
	CREATIVE_STATUS_ADVERTISER_BUDGET_EXCEED CreativeStatus = "CREATIVE_STATUS_ADVERTISER_BUDGET_EXCEED" // 超出账户日预算
)
