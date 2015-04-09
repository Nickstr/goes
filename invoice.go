package main
import "encoding/json"

type Invoice struct {
    Id string
    Title string
    LineItems []*LineItem
}

type LineItem struct {
    Id string
    Task string
    Description string
    Rate string
    Hours float64
}

type InvoiceCreatedEvent struct {
    Id string
    Title string
}

type InvoiceTitleUpdatedEvent struct {
    Id string `json:"id"`
    Title string
}

type LineItemAddedEvent struct {
    Id string
    Task string
    Description string
    Rate string
    Hours float64
}

type LineItemTaskUpdatedEvent struct {
    Id string
    Task string
}

func BuildInvoiceFromStream(id string) *Invoice {
    i := &Invoice{}
    stream := Read(id)
    for _, event := range stream.GetEvents() {
        switch event.EventType {
            case "InvoiceCreatedEvent":
                e := &InvoiceCreatedEvent{}
                json.Unmarshal([]byte(event.Data), e)
                i.ApplyInvoiceCreatedEvent(e)
            case "LineItemAddedEvent":
                e := &LineItemAddedEvent{}
                json.Unmarshal([]byte(event.Data), e)
                i.ApplyLineItemAddedEvent(e)
            case "LineItemTaskUpdatedEvent":
                e := &LineItemTaskUpdatedEvent{}
                json.Unmarshal([]byte(event.Data), e)
                i.ApplyUpdateLineItemTaskEvent(e)
            case "InvoiceTitleUpdatedEvent":
                e := &InvoiceTitleUpdatedEvent{}
                json.Unmarshal([]byte(event.Data), e)
                i.ApplyInvoiceTitleUpdatedEvent(e)
            }
    }

    return i
}

func CreateInvoice(title string) *Invoice {
    i := &Invoice{}
    e := &InvoiceCreatedEvent{
        Id: GenerateStringUuid(),
        Title: title,
    }
    RaiseEvent(e.Id, e)
    i.ApplyInvoiceCreatedEvent(e)
    return i
}

func (i *Invoice) UpdateInvoiceTitle(title string) {
    e := &InvoiceTitleUpdatedEvent{
        Id: GenerateStringUuid(),
        Title: title,
    }
    RaiseEvent(i.Id, e)
    i.ApplyInvoiceTitleUpdatedEvent(e)
}

func (i *Invoice) AddLineItem(task string, description string, rate string, hours float64) {
    e := &LineItemAddedEvent{
        Id: GenerateStringUuid(),
        Task: task,
        Description: description,
        Rate: rate,
        Hours: hours,
    }
    RaiseEvent(i.Id, e)
    i.ApplyLineItemAddedEvent(e)
}

func (i *Invoice) ApplyInvoiceCreatedEvent(e *InvoiceCreatedEvent) {
    i.Id = e.Id
    i.Title = e.Title
}

func (i *Invoice) ApplyInvoiceTitleUpdatedEvent(e *InvoiceTitleUpdatedEvent) {
    i.Title = e.Title
}

func (i *Invoice) ApplyLineItemAddedEvent(e *LineItemAddedEvent) {
    l := &LineItem{
        Id: e.Id,
        Task: e.Task,
        Description: e.Description,
        Rate: e.Rate,
        Hours: e.Hours,
    }
    i.LineItems = append(i.LineItems, l)
}

func (i *Invoice) UpdateLineItemTask(id string, task string) {
    e := &LineItemTaskUpdatedEvent{
        Id: id,
        Task: task,
    }
    i.ApplyUpdateLineItemTaskEvent(e)
    RaiseEvent(i.Id, e)
}

func (i *Invoice) ApplyUpdateLineItemTaskEvent(e *LineItemTaskUpdatedEvent) {
    l := i.findLineItem(e.Id);
    l.Task = e.Task
}

func (i *Invoice) findLineItem(id string) *LineItem {
    for _, item := range i.LineItems {
        if item.Id == id {
            return item
        }
    }
    return nil
}
