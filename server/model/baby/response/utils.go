package response

import (
	"encoding/json"
	"fmt"
	"strings"
)

// formatDuration 格式化时长显示
func formatDuration(seconds int) string {
	if seconds <= 0 {
		return "00:00"
	}
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

// parseTags 解析标签
func parseTags(tagsStr string) []string {
	if tagsStr == "" {
		return []string{}
	}

	var tags []string
	if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
		// 如果JSON解析失败，尝试按逗号分割
		tags = strings.Split(tagsStr, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		// 过滤空字符串
		var filteredTags []string
		for _, tag := range tags {
			if tag != "" {
				filteredTags = append(filteredTags, tag)
			}
		}
		return filteredTags
	}
	return tags
}
