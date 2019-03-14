# Bitly API

## Major design decisions
___

## Getting Started

**Go version**

- go version go1.12

**Move Project**

Move this project into your go path under `$GOPATH/src/github.com/kwtucker`

```shell
 $ mv -r./bitlyapi $GOPATH/src/github.com/kwtucker/bitlyapi
 $ mkdir -p $GOPATH/src/github.com/kwtucker && mv ./bitlyapi $GOPATH/src/github.com/kwtucker/bitlyapi
 $ cd $GOPATH/src/github.com/kwtucker/bitlyapi
```

**Dependencies**

```shell
$ go get ./...
```

**Test**

```shell
$ go test -cover -v  ./...
```

**Build**

```shell
$ go build
```

**Run** (Default port :8080)

```shell
$ ./bitlyapi
```

- `-port {PORT}`

___
## Documentation

### Number of clicks for Bitlinks in a user default group.
**[GET]**

```shell
 /api/v1/groups/default/countries
```

- `unit` [ String ]
  - Default: "day"
  - Enum:"minute" "hour" "day" "week" "month"
- `units` [ Integer ]
  - Default: "0"
  - An integer representing the time units to query data for. pass 0 to return all units of time.
- `average` [ Boolean ]
  - Default: false
  - Average number of clicks, per country, for the Bitlinks in a user's default group.

---


## Sample Request

### **"Average"** Number Clicks in 30 days for Bitlinks in user default group.
**[GET]**
`
 /api/v1/groups/default/countries?units=30&unit=day&average=true
`
#### Description
Get the average number of clicks, per country, within the last 30 days, for the Bitlinks in a user's default group.
#### Curl
```shell
$ curl -i 'localhost:8080/api/v1/groups/default/countries?units=30&unit=day&average=true' -H "Authorization: Bearer {BITLY_ACCESS_TOKEN}"
```
#### HTTPIE
```shell
$ http localhost:8080/api/v1/groups/default/countries units==30 unit==day average==true Authorization:"Bearer {BITLY_ACCESS_TOKEN}"
```