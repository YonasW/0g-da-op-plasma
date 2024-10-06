## Optimism 0G Integration Guide

<br />
<p align="center">
  <p align="center" href="https://optimism.io/">
  <img src="https://www.dropbox.com/scl/fi/z0skana9otxqlxqopcofa/Optimism-Red-48.svg?rlkey=hgxop0kdzplncn199kwbqzkui&e=1&dl=0" alt="Logo" width="120" height="120">
  </p>
  <h3 align="center">+</h3>
  <p align="center" href="0g.ia/">
  <img src="https://drive.google.com/file/d/1kPUEcdTAzZgfyIZ9A41XMUE4p6LPYL-H/view" alt="Logo" width="140" height="140">
  </p>
    <br />
  </p>
</p>
### Overview

The Optimism 0G integration allows developers to deploy an OP Stack-based chain using 0G (Zero Gravity) for data availability. This integration offers an alternative to Optimism's default data availability solution, providing a cost-effective and secure option for storing transaction data.

### Key Components

1. DA Server: Implements the data availability server interface for 0G DA.
2. OP Stack Configuration: Customizes the OP Stack components to work with 0G DA.
3. 0G Integration: Ensures data integrity and availability through 0G's network.

### 0G DA Server Implementation

The core logic for posting and retrieving data is implemented in the da-server. Key features include:

- DA Server: Manages the connection to the 0G DA client.
- HTTP Server: Handles requests from OP Stack components for data storage and retrieval.
- Integration with OP Stack: Ensures seamless communication between Optimism components and 0G DA.

### Setting Up Your Chain

1. Deploy da-server:
   - Follow the deployment instructions in the da-server README.

2. Deploy OP Stack components:
   - Modify the `rollup.json` configuration for op-node.
   - Set specific CLI configurations for op-node and op-batcher.

3. Start the system:
   - Launch all components following Optimism's general instructions with the 0G-specific modifications.

### Important Notes

- This integration is in beta, with ongoing active development.
- Ensure proper firewall configuration for all necessary ports.
- Refer to [Optimism documentation](https://docs.optimism.io/) for additional configuration options and troubleshooting.

### Learn More About 0G

[0G Website](https://0g.ai/)
[0G Github](https://github.com/0glabs)

### Learn More About Optimism

[Optimism Documentation](https://docs.optimism.io/)
[OP Stack Github](https://github.com/ethereum-optimism/optimism)

## Guidance to run OP Stack with 0G DA
[How to Use the OP Stack with 0G DA](./OP%20Stack%20integration.md)
