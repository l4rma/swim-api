package utils

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ParseInt(attr types.AttributeValue) int {
	if attr == nil {
		return 0
	}
	val, err := strconv.Atoi(attr.(*types.AttributeValueMemberN).Value)
	if err != nil {
		return 0
	}
	return val
}

func ParseTime(attr types.AttributeValue) time.Time {
	if attr == nil {
		return time.Time{}
	}
	parsed, err := time.Parse("2006-01-02T15:04:05Z", attr.(*types.AttributeValueMemberS).Value)
	if err != nil {
		return time.Time{}
	}
	return parsed
}

func ParseBool(attr types.AttributeValue) bool {
	if attr == nil {
		return false
	}
	return attr.(*types.AttributeValueMemberBOOL).Value
}

func ParseString(attr types.AttributeValue) string {
	if attr == nil {
		return ""
	}
	return attr.(*types.AttributeValueMemberS).Value
}
