package dbops

import (
	"fmt"
	"testing"
)

func testAddUserCreate(t *testing.T) {
	err := AddUserCreate("lx","1234567")
	if err != nil{
		t.Error("ADD fail err:",err)
	}
}

func testGetUser(t *testing.T) {
	_,err := GetUser("lx")
	if err != nil {
		t.Error("GET FAIL err:",err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("lx","1234567")
	if err != nil {
		t.Error(err)
	}
}

func testAddVideo(t *testing.T) {
	if err := AddVideo(1,"西红柿首富"); err != nil {
		t.Fail()
	}
}

func testGetVideo(t *testing.T) {
	if _,err := GetVideo("31d621c5-23e8-48b5-bbec-2a32e1a8787e")	;err!= nil{
		t.Fail()
	}
}

func testDelVideo(t *testing.T) {
	if err := DelVideo("31d621c5-23e8-48b5-bbec-2a32e1a8787e"); err != nil{
		t.Fail()
	}
}

func testAddCommonts(t *testing.T) {
	for i:=0;i<10;i++ {
		//生成10条评论，注意在其他地方输入的测试数据可能需要重新生成
		if err:=AddCommont("c882b4ae-cd4d-453c-939f-34c98655c8c1", 1, "这是我的回答");err != nil{
			t.Fail()
		}
	}
}

func testGetCommonts(t *testing.T) {
	comms,err := GetCommonts("c882b4ae-cd4d-453c-939f-34c98655c8c1")
	if err != nil || len(comms)!=10{
		t.Fail()
	}
	fmt.Println(comms)
}

func testDelCommont(t *testing.T) {
	if err := DelCommont("fde6d106-5a73-4400-83fa-97fb13605292");err != nil{
		t.Fail()
	}
}
func TestALL(t *testing.T){
	//t.Run("add",testAddUserCreate)
	//t.Run("GET",testGetUser)
	//t.Run("DEL",testDeleteUser)
	//t.Run("videoADD",testAddVideo)
	//t.Run("GetVideo",testGetVideo)
	//t.Run("DelVideo",testDelVideo)
	//t.Run("AddCommonts",testAddCommonts)
	//t.Run("GetCommonts",testGetCommonts)
	//t.Run("DelComments",testDelCommont)
}
