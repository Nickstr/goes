package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "strings"
    "reflect"
    "github.com/nu7hatch/gouuid"
    "bytes")

type DomainEvent struct {
    Id *uuid.UUID   `json:"id"`
    Type string     `json:"type"`
    Event []byte    `json:"event"`
}

func RaiseEvent(stream string, event interface{}) {
    eventType := strings.Split(reflect.TypeOf(event).String(), ".")
    de := DomainEvent{
        Id: GenerateUuid(),
        Type: eventType[len(eventType) - 1],
    }
    e, _ := json.Marshal(event)
    de.Event = e
    write(stream, de)
}

func write(stream string, e DomainEvent) {
    url := "http://10.10.10.10:2113/streams/" + stream
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(e.Event))
    req.Header.Set("ES-EventType", e.Type)
    req.Header.Set("ES-EventId", GenerateUuid().String())
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
}

func Read(s string) *EventStoreStream {
    url := "http://10.10.10.10:2113/streams/" + s + "?embed=body"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", "application/json")
    req.SetBasicAuth("admin", "changeit")

    stream := &EventStoreStream{}
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    jsonContent, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(jsonContent, stream)

    for _, event := range stream.EntryList {
        fmt.Println(event.EventType)
    }

    defer resp.Body.Close()
    fmt.Println("response Status:", resp.Status)
    return stream
}
