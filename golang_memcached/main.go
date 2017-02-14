package main

import (
    "fmt"
    "github.com/bradfitz/gomemcache/memcache"
    "github.com/golang/glog"
    "os"
)

func main() {

    glog.Info("Starting transaction...")

    if len(os.Args) < 3 {
        glog.Exitln("Usage: main key value")
    }

    key := os.Args[1]
    value := os.Args[2]

    mc := memcache.New("localhost:11211")
    mc.Set(&memcache.Item{Key: key, Value: []byte(value)})

    it, err := mc.Get(key)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(it)
        fmt.Println(string(it.Value[:]))
    }
}
