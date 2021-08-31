package array

import (
	"strconv"
	"strings"
)

func CommaSeparatedIDs(IDs []int64) string {
	var sIDs []string

	for _, v := range IDs {
		sIDs = append(sIDs, strconv.FormatInt(v, 10))
	}

	return strings.Join(sIDs, ",")
}
