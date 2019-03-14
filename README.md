# Bitly API

## Major Design Decisions

I wanted to make the api versioned by adding a sub route of /v1 so it would help with future iterations. I did opt to pass the access token through an authorization header so there was no secret storage in the api. For the query parameters, I decided to follow adopt the “unit” and “units” parameters from the real bitly api  to handle the time frames. Average parameter was also added to endpoint to allow the client to decide if they want the average or the default response of the total amount of clicks. I did not make the group value dynamic since the goal of this assessment was to get the average clicks of a user’s default group by countries (/api/v1/groups/default/countries). I kept “groups” and “countries” plural to prevent having to handle singular versus plural endpoints like (person/people, goose/geese).
___

## Getting Started

**Go version**

- go version go1.12

**Setup Project**

Setup this project in your go path under `$GOPATH/src/github.com/kwtucker`

```shell
 $ mkdir -p $GOPATH/src/github.com/kwtucker
 $ mv ./bitlyapi $GOPATH/src/github.com/kwtucker/bitlyapi
 $ cd $GOPATH/src/github.com/kwtucker/bitlyapi
```

**Dependencies**

```shell
$ go get -t ./...
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

**Header**

`Authorization:"Bearer {BITLY_ACCESS_TOKEN}"`

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