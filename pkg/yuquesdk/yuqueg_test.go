package yuquesdk_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
)

var (
	yu  *yuquesdk.Service
	rst string
	l   = yuquesdk.L
)

func setup() {
	yu = yuquesdk.NewService("xxx")
	yuquesdk.SetDebugLevel()
}

func shutdown() {
	l.Info(rst)
}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	if code == 0 {
		shutdown()
	}
	os.Exit(code)
}

func TestUserGet(t *testing.T) {
	d, err := yu.User.Get("")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	jsonString, _ := json.Marshal(d)
	rst = string(jsonString)
}
func TestGroupList(t *testing.T) {
	d, err := yu.Group.List("u22579")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	jsonString, _ := json.Marshal(d)
	rst = string(jsonString)
}
func TestDocList(t *testing.T) {
	d, err := yu.Doc.List("u22579/xcd0mr")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	jsonString, _ := json.Marshal(d)
	rst = string(jsonString)
}
