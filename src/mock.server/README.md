# Mock Server APIs

#### /demo/:id

`curl -v "http://127.0.0.1:17891/demo/01?userid=xxx&username=xxx"`

`curl -v "http://127.0.0.1:17891/demo/02?userid=xxx&username=xxx&key=val1&key=val2"`

`curl -v -X POST "http://127.0.0.1:17891/demo/03" -H "Content-Type:application/json" --data-binary @test.json`

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

`curl -v -X POST "http://127.0.0.1:17891/demo/04" --cookie "user=user_001;pwd=test_com" --form key1=val1 --form key2=val2`

`curl -v -X POST "http://127.0.0.1:17891/demo/04" --cookie "user=user_001;pwd=test_com" --data-binary @test.out`

```
key1=val1;key2=val2
```

`curl -v "http://127.0.0.1:17891/demo/05"`
