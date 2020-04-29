# api
[Protobuf](https://developers.google.com/protocol-buffers) implementation of the Spacemesh API. This is a work in progress. See the more complete [master API spec](https://docs.google.com/spreadsheets/d/1P89OVWdgJocPy0CGM43Ge7Sx_6dabCBEagaVQfOk9us/edit).

## Building

Use the [`buf`](https://buf.build/) tool to compile the API to an [image](https://buf.build/docs/inputs). First, [install `buf`](https://buf.build/docs/installation), then run:

```
> buf image build -o /dev/null
```

to test the build. To output the image in json format, run:

```
> buf image build --exclude-source-info -o -#format=json
```

## Linting

`buf` runs several [linters](https://buf.build/docs/lint-checkers).

```
> buf check lint
```

This command should have exit code 0 and no output. See the [style guide](https://buf.build/docs/style-guide).

## API Usage Patterns

### Data processing flow
1. Start a full node with a flag to not start syncing and to open the V2 GRPC APIs.
2. Client Registers on the streaming GRPC api methods that are of interest to it.
3. Client call NodeSyncStart() to request the node to start syncing.
4. Client processes server sides streaming data pushed by the node to it.
5. Client monitors node using NodeSyncStatuses() and NodeErrors() and handle node critical errors. e.g. Got to step 1.
6. Client gracefully shuts down the node by calling NodeShutdown() when it is done processing data.
