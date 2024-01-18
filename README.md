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

## via env vars

| Name               | Type     | Description            |
|--------------------|:---------|------------------------|
| `AIVEN_TOKEN`      | `string` | API auth token         |
| `AIVEN_WEB_URL`    | `string` | API Server location    |
| `AIVEN_USER_AGENT` | `string` | Client user agent      |
| `AIVEN_DEBUG`      | `bool`   | Enables Stderr logging |


## via constructor options

```go
import "github.com/aiven/aiven-go-client-v3"

client, err := aiven.NewClient(DebugOpt(true), UserAgentOpt("smith"))
if err != nil {
	return err
}

services, err := client.ServiceList(ctx, "my-project")
```

## Design

The `aiven.Client` exposes all Aiven methods by operation id.
For instance, [`ServiceList`](https://api.aiven.io/doc/#tag/Service/operation/ServiceList) is literally just:

```go
client.ServiceList(ctx, "my-project")
```

This approach allows the creation of custom methods subsets that might be helpful with testing and mocking.

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
