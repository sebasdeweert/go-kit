package types

import "time"

func PString(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}

func PInt64(v int64) *int64 {
	if v == 0 {
		return nil
	}

	return &v
}

func PTime(v time.Time) *time.Time {
	if v == time.Date(1970, 1, 1, 0,0,0,0, time.UTC) {
		return nil
	}

	return &v
}