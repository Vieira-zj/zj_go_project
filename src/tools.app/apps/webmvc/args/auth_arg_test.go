package args

import (
	"encoding/json"
	"testing"

	"tools.app/apps/webmvc/model"
)

func TestAuthArg(t *testing.T) {
	arg := &AuthArg{
		User: model.User{
			ID:       111,
			NickName: "test_user01",
		},
		PageArg: PageArg{
			Pagefrom: 1,
			Pagesize: 1,
			Asc:      "field1",
			Desc:     "field2",
		},
		Code: "200",
	}
	b, err := json.Marshal(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("json string:", string(b))
}
