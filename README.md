# Song Contest Rater Protos

This repository contains the Protocol Buffer definitions for the Song Contest Rater application. It uses [Buf](https://buf.build/) to manage the Protocol Buffer definitions and generate code for both Go and TypeScript.

## Prerequisites

- [Buf CLI](https://buf.build/docs/installation)

## Development

### Generate Code

To generate code for both Go and TypeScript, run:

```sh
make generate
```

### Build Code

To build the Protocol Buffer files, run:

```sh
make build
```

This will generate code in the `gen` directory.

### Lint Protocol Buffer Files

To lint the Protocol Buffer files, run:

```sh
make lint
```

### Format Protocol Buffer Files

To format the Protocol Buffer files, run:

```sh
make format
```

### Check for Breaking Changes

To check for breaking changes against the registry, run:

```sh
make breaking
```

## Publishing

To publish a new version to the Buf Schema Registry, run:

```sh
make push
```

## License

MIT
