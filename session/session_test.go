package session

import (
	"fmt"
	"sync"
	"testing"
)

var newsession ISession

func TestMain(m *testing.M) {
	fmt.Println("test main first")
	newsession = &Session{SessionID:"31fd5387-57ed-402e-9e29-82482a30296d",sycMap:&sync.Map{}}

	m.Run()
}

func TestALL(t *testing.T) {
	t.Run("overdue",testOverdue)
	t.Run("attr",testAttr)
	//t.Run("isexp",testIsExp)
	//t.Run("del",DelAttr)

	t.Run("writeDB",testWriteDB)
	t.Run("readDB",testReadDB)
	t.Run("AddMap",testSesManger_AddMap)
	t.Run("GetSession",testSesManger_GetSession)
	t.Run("is过期",testSesManger_IsExpiration)
}

func testOverdue(t *testing.T){
	newsession.SetOverdueTime(1000)
	timestr := newsession.GetOverdueTime()
	fmt.Println(timestr)

}

func testAttr(t *testing.T){
	if err := newsession.SetAttr("lx","123"); err != nil {
		t.Errorf("setattr error:")
	}

	if err := newsession.SetAttr("xiang",456); err != nil {
		t.Errorf("setattr error:")
	}

	if str,err := newsession.GetAttr("lx"); err != nil {
		t.Errorf("getattr error")
	}else {
		fmt.Println(str.(string))
	}

	if str,err := newsession.GetAttr("xiang"); err != nil {
		t.Errorf("getattr error")
	}else{
		fmt.Println(str.(int))
	}

}

func DelAttr(t *testing.T){
	newsession.DelAttr("abc")
}

func testIsExp(t *testing.T){
	//time.Sleep(6*time.Second)
	if b := newsession.ISExpiration(); b {
		fmt.Println("已过期")
		return
	}
	t.Errorf("没过期就是出错了")

}

func testWriteDB(t *testing.T){
	err := newsession.WriteDB()
	if err != nil {
		t.Error(err)
	}
}

func testReadDB(t *testing.T){
	err := newsession.ReadDB()
	newsession.SetOverdueTime(100000)
	if err != nil {
		t.Error(err)
	}
	newsession.GetSyncMap().Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}

