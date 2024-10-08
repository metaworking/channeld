# Core features
- [x] Channel pub/sub
- [x] Data update and fan-out
- [x] FSM-based message filtering
- [x] Message broadcasting
- [x] Authentication
- [x] Channel ACL
- [x] DDoS Protection
- [x] Spatial-based pub/sub
- [ ] Health check
- [ ] Disaster recovery
- [ ] Spatial-based load-balancing
- [ ] Distributed channels

# Modules
- [x] Stub(RPC) support
- [x] WebSocket support
- [x] KCP support
- [x] [Snappy](https://github.com/golang/snappy) compression
- [ ] [Markov-chain](https://en.wikipedia.org/wiki/Markov_chain) compression
- [ ] Encryption
- [x] Replay
- [x] Prometheus integration

# Optimizations
- [x] Read/write the packet using Protobuf
- ~~[ ] Use [gogoprotobuf](https://github.com/gogo/protobuf) for faster marshalling/unmarshalling~~
- [x] Enable custom merge of channel data messages
- [ ] Non-reflection-based merge of channel data messages

# Tests
- [x] Unit tests
- [x] Benchmark tests

# SDKs
- [ ] Javascript SDK
- [x] Unity C# SDK
- [x] Unreal C++ SDK

# Tools
- [x] Simulated client (Go)
- [x] Replication code generators
- [ ] Channel monitor

# Example projects
- [x] Web chat rooms
    - [x] Implement the Javascript client library
    - [x] Implement the commands
    - [ ] Scale test with 10K connections
    - [ ] Complete the UI
- [ ] Unity tank game
    - [x] Implement the C# client library
    - [x] Mirror Integration
        - [x] Transport
        - [x] SyncVar and NetworkTransform
        - [x] Observers and Interest Management
        - [ ] SyncVar and RPC code generation
    - [x] Multi-server support
- [x] Unreal seamless world travelling
    - [x] Implement the C++ client library
    - [x] Blueprint support
    - [x] Integrate with Unreal's networking stack
    - [x] Integrate with Unreal's Replication system
    - [x] Replication codegen
    - [x] Multi-server support
    - [x] Client interest management
    - [x] Editor toolbar extension
    - [ ] KCP support
    - [ ] Data traffic compression
- [ ] Dynamic region load-balancing
