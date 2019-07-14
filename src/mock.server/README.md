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

Test, mock return bytes body with wait.
- `curl -v "http://127.0.0.1:17891/mocktest/one/1?size=128&wait=1"`

Test, mock return file content with wait.
- `curl -v "http://127.0.0.1:17891/mocktest/one/2?file=test.out&wait=1"`

Test, mock return custom error code.
- `curl -v "http://127.0.0.1:17891/mocktest/one/3?code=403"`

Test, mock httpdns server which returns json string.
- `curl -v "http://127.0.0.1:17891/mocktest/one/4?wait=1"`

Test, mock gzip and chunk http response.
- `curl -v "http://127.0.0.1:17891/mocktest/one/5"`

Test, mock http response diff mimetype.
- `curl -v "http://127.0.0.1:17891/mocktest/one/6?type=txt&errlen=false"`

### /mocktest/two/:id

Test, mock 403 Forbidden, or return file content.
- `curl -v "http://127.0.0.1:17891/mocktest/two/1?iserr=false"`

Test, mock return bytes data by flush.
- `curl -v "http://127.0.0.1:17891/mocktest/two/2"`

Test, mock return bytes by range with wait.

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

Test, mock kb data with wait between each read.
- `curl -v "http://127.0.0.1:17891/mocktest/two/4?kb=3&wait=2"`

Test, mock server side disconnect.
- `curl -v "http://127.0.0.1:17891/mocktest/two/5?wait=1"`

### /mockqiniu/:id

Mock mirror file server handler.
- `curl -v "http://127.0.0.1:17891/mockqiniu/1?wait=1"`

CDN refresh request handler.
- `curl -v "http://127.0.0.1:17891/mockqiniu/2"`

Mock return diff file content by arg "start".
- `curl -v "http://127.0.0.1:17891/mockqiniu/3?start=100"`
