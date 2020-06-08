# Spacemesh API
[Protobuf](https://developers.google.com/protocol-buffers) implementation of the Spacemesh API. This is a work in progress. See the more complete [master API spec](https://docs.google.com/spreadsheets/d/1P89OVWdgJocPy0CGM43Ge7Sx_6dabCBEagaVQfOk9us/edit).

## Design

The API was designed with the following considerations in mind.

### Mesh vs. global state

In Spacemesh, "the mesh" refers to data structures that are _explicitly_ stored by all full nodes, and are subject to consensus. This consists of transactions, collated into blocks, which in turn are collated into layers. Note that, in addition to transactions, blocks contain metadata such as layer number and signature. The mesh also includes ATXs (activations).

By contrast, "global state" refers to data structures that are calculated _implicitly_ based on mesh data. These data are not stored anywhere in the mesh explicitly. Global state includes account state (balance, counter/nonce value, and, for smart contract accounts, code), transaction receipts, and smart contract event logs. These data need not be stored indefinitely by all full nodes (although they should be stored indefinitely by archive nodes).

The API provides access to both types of data, but they are divided into different API services. For more information on this distinction, see [SMIP-0003: Global state data, STF, APIs](https://github.com/spacemeshos/SMIPS/issues/13).

### Transactions

Transactions span mesh and global state data. They are submitted to a node, which may or may not admit the transaction to its mempool. If the transaction is admitted to the mempool, it will probably end up being added to a newly-mined block, and that block will be submitted to the mesh in some layer. After that, the layer containing the block will eventually be approved, and then confirmed, by the consensus mechanism. After the layer is approved, the transaction will be run through the STF (state transition function), and if it succeeds, it may update global state.

Since transactions span multiple layers of abstraction, the API exposes transaction data in its own service, TransactionService.

### Queries and streams

Most API endpoints are one of two types: a query, or a stream. This is an important distinction, and in many cases, the same data are exposed through both a query and a stream endpoint.

- **Queries** are used to read paginated historical data. A `*Query` endpoint accepts a `*Params` message which typically contains the following:
    - `filter`: A filter (see Streams, below)
    - `min_layer`: The first layer to return results from
    - `max_results`: The maximum number of results to return
    - `offset`: Page offset
- **Streams** are used to read realtime data. They do not return historical data. Each time the node creates, or learns of, a piece of data matching the filter and type, or sees an update to a matching piece of data, it sends it over the stream. A `*Stream` endpoint accepts a `*Filter` message which typically contains the following:
    - `*_id`: The ID of the data type to filter on (e.g., "show me all data items that touch this `account_id`")
    - `flags`: A bit field that allows the client to select which, among multiple types multiplexed on this stream, to receive

## Services

The Spacemesh API consists of several logical services, each of which contains a set of one or more RPC endpoints. The node operator can enable or disable each service independently using the CLI. The current set of services is as follows:

- [NodeService](proto/spacemesh/node.proto) is a readonly interface for reading basic node-related data such as node status, software version and build number, and errors. It also allows a consumer to request that the node start the sync process, thus enabling the stream endpoints.
- [MeshService](proto/spacemesh/mesh.proto) is a readonly interface that provides access to mesh data such as layer number, epoch number, and network ID. It provides streams for watching layers (which contain blocks, transactions, etc.). In the future this service will be expanded to include other mesh-related endpoints.
- [GlobalStateService](proto/spacemesh/global_state.proto) is a readonly interfact that provides access to data elements that are not explicitly part of the mesh such as accounts, rewards, and transaction state and receipts. In the future this service will be expanded to include additional endpoints for things such as global state hash and events emitted by smart contracts.

In addition to these services, there is also a set of [global types](proto/spacemesh/types.proto) which are shared among all of the services.

## Intended Usage Pattern

### Mesh data processing flow
1. Client starts a full node with flags set to turn syncing off and to open the GRPC APIs
1. Client registers on the streaming GRPC api methods that are of interest
1. Client calls `node.SyncStart()` to request that the node start syncing
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

### Breaking changes detection

`buf` also supports [detection of breaking changes](https://buf.build/docs/tour-5). To do this, first create an image from the current state:

```
> buf image build -o image.bin
```

Make a breaking change, then run against this change:

```
> buf check breaking --against-input image.bin
```

`buf` will report all breaking changes.

Detection of breaking changes is turned off in GitHub actions [continuous integration](.github/workflows/ci.yml) for now, until the API has stabilized, but it's still good practice to run this check when committing changes.

### Linting

`buf` runs several [linters](https://buf.build/docs/lint-checkers).

```
> buf check lint
```

This command should have exit code 0 and no output. See the [style guide](https://buf.build/docs/style-guide).
