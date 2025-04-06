package utils

import (
	"strings"

	"github.com/himaatluri/tagion/pkg/types"
)

func SupportsTags(resourceType string) bool {
	supportedTypes := []string{
		"AWS::EC2::",
		"AWS::S3::",
		"AWS::RDS::",
		"AWS::DynamoDB::",
		"AWS::Lambda::",
	}

	for _, t := range supportedTypes {
		if strings.HasPrefix(resourceType, t) {
			return true
		}
	}
	return false
}

func CreateNewTags(configTags map[string]string) []types.Tags {
	newTags := make([]types.Tags, 0)
	for k, v := range configTags {
		newTags = append(newTags, types.Tags{Key: k, Value: v})
	}
	return newTags
}

func HasTag(tags []types.Tags, key string) bool {
	for _, tag := range tags {
		if tag.Key == key {
			return true
		}
	}
	return false
}

func ConvertToTagsArray(tags interface{}) []types.Tags {
	var result []types.Tags

	switch v := tags.(type) {
	case []interface{}:
		for _, tag := range v {
			if m, ok := tag.(map[string]interface{}); ok {
				result = append(result, types.Tags{
					Key:   m["Key"].(string),
					Value: m["Value"].(string),
				})
			}
		}
	case []types.Tags:
		result = v
	}

	return result
}

func MergeTags(existingTags []types.Tags, newTags map[string]string) []types.Tags {
	modified := false
	for k, v := range newTags {
		if !HasTag(existingTags, k) {
			existingTags = append(existingTags, types.Tags{Key: k, Value: v})
			modified = true
		}
	}
	if modified {
		return existingTags
	}
	return nil
}
