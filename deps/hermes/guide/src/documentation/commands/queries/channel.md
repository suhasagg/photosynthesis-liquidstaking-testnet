# Table of Contents

<!-- toc -->

# Query Channels

Use the `query channels` command to query the identifiers of all channels on a
given chain.

```shell
USAGE:
    hermes query channels [OPTIONS] --chain <CHAIN_ID>

DESCRIPTION:
    Query the identifiers of all channels on a given chain

OPTIONS:
        --counterparty-chain <COUNTERPARTY_CHAIN_ID>
            Filter the query response by the this counterparty chain

    -h, --help
            Print help information

        --show-counterparty
            Show the counterparty chain, port, and channel

        --verbose
            Enable verbose output, displaying the client and connection ids for each channel in the
            response

REQUIRED:
        --chain <CHAIN_ID>    Identifier of the chain to query
```

**Example**

Query all channels on `ibc-1`:

```shell
hermes query channels --chain ibc-1
```

```json
Success: [
    PortChannelId {
        channel_id: ChannelId(
            "channel-0",
        ),
        port_id: PortId(
            "transfer",
        ),
    },
    PortChannelId {
        channel_id: ChannelId(
            "channel-1",
        ),
        port_id: PortId(
            "transfer",
        ),
    },
]
```

# Query Channel Data

Use the `query channel` commands to query the information about a specific
channel.

```shell
USAGE:
    hermes query channel <SUBCOMMAND>

DESCRIPTION:
    Query information about channels

SUBCOMMANDS:
    client     Query channel's client state
    end        Query channel end
    ends       Query channel ends and underlying connection and client objects
```

## Query the channel end data

Use the `query channel end` command to query the channel end:

```shell
USAGE:
    hermes query channel end [OPTIONS] --chain <CHAIN_ID> --port <PORT_ID> --channel <CHANNEL_ID>

DESCRIPTION:
    Query channel end

OPTIONS:
        --height <HEIGHT>    Height of the state to query

REQUIRED:
        --chain <CHAIN_ID>        Identifier of the chain to query
        --channel <CHANNEL_ID>    Identifier of the channel to query [aliases: chan]
        --port <PORT_ID>          Identifier of the port to query
```

**Example**

Query the channel end of channel `channel-1` on port `transfer` on `ibc-1`:

```shell
hermes query channel end --chain ibc-1 --port transfer --channel channel-1
```

```json
Success: ChannelEnd {
    state: Open,
    ordering: Unordered,
    remote: Counterparty {
        port_id: PortId(
            "transfer",
        ),
        channel_id: Some(
            ChannelId(
                "channel-0",
            ),
        ),
    },
    connection_hops: [
        ConnectionId(
            "connection-1",
        ),
    ],
    version: "ics20-1",
}
```

## Query the channel data for both ends of a channel

Use the `query channel ends` command to obtain both ends of a channel:

```shell
USAGE:
    hermes query channel ends [OPTIONS] --chain <CHAIN_ID> --port <PORT_ID> --channel <CHANNEL_ID>

DESCRIPTION:
    Query channel ends and underlying connection and client objects

OPTIONS:
        --height <HEIGHT>    Height of the state to query
        --verbose            Enable verbose output, displaying all details of channels, connections
                             & clients

REQUIRED:
        --chain <CHAIN_ID>        Identifier of the chain to query
        --channel <CHANNEL_ID>    Identifier of the channel to query [aliases: chan]
        --port <PORT_ID>          Identifier of the port to query
```

**Example**

Query the channel end of channel `channel-1` on port `transfer` on `ibc-0`:

```shell
hermes query channel ends --chain ibc-0 --port transfer --channel channel-1
```

```json
Success: ChannelEndsSummary {
    chain_id: ChainId {
        id: "ibc-0",
        version: 0,
    },
    client_id: ClientId(
        "07-tendermint-1",
    ),
    connection_id: ConnectionId(
        "connection-1",
    ),
    channel_id: ChannelId(
        "channel-1",
    ),
    port_id: PortId(
        "transfer",
    ),
    counterparty_chain_id: ChainId {
        id: "ibc-2",
        version: 2,
    },
    counterparty_client_id: ClientId(
        "07-tendermint-1",
    ),
    counterparty_connection_id: ConnectionId(
        "connection-1",
    ),
    counterparty_channel_id: ChannelId(
        "channel-1",
    ),
    counterparty_port_id: PortId(
        "transfer",
    ),
}
```

Passing the `--verbose` flag will additionally print all the details of the
channel, connection, and client on both ends.

## Query the channel client state

Use the `query channel client` command to obtain the channel's client state:

