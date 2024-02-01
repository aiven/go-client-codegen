# Code-generated Aiven Go Client

## Known limitations

- doesn't support query params
- doesn't support custom certificates
- `string` type is never a pointer, as it is not expected to send `""`
- authorization is by token only
- slices never omitted (always sent), as it makes development much easier

## Installation

``` shell
go get github.com/aiven/aiven-go-client-v3
```

## Configuration and usage

### via env vars

| Name               | Type     | Description            |
|--------------------|:---------|------------------------|
| `AIVEN_TOKEN`      | `string` | API auth token         |
| `AIVEN_WEB_URL`    | `string` | API Server location    |
| `AIVEN_USER_AGENT` | `string` | Client user agent      |
| `AIVEN_DEBUG`      | `bool`   | Enables Stderr logging |


### via constructor options

```go
import "github.com/aiven/aiven-go-client-v3"

client, err := aiven.NewClient(DebugOpt(true), UserAgentOpt("smith"))
if err != nil {
	return err
}

services, err := client.ServiceList(ctx, "my-project")
```

## Design decisions

### Al-in-one interface

The `aiven.Client` exposes all Aiven methods by OperationID instead of providing with scoped/grouped handlers. 
This approach has several benefits:

1. The OperationID is immutable, means the interface should not dramatically change if spec is changed
2. Easier mocking/testing, as it is possible to create a subset of client methods:

```go
type sweeperClient interface {
    ServiceDelete(ctx context.Context, project string, serviceName string) error
    ServiceUserDelete(ctx context.Context, project string, serviceName string, serviceUsername string) error
    VpcDelete(ctx context.Context, project string, projectVpcId string) (*vpc.VpcDeleteOut, error)
}

func sweeper(client serviceShredderClient) error {
	... // sweep
}
```

### Pointers for reference types

The Aiven API distinguishes between cases when the field is missing or has a zero-value.  
In some cases, we must avoid sending empty arrays or objects even if that works as expected.
For instance, sending `tech_emails` triggers creation of an additional event log entry in Console.
As a universal solution, the client takes `nil` as "missing".

### `[]Foo`, not `[]*Foo`

The generator doesn't create pointers for array elements.
Because technically that means it might contain `nil` values.
Therefore `nil` checks _must_ be performed.

### Response objects

Request and response objects are separated and do not share code, except enums.
Even though if they look similar:

```go
type UserIn struct {
	Name string `json:"name"`
}

type UserOut struct {
	Name string `json:"name"`
}
```

That's made on purpose, so if a request or response object has been changed, hence `UserIn != UserOut` 
it won't generate a new struct.
Which in turn might regenerate other objects because of the name collision.
Keeping things separate makes generated code more durable.
