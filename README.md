# Spacemesh API
[Protobuf](https://developers.google.com/protocol-buffers) implementation of the Spacemesh API. This is a work in progress. See the more complete [master API spec](https://docs.google.com/spreadsheets/d/1P89OVWdgJocPy0CGM43Ge7Sx_6dabCBEagaVQfOk9us/edit).

## Intended Usage Pattern

### Mesh data processing flow
1. Client starts a full node with flags set to turn syncing off and to open the GRPC APIs
1. Client registers on the streaming GRPC api methods that are of interest
1. Client calls `node.SyncStart()` to request that the node start syncing
1. Client processes streaming data it receives from the node
1. Client monitors node using `node.SyncStatusStream()` and `node.ErrorStream()` and handle node critical errors. Return to step 1 as necessary.

## Dev

### Building

Use the [`buf`](https://buf.build/) tool to compile the API to an [image](https://buf.build/docs/inputs). First, [install `buf`](https://buf.build/docs/installation), then run:

```
> buf image build -o /dev/null
```

to test the build. To output the image in json format, run:

```
> buf image build --exclude-source-info -o -#format=json
```

### Linting

`buf` runs several [linters](https://buf.build/docs/lint-checkers).

```
> buf check lint
```

This command should have exit code 0 and no output. See the [style guide](https://buf.build/docs/style-guide).

