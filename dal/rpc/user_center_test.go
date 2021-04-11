package rpc

import (
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	resp, err := GetUserInfo(ctx, 1, "")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v", resp)
}
