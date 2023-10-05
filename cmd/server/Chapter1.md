## test
```
curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzIK"}}'

curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzMK"}}'
```

動作確認例
```
# ~/Practice/Go/proglog on git:main 
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
{"offset":0}

# ~/Practice/Go/proglog on git:main
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzIK"}}'
{"offset":1}

# ~/Practice/Go/proglog on git:main
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzMK"}}'
{"offset":2}
```

呼べば呼ぶだけカウントされていくみたい
```
同じvalueで何度か実行した例
curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
{"offset":0}

curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
{"offset":1}

curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
{"offset":2}

```


GETもできる
```
curl -X GET localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
{"record":{"value":"TGV0J3MgR28gIzEK","offset":0}}
```