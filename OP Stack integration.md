# OP Stack Integration

[OP Stack](https://stack.optimism.io/) is the set of [software components](https://github.com/ethereum-optimism/optimism) that run the [Optimism](https://www.optimism.io/) rollup and can be deployed independently to power third-party rollups.

By default, OP Stack sequencers write batches to Ethereum in the form of calldata or 4844 blobs to commit to the transactions included in the canonical L2 chain. In Alt-DA mode, OP Stack sequencers and full nodes are configured talk to a third-party HTTP server for writing and reading tx batches to and from DA. Optimism's Alt-DA spec contains a more in-depth breakdown of how this works.

To implement this server spec, 0G DA provides the da-server, which can be run alongside OP Stack sequencers and full nodes to securely communicate with the 0G DA Client.

## Deploying

### Deploy da-server

[Run DA Server](./README.md#Deployment)

### Deploying OP Stack

Next deploy the OP Stack components according to the official OP Stack [deployment docs](https://docs.optimism.io/builders/chain-operators/tutorials/create-l2-rollup), but with the following modifications:


#### op-node rollup.json configuration

In the op-node `rollup.json` configuration the following should be set:

```json
{
  "plasma_config": {
    "da_challenge_contract_address": "0x0000000000000000000000000000000000000000",
    "da_commitment_type": "GenericCommitment",
    "da_challenge_window": 300,
    "da_resolve_window": 300
  }
}
```

#### op-node CLI configuration

The following configuration values should be set to ensure proper communication between op-node and da-server.

```
    --plasma.enabled=true
    --plasma.da-server=http://localhost:3100
    --plasma.verify-on-read=false
    --plasma.da-service=true
```

#### op-batcher CLI configuration

The following configuration values should be set accordingly to ensure proper communication between OP Batcher and da-server.

```
    --plasma.enabled=true
    --plasma.da-server=http://localhost:3100
    --plasma.verify-on-read=false
    --plasma.da-service=true
```

## Reference
[How to Run an Alt-DA Mode Chain](https://docs.optimism.io/builders/chain-operators/features/alt-da-mode)
