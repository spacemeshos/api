# Spacemesh API
[Protobuf](https://developers.google.com/protocol-buffers) implementation of the Spacemesh API. This is a work in progress. See the more complete [master API spec](https://docs.google.com/spreadsheets/d/1P89OVWdgJocPy0CGM43Ge7Sx_6dabCBEagaVQfOk9us/edit).

## Design

The Spacemesh API consists of several logical services, each of which contains a set of one or more RPC endpoints. The node operator can enable or disable each service independently using the CLI. The current set of services is as follows:

- [NodeService](proto/spacemesh/node.proto) is a readonly interface for reading basic node-related data such as node status, software version and build number, and errors. It also allows a consumer to request that the node start the sync process, thus enabling the stream endpoints.
- [MeshService](proto/spacemesh/mesh.proto) is a readonly interface that provides access to mesh data such as layer and epoch number, and network ID. It provides streams for watching layers (which contain blocks, transactions, etc.). In the future this service will be expanded to include other mesh-related endpoints.
- [GlobalStateService](proto/spacemesh/global_state.proto) is a readonly interfact that provides access to data elements that are not explicitly part of the mesh such as accounts, rewards, and transaction state and receipts. In the future this service will be expanded to include additional endpoints for things such as global state hash and events emitted by smart contracts.

In addition to these services, there is also a set of [global types](proto/spacemesh/types.proto) which are shared among all of the services.
 
## Intended Usage Pattern

### Mesh data processing flow
1. Client starts a full node with flags set to turn syncing off and to open the GRPC APIs
1. Client registers on the streaming GRPC api methods that are of interest
1. Client calls `NodeSyncStart()` to request that the node start syncing
1. Client processes streaming data it receives from the node
1. Client monitors node using `node.SyncStatusStream()` and `node.ErrorStream()` and handle node critical errors. Return to step 1 as necessary.
1. Client gracefully shuts down the node by calling `node.Shutdown()` when it is done processing data.

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

