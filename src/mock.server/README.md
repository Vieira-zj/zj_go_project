# Mock Server APIs

### /

1. Default page:

`curl -v "http://127.0.0.1:17891/`

### /demo/:id

1. Demo, parse Get request:

`curl -v "http://127.0.0.1:17891/demo/1?userid=xxx&username=xxx"`

2. Demo, parse Get request:

`curl -v "http://127.0.0.1:17891/demo/2?userid=xxx&username=xxx&key=val1&key=val2"`

3. Demo, parse Post json body:

`curl -v -X POST "http://127.0.0.1:17891/demo/3" -H "Content-Type:application/json" --data-binary @test.json`

```json
{
    "server_list": [
        {
            "server_name": "svr_a_002",
            "server_ip": "127.0.1.2"
        },
        {
            "server_name": "svr_a_013",
            "server_ip": "127.0.1.13"
        }
    ],
    "server_group_id": "svr_grp_001"
}
```

4. Demo, parse Post form with cookie:

`curl -v -X POST "http://127.0.0.1:17891/demo/4" --cookie "user=user_001;pwd=test_com" --form key1=val1 --form key2=val2`

`curl -v -X POST "http://127.0.0.1:17891/demo/4" --cookie "user=user_001;pwd=test_com" --data-binary @test.out`

```text
key1=val1;key2=val2
```

5. Demo, show access count which is stored in redis:

`curl -v "http://127.0.0.1:17891/demo/5"`

### /mocktest/one/:id

1. Mock test, returns bytes body with wait:

`curl -v "http://127.0.0.1:17891/mocktest/one/1?size=128&wait=1"`

2. Mock test, returns file content with wait:

`curl -v "http://127.0.0.1:17891/mocktest/one/2?file=test_log.txt&wait=1"`

3. Mock test, returns custom error code, like 403, 502:

`curl -v "http://127.0.0.1:17891/mocktest/one/3?code=403"`

4. Mock test, returns httpdns json string:

`curl -v "http://127.0.0.1:17891/mocktest/one/4?wait=1"`

5. Mock test, returns gzip and chunked http response:

`curl -v "http://127.0.0.1:17891/mocktest/one/5"`

6. Mock test, returns http response with diff mimetype:

`curl -v "http://127.0.0.1:17891/mocktest/one/6?type=txt&errlen=false"`

### /mocktest/two/:id

1. Mock test, returns 403 Forbidden, or file content:

`curl -v "http://127.0.0.1:17891/mocktest/two/1?iserr=false"`

2. Mock test, returns chunked of bytes by flush:

`curl -v "http://127.0.0.1:17891/mocktest/two/2"`

3. Mock test, returns bytes by range with wait:

```sh
#!/bin/bash
start=0
for i in {1..10};do
    end=$(($start+1023))
    curl -v "http://127.0.0.1:17891/mocktest/two/3" -H "Range:bytes=${start}-${end}"
    start=$(($end+1))
done
echo "done"
```

4. Mock test, returns kb data with wait in each read:

`curl -v "http://127.0.0.1:17891/mocktest/two/4?wait=100&kb=3"`

5. Mock test, server side close connection:

`curl -v "http://127.0.0.1:17891/mocktest/two/5?wait=1"`

### /mockqiniu/:id

1. Mock mirror file server handler:

`curl -v "http://127.0.0.1:17891/mockqiniu/1?wait=2"`

2. Mock CDN refresh request handler:

`curl -v "http://127.0.0.1:17891/mockqiniu/2"`

3. Mock return diff file content by arg "start":

`curl -v "http://127.0.0.1:17891/mockqiniu/3?start=100"`