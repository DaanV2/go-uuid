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

// Parsing (canonical, braced, urn:uuid: and dash-less forms are all accepted)
u, err := uuid.Parse("01234567-89ab-cdef-0123-456789abcdef")
u := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")

// Formatting
u.String()    // 01234567-89ab-cdef-0123-456789abcdef
u.StringHex() // 0123456789abcdef0123456789abcdef
u.URN()       // urn:uuid:01234567-89ab-cdef-0123-456789abcdef

// Comparing and sorting
a.Equal(b)      // bool
a.Compare(b)    // -1, 0 or 1
slices.SortFunc(ids, uuid.UUID.Compare)

```

## Encoding & interoperability

`UUID` implements the standard library interfaces so it works out of the box
with JSON, XML, `database/sql`, and any package that relies on them:

- `encoding.TextMarshaler` / `encoding.TextUnmarshaler`
- `encoding.BinaryMarshaler` / `encoding.BinaryUnmarshaler`
- `json.Marshaler` / `json.Unmarshaler`
- `driver.Valuer` / `sql.Scanner`

```go
type User struct {
    ID uuid.UUID `json:"id"` // encodes/decodes as a canonical string
}

// Works directly as a SQL argument and scan target.
row := db.QueryRow("SELECT id FROM users WHERE id = ?", user.ID)
_ = row.Scan(&user.ID)
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