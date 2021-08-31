package types

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func LString(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

func LInt64(v *int64) int64 {
	if v == nil {
		return 0
	}

	return *v
}

func LPBTimestamp(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}

	return timestamppb.New(*v)
}

func LBool(v *bool) bool {
	if v == nil {
		return false
	}

	return *v
}