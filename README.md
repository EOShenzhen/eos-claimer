# eos-claimer


### config main.go
```
var bp_name string = "your block producer name"
var claim_key string = "your claimer key, active key is not advised"
var claim_key_permission_name string = "claimer"
var end_points = []string{}  # some active api endpoints
```

### build
``` 
# download dependence
$ go get github.com/eoscanada/eos-go

# build binary for various platform 
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eos-claimer-linux
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o eos-claimer-windows
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o eos-claimer-darwin

```

### run
``` 
./eos-claimer-darwin

```

### install
``` 
cp eos-claimer-linux /usr/bin/eos-worker

```


### crontab
```
crontab -e
*/1 * * * * /usr/bin/eos-worker  1>>/tmp/tmp.log

# 注意：在上一行的最后不可以加2>>&1 会导致整行命令不执行
```
