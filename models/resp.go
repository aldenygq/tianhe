package models

type RespOncallRules struct{
	Count int64 `json:"count"`
	Rules []*OncallRule `json:"rules"`
}
type RespDutyPerson struct {
	Rule string `json:"rule"`
	Person []string `json:"person"`
	RuleName string `json:"rule_name"`
}
