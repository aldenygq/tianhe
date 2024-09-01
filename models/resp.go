package models

type RespOncallRules struct{
	Count int64 `json:"count"`
	Rules []*OncallRuleInfo `json:"rules"`
}
type RespDutyPerson struct {
	Rule string `json:"rule"`
	Person []string `json:"person"`
	RuleName string `json:"rule_name"`
}
type RespCreateTime struct{
	CreateTime int64 `json:"create_time"`
}
type RespUpdateTime struct{
	UpdateTime int64 `json:"update_time"`
}
type OncallRuleInfo struct {
	ParamOncallRule
	CreatorInfo
	UpdatorInfo
	ParamOncallInfo
	RespCreateTime
	RespUpdateTime
}
