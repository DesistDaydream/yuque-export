package yuquesdk_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
)

var (
	yu  *yuquesdk.Service
	rst string
	l   = core.L
)

func setup() {
	yu = yuquesdk.NewService("4agpJJf7G2Xo0rIKCVj3n6GYPSGvQRKPQ4XHZK5Z")
	core.SetDebugLevel()
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
	d, err := yu.Group.List("desistdaydream")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	jsonString, _ := json.Marshal(d)
	rst = string(jsonString)
}
func TestDocList(t *testing.T) {
	d, err := yu.Doc.List("desistdaydream/entertainment")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	jsonString, _ := json.Marshal(d)
	rst = string(jsonString)
}
