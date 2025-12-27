# gobotirc-lastfm

get last track played from `https://www.last.fm/api/show/user.getRecentTracks`

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
