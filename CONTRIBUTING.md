# Contributing

## Welcome

Contributions are very welcome on go-api-schemas. When contributing please keep this in mind:

- Open an issue to discuss new bigger features.
- Write code consistent with the project style and make sure the tests are passing.
- Stay in touch with us if we have follow up questions or requests for further changes.

## Architectural Choices

### Unified Interface Concept

The `aiven.Client` offers a singular interface to access all Aiven functionalities through `OperationID`, rather than
organizing them into categorized handlers. This design has several advantages:

1. The permanence of `OperationID` ensures stability in the interface, even when specifications undergo modifications.
2. It simplifies the process of mocking/testing by allowing the creation of client method subsets:

```go
type sweeperClient interface {
    ServiceDelete(ctx context.Context, project string, serviceName string) error
    ServiceUserDelete(ctx context.Context, project string, serviceName string, serviceUsername string) error
    VpcDelete(ctx context.Context, project string, projectVpcId string) (*vpc.VpcDeleteOut, error)
}

func sweeper(client serviceShredderClient) error {
	... // cleaning process
}
```

### Use of Pointers for Nullable Types

To differentiate between absent fields and fields with zero-values, the Aiven API uses a distinct approach. In certain
scenarios, it's crucial to avoid sending empty structures or arrays, as they may inadvertently trigger additional
actions, like creating an event log entry when `tech_emails` is sent. The client, therefore, treats `nil` as indicative
of a field's absence.

### Array Elements Without Pointers

The code generator is designed not to use pointers for elements within arrays, to avoid the inclusion of `nil` values
within these arrays. Consequently, it necessitates explicit `nil` checks.

### Distinct Request and Response Structures

Despite similarities in appearance, request and response structures are intentionally kept distinct, with no shared
codebase except for enums. This is to ensure that changes in either request or response objects don't lead to the
creation of new structures due to naming conflicts, thus enhancing the robustness of the generated code by maintaining
separation:

```go
type UserIn struct {
	Name string `json:"name"`
}

type UserOut struct {
	Name string `json:"name"`
}
```

## Development

### Local Environment

Place the OpenAPI specification file in the root directory of the project as `openapi.json`, and run the following
command to generate the client code:

```bash
task generate
```

### Tests

```bash
task test
```

### Static checking and Linting

We use [Trunk.io](https://trunk.io/) for static checking and linting. Install it locally and you'll be good to go.

## Opening a PR

- Commit messages should describe the changes, not the filenames. Win our admiration by following the
  [excellent advice from Chris Beams](https://chris.beams.io/posts/git-commit/) when composing commit messages.
- Choose a meaningful title for your pull request.
- The pull request description should focus on what changed and why.
- Check that the tests pass (and add test coverage for your changes if appropriate).

### Commit Messages

This project adheres to the [Conventional Commits](https://conventionalcommits.org/en/v1.0.0/) specification.
Please, make sure that your commit messages follow that specification.
