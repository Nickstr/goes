package main

import "testing"

func TestUpdateInvoiceTitle(t *testing.T) {
    i := createInvoiceWithLineItems("foo")
    i.UpdateInvoiceTitle("bar")
    if i.Title != "bar"  {
        t.Error("Title should be Bar")
    }
}

func TestAddLineItem(t *testing.T) {
    i := createInvoiceWithLineItems("foo")

    if i.LineItems[0].Task != "foo" {
        t.Error("Task should be foo")
    }
    if i.LineItems[1].Task != "bar" {
        t.Error("Task should be bar")
    }
}

func TestUpdateLineItemTask(t *testing.T) {
    i := createInvoiceWithLineItems("foo")

    i.UpdateLineItemTask(i.LineItems[0].Id, "bar")
    if i.LineItems[0].Task != "bar" {
        t.Error("Task should be bar")
    }

    i.ApplyUpdateLineItemTaskEvent(&LineItemTaskUpdatedEvent{
        Id: i.LineItems[1].Id,
        Task: "foo",
    })
    if i.LineItems[1].Task != "foo" {
        t.Error("Task should be foo")
    }
}

func createInvoiceWithLineItems(title string) *Invoice {
    i := CreateInvoice(title)
    i.AddLineItem("foo", "test title", "50,00", 5)
    i.AddLineItem("bar", "test title", "50,00", 5)
    return i
}
