package args

import (
	"encoding/json"
	"testing"

	"src/tools.app/apps/webmvc/model"
)

func TestAuthArg01(t *testing.T) {
	arg := &AuthArg{
		PageArg: PageArg{
			Pagefrom: 1,
			Pagesize: 1,
			Asc:      "date",
			Desc:     "time",
		},
		User: model.User{
			ID:       111,
			NickName: "test_user01",
		},
		Code: "200",
	}
	b, err := json.Marshal(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("json string:", string(b))
}

func TestAuthArg02(t *testing.T) {
	var arg AuthArg
	arg.Pagefrom = 1
	arg.Pagesize = 2
	arg.Asc = "year"
	arg.Desc = "month"
	arg.ID = 112
	arg.NickName = "test_user02"
	arg.Code = "201"

	b, err := json.Marshal(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("json string:", string(b))
}

func TestAuthArg03(t *testing.T) {
	arg := new(AuthArg)
	arg.Pagefrom = 1
	arg.Pagesize = 2
	arg.Asc = "year"
	arg.Desc = "month"
	arg.ID = 113
	arg.NickName = "test_user03"
	arg.Code = "203"

	b, err := json.Marshal(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("json string:", string(b))
}
