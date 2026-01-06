# Go UUID

[![Pipeline](https://github.com/DaanV2/go-uuid/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-uuid/actions/workflows/pipeline.yaml)

A library that provides a way to handle, and generate UUIDs. Complies with RFC 4122 and RFC 9562 standards. Based on the code I wrote for [.dotnet - DaanV2.UUID.Net](https://github.com/DaanV2/DaanV2.UUID.Net)


```go
import "github.com/DaanV2/go-uuid"

u := uuid.New()
u := uuid.NewString()

//Specific version
u := uuid.V4.New()
u := uuid.V4.NewString()

// Hash-based UUIDs
u := uuid.V5.New([]byte("Some data"))
u := uuid.V3.New([]byte("Some data"))

// Time-based UUIDs (RFC 9562)
u, err := uuid.V6.New() // Reordered timestamp UUID (better for databases)
u := uuid.V7.New()      // Unix timestamp-based UUID (time-ordered)

// Custom/vendor-specific UUID (RFC 9562)
u := uuid.V8.New()                    // Random
u := uuid.V8.From([]byte("custom"))   // From custom data

// Batch generation
u := uuid.V4.NewBatch(10)

```

## Supported UUID Versions

- **V1**: Timestamp and MAC address-based UUID (RFC 4122)
- **V3**: MD5 hash-based UUID (RFC 4122)
- **V4**: Random UUID (RFC 4122)
- **V5**: SHA-1 hash-based UUID (RFC 4122)
- **V6**: Reordered timestamp UUID - better database indexing than V1 (RFC 9562)
- **V7**: Unix timestamp-based UUID with random data - time-ordered (RFC 9562)
- **V8**: Custom/vendor-specific UUID format (RFC 9562)

## Install

```bash
go get github.com/DaanV2/go-uuid
```