```shell
USAGE:
    hermes query channel client --chain <CHAIN_ID> --port <PORT_ID> --channel <CHANNEL_ID>

DESCRIPTION:
    Query channel's client state

REQUIRED:
        --chain <CHAIN_ID>        Identifier of the chain to query
        --channel <CHANNEL_ID>    Identifier of the channel to query [aliases: chan]
        --port <PORT_ID>          Identifier of the port to query
```

If the command is successful a message with the following format will be
displayed:

    Success: Some(
        IdentifiedAnyClientState {
            client_id: ClientId(
                "07-tendermint-0",
            ),
            client_state: Tendermint(
                ClientState {
                    chain_id: ChainId {
                        id: "network2",
                        version: 0,
                    },
                    trust_level: TrustThreshold {
                        numerator: 1,
                        denominator: 3,
                    },
                    trusting_period: 1209600s,
                    unbonding_period: 1814400s,
                    max_clock_drift: 40s,
                    latest_height: Height {
                        revision: 0,
                        height: 2775,
                    },
                    proof_specs: ProofSpecs(
                        [
                            ProofSpec(
                                ProofSpec {
                                    leaf_spec: Some(
                                        LeafOp {
                                            hash: Sha256,
                                            prehash_key: NoHash,
                                            prehash_value: Sha256,
                                            length: VarProto,
                                            prefix: [
                                                0,
                                            ],
                                        },
                                    ),
                                    inner_spec: Some(
                                        InnerSpec {
                                            child_order: [
                                                0,
                                                1,
                                            ],
                                            child_size: 33,
                                            min_prefix_length: 4,
                                            max_prefix_length: 12,
                                            empty_child: [],
                                            hash: Sha256,
                                        },
                                    ),
                                    max_depth: 0,
                                    min_depth: 0,
                                },
                            ),
                            ProofSpec(
                                ProofSpec {
                                    leaf_spec: Some(
                                        LeafOp {
                                            hash: Sha256,
                                            prehash_key: NoHash,
                                            prehash_value: Sha256,
                                            length: VarProto,
                                            prefix: [
                                                0,
                                            ],
                                        },
                                    ),
                                    inner_spec: Some(
                                        InnerSpec {
                                            child_order: [
                                                0,
                                                1,
                                            ],
                                            child_size: 32,
                                            min_prefix_length: 1,
                                            max_prefix_length: 1,
                                            empty_child: [],
                                            hash: Sha256,
                                        },
                                    ),
                                    max_depth: 0,
                                    min_depth: 0,
                                },
                            ),
                        ],
                    ),
                    upgrade_path: [
                        "upgrade",
                        "upgradedIBCState",
                    ],
                    allow_update: AllowUpdate {
                        after_expiry: true,
                        after_misbehaviour: true,
                    },
                    frozen_height: None,
                },
            ),
        },
    )

**JSON:**

```shell
    hermes --json query channel client --chain <CHAIN_ID> --port <PORT_ID> --channel <CHANNEL_ID>
```

If the command is successful a message with the following format will be
displayed:

```json
{
    "result":
    {
        "client_id":"07-tendermint-0",
        "client_state":
        {
            "allow_update":
            {
                "after_expiry":true,
                "after_misbehaviour":true
            },
            "chain_id":"network2",
            "frozen_height":null,
            "latest_height":
            {
                "revision_height":2775,
                "revision_number":0
            },
            "max_clock_drift":
            {
                "nanos":0,
                "secs":40
            },
            "proof_specs":
            [
                {
                    "inner_spec":
                    {
                        "child_order":[0,1],
                        "child_size":33,
                        "empty_child":"",
                        "hash":1,
                        "max_prefix_length":12,
                        "min_prefix_length":4
                    },
                    "leaf_spec":
                    {
                        "hash":1,
                        "length":1,
                        "prefix":"AA==",
                        "prehash_key":0,
                        "prehash_value":1
                    },
                    "max_depth":0,
                    "min_depth":0
                },
                {
                    "inner_spec":
                    {
                        "child_order":[0,1],
                        "child_size":32,
                        "empty_child":"",
                        "hash":1,
                        "max_prefix_length":1,
                        "min_prefix_length":1
                    },
                    "leaf_spec":
                    {
                        "hash":1,
                        "length":1,
                        "prefix":"AA==",
                        "prehash_key":0,
                        "prehash_value":1
                    },
                    "max_depth":0,
                    "min_depth":0
                }
            ],
            "trust_level":
            {
                "denominator":3,
                "numerator":1
            },
            "trusting_period":
            {
                "nanos":0,
                "secs":1209600
            },
            "type":"Tendermint",
            "unbonding_period":
            {
                "nanos":0,
                "secs":1814400
            },
            "upgrade_path":["upgrade","upgradedIBCState"]
        },
        "type":"IdentifiedAnyClientState"
    },
    "status":"success"
}
```
