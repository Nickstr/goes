package main

import (
    "github.com/nu7hatch/gouuid"
)

func GenerateUuid() *uuid.UUID {
    id, _ := uuid.NewV4()
    return id
}

func GenerateStringUuid() string {
    return GenerateUuid().String()
}
