# Go UUID

[![Pipeline](https://github.com/DaanV2/go-uuid/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-uuid/actions/workflows/pipeline.yaml)

A library that provides a way to handle, and generate UUIDs. Complies with the RFC 4122 standard. based the code on what I wrote for [.dotnet - DaanV2.UUID.Net](https://github.com/DaanV2/DaanV2.UUID.Net)


```go
import "github.com/DaanV2/go-uuid"

u := uuid.New()
u := uuid.NewString()

//Specific version
u := uuid.V4.New()
u := uuid.V4.NewString()

u := uuid.V5.New([]byte("Some data"))
u := uuid.V3.New([]byte("Some data"))

// Batch generation
u := uuid.V4.NewBatch(10)

```

## Install

```bash
go get github.com/DaanV2/go-uuid
```