package main

import (
    "fmt"
)

func main() {
    i := CreateInvoice("original title")
    i.AddLineItem("original task", "test title", "50,00", 5)
    i.AddLineItem("original task", "test title", "50,00", 5)
    i.AddLineItem("original task", "test title", "50,00", 5)

    for _, item := range i.LineItems {
        i.UpdateLineItemTask(item.Id, "updated task")
    }

    i.UpdateInvoiceTitle("updated title")
    buildFromStream(i.Id)
}

func buildFromStream(id string) {
    ie := BuildInvoiceFromStream(id)

    fmt.Println(ie.Title)
    for _, item := range ie.LineItems {
        fmt.Println(item.Task)
    }
}
