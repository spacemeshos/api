# Spacemesh API
![](https://github.com/spacemeshos/api/workflows/CI/badge.svg)

[Protobuf](https://developers.google.com/protocol-buffers) implementation of the Spacemesh API. This repository contains only API design, not implementation. For implementation work, see [go-spacemesh](https://github.com/spacemeshos/go-spacemesh/tree/develop/api/grpcserver). Note that API implementation may lag design.

## Design

The API was designed with the following considerations in mind.

### Mesh vs. global state

In Spacemesh, "the mesh" refers to data structures that are _explicitly_ stored by all full nodes and are subject to consensus. This consists of transactions, collated into blocks, which in turn are collated into layers. Note that, in addition to transactions, blocks contain metadata such as layer number and signature. The mesh also includes ATXs (activations).

By contrast, "global state" refers to data structures that are calculated _implicitly_ based on mesh data. These data are not explicitly stored anywhere in the mesh. Global state includes account state (balance, counter/nonce value, and, for smart contract accounts, code), transaction receipts, and smart contract event logs. These data need not be stored indefinitely by all full nodes (although they should be stored indefinitely by archive nodes).

The API provides access to both types of data, but they are divided into different API services. For more information on this distinction, see [SMIP-0003: Global state data, STF, APIs](https://github.com/spacemeshos/SMIPS/issues/13), as well as the [`MeshService`](/proto/spacemesh/v1/mesh.proto) and the [`GlobalStateService`](/proto/spacemesh/v1/global_state.proto).

### Transactions

Transactions span mesh and global state data. They are submitted to a node, which may or may not admit the transaction to its mempool. If the transaction is admitted to the mempool, it will probably end up being added to a newly-mined block, and that block will be submitted to the mesh in some layer. After that, the layer containing the block will eventually be approved, and then confirmed, by the consensus mechanism. After the layer is approved, the transaction will be run through the STF (state transition function), and if it succeeds, it may update global state.

Since transactions span multiple layers of abstraction, the API exposes transaction data in its own service, [`TransactionService`](/proto/spacemesh/v1/tx.proto).

### Types of endpoints

Broadly speaking, there are four types of endpoints: simple, command, query, and stream. Each type is described below. Note that in some cases, the same data are exposed through multiple endpoints, e.g., both a query and a stream endpoint.

- **Simple** endpoints are used to query a single data element. Some simple endpoints accept a request object (e.g., [`GlobalStateService.Account`](https://github.com/spacemeshos/api/blob/ab285df4d52af4663335a4bcccb0b52f1e5003ee/proto/spacemesh/v1/global_state.proto#L14)), while some which return data that's global to a node accept no request object (e.g., [`NodeService.Version`](https://github.com/spacemeshos/api/blob/ab285df4d52af4663335a4bcccb0b52f1e5003ee/proto/spacemesh/v1/node.proto#L13)).
- **Command** endpoints are used to send a command to a node. Examples include [`TransactionService.SubmitTransaction`](https://github.com/spacemeshos/api/blob/ab285df4d52af4663335a4bcccb0b52f1e5003ee/proto/spacemesh/v1/tx.proto#L15) and [`SmesherService.StartSmeshing`](https://github.com/spacemeshos/api/blob/ab285df4d52af4663335a4bcccb0b52f1e5003ee/proto/spacemesh/v1/smesher.proto#L16).
- **Query** endpoints are used to read paginated historical data. A `*Query` endpoint accepts a `*Request` message that typically contains the following:
    - `filter`: A filter (see Streams, below)
    - `min_layer`: The first layer to return results from
    - `max_results`: The maximum number of results to return
    - `offset`: Page offset
- **Stream** endpoints are used to read realtime data. They do not return historical data. Each time the node creates, or learns of, a piece of data matching the filter and type, or sees an update to a matching piece of data, it sends it over the stream. A `*Stream` endpoint accepts a `*Request` message, containing the following, that functions as a filter:
    - `*_id`: The ID of the data type to filter on (e.g., "show me all data items that touch this `account_id`")
    - `flags`: A bit field that allows the client to select which, among multiple types multiplexed on this stream, to receive

## Services

The Spacemesh API consists of several logical services, each of which contains a set of one or more RPC endpoints. The node operator can enable or disable each service independently using the CLI. The current set of services is as follows:

- [GatewayService](/proto/spacemesh/v1/gateway.proto) is a read-write interface that allows a [poet server](https://github.com/spacemeshos/poet) to broadcast data to the network via a poet gateway node.
- [GlobalStateService](/proto/spacemesh/v1/global_state.proto) is a readonly interface that provides access to data elements that are not explicitly part of the mesh such as accounts, rewards, and transaction state and receipts.
- [MeshService](/proto/spacemesh/v1/mesh.proto) is a readonly interface that provides access to mesh data such as layer number, epoch number, and network ID. It provides streams for watching layers (which contain blocks, transactions, etc.). In the future this service will be expanded to include other mesh-related endpoints.
- [NodeService](/proto/spacemesh/v1/node.proto) is a readonly interface for reading basic node-related data such as node status, software version and build number, and errors. It also allows a consumer to request that the node start the sync process, thus enabling the stream endpoints.
- [SmesherService](/proto/spacemesh/v1/smesher.proto) is a read-write interface that allows the client to query, and set, parameters related to smeshing (mining), such as PoST commitment, coinbase, etc.
- [TransactionService](/proto/spacemesh/v1/tx.proto) is a read-write interface that allows the client to submit a new transaction, and to follow the state of one or more transactions on its journey from submission to mempool to block to mesh to STF.

Each of these services relies on one or more sets of message types, which live in `*types.proto` files in the same directory as the service definition files.

## Intended Usage Pattern

### Mesh data processing flow

1. Client starts a full node with one or more relevant GRPC endpoints enabled
1. Client subscribes to the streaming GRPC api methods that are of interest
1. Client calls `NodeService.SyncStart()` to request that the node start syncing (note that, at present, sync is on by default and this step is unnecessary, but in future, it will be possible to [start the node with sync turned off](https://github.com/spacemeshos/go-spacemesh/issues/2080) so that the client can subscribe to streams before the sync process begins, ensuring they don't miss any data)
1. Client processes streaming data it receives from the node
1. Client monitors node using `NodeService.SyncStatusStream()` and `NodeService.ErrorStream()` and handles node critical errors. Return to step 1 as necessary.
1. Client gracefully shuts down the node by calling `NodeService.Shutdown()` when it is done processing data.

## Development

### Versioning

We use standard [semantic versioning](https://semver.org/). Please regularly cut releases against the `master` branch and increment the version accordingly. Releases are managed at [Releases](https://github.com/spacemeshos/api/releases) and the current version line is 1.x. Note that this is especially important for downstream code that relies on individual builds, such as the golang build.

### Build targets

This repository currently contains builds for two targets: golang and grpc-gateway. **Every time a protobuf definition file is changed, you must update the build** and include the updated build files with your PR in order to keep everything in sync. You can check this at any time by running `make check`, and it's also enforced by CI (see below for more information).

- golang builds live in [`release/go`](/release/go). You may use this repository directly as a go module with an import statement such as `import "github.com/spacemeshos/api/release/go/spacemesh/v1"`.
- grpc-gateway builds are configured in [`proto/spacemesh/v1/api_config.yaml`](/proto/spacemesh/v1/api_config.yaml) and live alongside the golang builds in [`release/go/spacemesh/v1`](/release/go/spacemesh/v1) (they have a `.gw.go` extension).

### Makefile

The repository includes a [`Makefile`](Makefile) that makes it easy to run most regular tasks:

- `make deps` installs the required dependencies, including `buf` (see below)
- `make lint` runs the linter (see below)
- `make local` runs the linter and checks for breaking changes (see below)
- `make build` builds the API for all targets
- `make check` ensures that the build is up to date with respect to the proto source files

Under the hood, it uses a helpful tool called `buf`.

### Buf

In addition to running `make` commands, you can also manually use the [`buf`](https://buf.build/) tool to compile the API to an [image](https://buf.build/docs/inputs). First, [install `buf`](https://buf.build/docs/installation), then run:

```
> buf image build -o /dev/null
```

to test the build. To output the image in json format, run:

```
> buf image build --exclude-source-info -o -#format=json
```

#### Breaking changes detection

`buf` also supports [detection of breaking changes](https://buf.build/docs/tour-5). To do this, first create an image from the current state:

```
> buf image build -o image.bin
```

Make a breaking change, then run against this change:

```
> buf check breaking --against-input image.bin
```

`buf` will report all breaking changes.

#### Linting

`buf` runs several [linters](https://buf.build/docs/lint-checkers). It's pretty strict about things such as naming conventions, to prevent downstream issues in the various languages and framework that rely upon the protobuf definition files. You can run the linter like this:

```
> buf check lint
```

If there are no issues, this command should have exit code 0 and no output.

For more information on linting, see the [style guide](https://buf.build/docs/style-guide). For more information on the difference between the `buf` tool and the `protoc` compiler, see [Use protoc input instead of the internal compiler](https://buf.build/docs/tour-7).

### Continuous integration

This repository has a continuous integration (CI) [workflow](.github/workflows/ci.yml) built on GitHub Actions. In addition to linting and breaking changes detection, it also runs the `protoc` compiler, since that tends to surface a slightly different set of warnings and errors than `buf`.

You can use a nifty tool called [`act`](https://github.com/nektos/act) to run the CI workflow locally, although it [doesn't always play nice](https://github.com/nektos/act/issues/322) with our workflow configuration.

### Third party files

The `third_party/` directory includes several third-party libraries, which are required by some of the Google extensions used in the API definition files. These have manually been copied in from two sources:

- https://github.com/googleapis/googleapis/tree/master/google (`google.api`, `google.rpc`)
- https://github.com/protocolbuffers/protobuf/tree/master/src/google/protobuf (`google.protobuf`)

These files do not change often and probably do not need to be updated. However, if updates were to be more common, it might make more sense to add a dynamic dependency on these external libraries using a Makefile or `git submodule`.
