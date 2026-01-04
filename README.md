# gobotirc-lastfm

Get last track played from `https://www.last.fm/api/show/user.getRecentTracks` To be used with https://github.com/StalkR/goircbot

## usage

```golang
package main

import (
    "fmt"
    "log"
    "github.com/4cm3/gobotirc-lastfm"
)

func main() {
    r, err := lastfm.GetRecentTrack("APIKEY", "USERNAME" )
    if err != nil {
        log.Fatalf("error: %v\n", err)
    }
    fmt.Println(r.String())
}

## documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/4cm3/gobotirc-lastfm.svg)](https://pkg.go.dev/github.com/4cm3/gobotirc-lastfm)