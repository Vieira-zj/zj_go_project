# Mock Server APIs

## Default

1. Ping:

```sh
curl -v "http://127.0.0.1:17891/ping"
```

## Demo

`demo/:id`:

1. Parse Get request:

```sh
curl -v "http://127.0.0.1:17891/demo/1?userid=001&username=foo"
```

2. Parse Get request:

```sh
curl -v "http://127.0.0.1:17891/demo/2?userid=001&username=bar&key=val1&key=val2"
```

3. Parse Post json body:

```sh
curl -v -X POST "http://127.0.0.1:17891/demo/3" -H "Content-Type:application/json" --data-binary @test.json
```

`test.json`:

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

4. Parse Post form with cookie:

```sh
curl -v -X POST "http://127.0.0.1:17891/demo/4" --cookie "user=user_001;pwd=test_com" --form key1=val1 --form key2=val2

curl -v -X POST "http://127.0.0.1:17891/demo/4" --cookie "user=user_001;pwd=test_com" --data-binary @data.txt
```

`data.txt`:

```text
key1=val1;key2=val2
```

5. Show access count which is stored in redis:

```sh
curl -v "http://127.0.0.1:17891/demo/5"
```

## Demo with templated json response

1. Parse Post request, and return templated json:

```sh
curl -v -X POST "http://127.0.0.1:17891/demo/2-1?userid=mmm&username=nnn&age=19&sex=male&key1=val1&key2=val2" \n
  -H "Content-Type:text/plain;charset=UTF-8" --data-binary @data.txt
```

2. Parse Post request with keywords (randint, randstr), and return templated json:

```sh
curl -v -X POST "http://127.0.0.1:17891/demo/2-1?userid=xxx&username=yyy&age=randint(37)&sex=randchoice(male,female)&key1=val1&key2=randstr(16)" \n
  -H "Content-Type:text/plain;charset=UTF-8" --data-binary @data.txt
```

request body `data.txt`:

```text
{
  "user_info": {
    "user_id": "{{.userid}}",
    "user_name": "{{.username}}",
    "user_age": {{.age}},
    "user_sex": "{{.sex}}",
    "meta": {
      "key1": "{{.key1}}",
      "key2": "{{.key2}}"
    }
  }
}
```

response json:

```json
{
  "user_info": {
    "user_id": "xxxxx",
    "user_name": "xxxxx",
    "user_age": 21,
    "user_sex": "male",
    "meta": {
      "key1": "val1",
      "key2": "rand_string"
    }
  }
}
```

## Mock Test

`/mocktest/one/:id`

1. Return bytes body with wait:

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/1?size=128&wait=1"
```

2. Return file content with wait:

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/2?file=test_log.txt&wait=1"
```

3. Return custom error code, like 403, 502:

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/3?code=403"
```

4. Return httpdns json string with wait (wait=sec/milli, sec by default):

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/4?wait=200&unit=milli"
```

5. Return gzip and chunked http response:

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/5"
```

6. Return http response with diff mimetype:

```sh
curl -v "http://127.0.0.1:17891/mocktest/one/6?type=txt&errlen=false"
```

## Mock Test 2

`/mocktest/two/:id`

1. Return 403 Forbidden, or file content:

```sh
curl -v "http://127.0.0.1:17891/mocktest/two/1?iserr=false"
```

2. Return chunked of bytes by flush:

```sh
curl -v "http://127.0.0.1:17891/mocktest/two/2"
```

3. Return bytes by range with wait:

`test.sh`:

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

4. Return kb data with wait in each read:

```sh
curl -v "http://127.0.0.1:17891/mocktest/two/4?wait=100&kb=3"
```

5. Mock server side close connection:

```sh
curl -v "http://127.0.0.1:17891/mocktest/two/5?wait=1"
```

## Mock Apis

1. Register a uri with params and template body (Post `/mock/register/:uri`):

```sh
# simple request
curl -v "http://127.0.0.1:17891/mock/register/mock-001" \
  -H "Content-Type:text/plain;charset=UTF-8" -d "hello world"

# request with keys
curl -v "http://127.0.0.1:17891/mock/register/mock-001?userid=xxx&username=yyy&age=randint(27)&sex=randchoice(male,female)&key1=val1&key2=randstr(12)" \
  -H "Content-Type:text/plain;charset=UTF-8" --data-binary @data.txt
```

request body `data.txt`:

```text
{
  "user_info": {
    "user_id": "{{.userid}}",
    "user_name": "{{.username}}",
    "user_age": {{.age}},
    "user_sex": "{{.sex}}",
    "meta": {
      "key1": "{{.key1}}",
      "key2": "{{.key2}}"
    }
  }
}
```

response json:

```json
{
  "meta": null,
  "data": {
    "status": 200,
    "message": "success"
  }
}
```

2. Access register uri, and get templated json body (Get `/mock/:uri`):

```sh
curl -v "http://127.0.0.1:17891/mock/mock-001"
```

response json:

```json
{
  "user_info": {
    "user_id": "xxx",
    "user_name": "yyy",
    "user_age": 21,
    "user_sex": "male",
    "meta": {
      "key1": "val1",
      "key2": "rand_string"
    }
  }
}
```

## Mock Qiniu Apis

`/mockqiniu/:id`

1. Mock mirror file server handler:

```sh
curl -v "http://127.0.0.1:17891/mockqiniu/1?wait=2"
```

2. Mock CDN refresh request handler:

```sh
curl -v "http://127.0.0.1:17891/mockqiniu/2"
```

3. Mock return diff file content by arg "start":

```sh
curl -v "http://127.0.0.1:17891/mockqiniu/3?start=100"
```

## Tools Apis

`/tools/:name`

1. Run shell commands:

```sh
curl -v -X POST "http://127.0.0.1:17891/tools/cmd" -H "Content-Type:application/json" --data-binary "@test.json"
```

request json:

```json
{
  "meta": "reserved keyword",
  "commands": [
    "hostname",
    "go version"
  ]
}
```

response json:

```json
{
  "meta": null,
  "data": {
    "status": 200,
    "message": "success",
    "results": "zjmbp\ngo version go1.11 darwin/amd64\n"
  }
}
```

2. Send an email:

```sh
curl -v -X POST "http://127.0.0.1:17891/tools/mail" -H "Content-Type:application/json" --data-binary "@test.json"
```

request json:

```json
{
  "meta": "reserved keyword",
  "receivers": [
    "zhengjin@4paradigm.com"
  ],
  "subject": "Go Mail Test",
  "body": "This is a go mail test from MockServer.",
  "attachments": [
    "./test_log.txt",
    "./logs"
  ],
  "archive": true
}
```

response json:

```json
{
  "meta": null,
  "data": {
    "status": 200,
    "message": "success"
  }
}
```

