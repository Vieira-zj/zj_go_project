#!/bin/bash
set -eu

if [[ $1 == "index" ]]; then
    curl -v "http://localhost:17891/index"
fi

if [[ $1 == "reg" ]]; then
    curl -v "http://localhost:17891/d/test?debug=true"
fi

if [[ $1 == "login" ]]; then
    curl -v -XPOST "http://localhost:17891/user/login" -H "Content-Type:application/json" \
	  -d '{"pagefrom":1,"pagesize":1,"asc":"field1","desc":"field2","id":111,"nickName":"test_user01","role":0,"code":"200"}'
fi

# echo "mvc web app test done."