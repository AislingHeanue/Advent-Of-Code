package util

import (
	"cmp"
	"fmt"
)

func Between[V cmp.Ordered](minimum, value, maximum V) V {
	if maximum < minimum {
		fmt.Printf("ERROR: max %v is less than min %v", maximum, minimum)
	}
	return max(min(value, maximum), minimum)
}
