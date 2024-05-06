# go-client-codegen

`go-client-codegen` is an automatically code generated Aiven Go Client from the Aiven API specification.

_Warning:_ this client is under heavy development.

## Setup

```bash
go get github.com/aiven/go-client-codegen
```

### Configuration and Usage

#### Via Environment Variables

| Name               | Type     | Description                    |
| ------------------ | :------- | ------------------------------ |
| `AIVEN_TOKEN`      | `string` | Aiven API Authentication Token |
| `AIVEN_WEB_URL`    | `string` | Aiven API URL                  |
| `AIVEN_USER_AGENT` | `string` | User Agent                     |
| `AIVEN_DEBUG`      | `bool`   | Debug Output Flag (stderr)     |

#### Via Constructor Options

```go
import "github.com/aiven/go-client-codegen"

client, err := aiven.NewClient(DebugOpt(true), UserAgentOpt("foo"))
if err != nil {
	return err
}

services, err := client.ServiceList(ctx, "bar-project")
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute to the development of go-client-codegen.

## License

go-client-codegen is licensed under the Apache license, version 2.0. Full license text is available in the
[LICENSE](LICENSE) file.

Please note that the project explicitly does not require a CLA (Contributor License Agreement) from its contributors.

## Contact

Bug reports and patches are very welcome, please post them as GitHub issues and pull requests at
https://github.com/aiven/go-client-codegen. To report any possible vulnerabilities or other serious issues please see
our [security](SECURITY.md) policy.
