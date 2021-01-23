package ipInfo

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	res := GetInfo("47.96.154.4a")
	fmt.Println(res)
}