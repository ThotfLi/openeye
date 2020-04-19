package session

import (
	"fmt"
	"testing"
)

var newSesManger ISesManger
func TestNewSesManger(t *testing.T) {
	newSesManger = NewSesManger()
}

func testSesManger_GetSession(t *testing.T) {
	if _,err := newSesManger.GetSession("31fd5387-57ed-402e-9e29-82482a30296d"); err != nil {
		t.Error(err)
	}
}

func testSesManger_AddMap(t *testing.T) {
	if err := newSesManger.AddMap("31fd5387-57ed-402e-9e29-82482a30296d",newsession); err != nil {
		t.Error(err)
	}
}

func testSesManger_IsExpiration(t *testing.T) {
	if _,err := newSesManger.IsExpiration("31fd5387-57ed-402e-9e29-82482a30296d"); err != nil {
		t.Error(err)
	}
	b,_ := newSesManger.IsExpiration("31fd5387-57ed-402e-9e29-82482a30296d")
	if b {
		fmt.Println("已过期")
	}
}
