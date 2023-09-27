# Code-generated Aiven Go Client

## Known limitations

- doesn't support query params
- `string` type is never a pointer, as it is not expected to send `""`
- authorization is by token only
- doesn't support custom certificates

## Configuration 

## via env vars

| Name               | Type     | Description            |
|--------------------|:---------|------------------------|
| `AIVEN_TOKEN`      | `string` | API auth token         |
| `AIVEN_WEB_URL`    | `string` | API Server location    |
| `AIVEN_USER_AGENT` | `string` | Client user agent      |
| `AIVEN_DEBUG`      | `bool`   | Enables Stderr logging |


## via constructor options

```go
client, err := aiven.NewClient(func(d *aiven.Doer){
	d.Debug = true
	d.UserAgent = "Smith"
})
```
