package data

import (
	"encoding/json"
	"strings"
	"unicode"
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

func sanitize(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		} else if unicode.IsSpace(r) {
			result = append(result, '-')
		}
	}
	return string(result)
}

func ComputedValueKeyGlobalTitleMetrics() string {
	return "global-title-metrics"
}

func ComputedValueKeyAgencyMetric(agencyId string) string {
	return createKey("agency-metrics", agencyId)
}

func ComputedValueKeySubAgencyMetric(parentId string, agencyName string) string {
	agencyId := strings.ToLower(agencyName)
	return createKey("sub-agency-metrics", parentId, sanitize(agencyId))
}
