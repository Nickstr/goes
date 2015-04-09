package main
import "fmt"

type EventStoreStream struct {
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Id string `json:"id"`
    Updated string `json:"updated"`
    Rights string `json:"rights"`
    Link string `json:"selfUrl"`
    Author Author `json:"author"`
    EntryList []Entry `json:"entries"`
}

type Link struct {
    Href string `json:"href,attr"`
}

type Author struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

type Entry struct {
    Title string `json:"title"`
    Summary string `json:"summary"`
    Content string `json:"content"`
    Id string `json:"id"`
    Updated string `json:"updated"`
    Link Link `json:"link"`
    Author Author `json:"author"`
    EventId string `json:"eventId"`
    EventType string `json:"eventType"`
    EventNumber int `json:"eventNumber"`
    Data string `json:"data"`
}

func (s *EventStoreStream) GetEvents() []Entry {
    var events []Entry
    for _, event := range s.EntryList {
        fmt.Println(event)
        events = append([]Entry{event}, events...)
    }
    return events
}
