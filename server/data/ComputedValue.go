package data

import (
	"encoding/json"
	"strings"
)

type ComputedValue struct {
	InternalId int             `json:"-"`
	ValueId    string          `json:"valueId"`
	Key        string          `json:"key"`
	Data       json.RawMessage `json:"data"`
}

func createKey(parts ...string) string {
	return strings.Join(parts, "__")
}

func ComputedValueKeyGlobalTitleMetrics() string {
	return "global-title-metrics"
}

func ComputedValueKeyAgencyMetric(agencyId string) string {
	return createKey("agency-metrics", agencyId)
}
