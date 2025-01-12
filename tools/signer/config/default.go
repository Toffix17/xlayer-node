package config

// DefaultValues is the default configuration
const DefaultValues = `
Port = 8080

[L1]
ChainId = 11155111
RPC = "https://rpc.ankr.com/eth_sepolia/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
PolygonZkEVMAddress = "0x812cB73e48841a6736bB94c65c56341817cE6304"
GlobalExitRootManagerAddr = "0x0e9Bb928351a50227ebFEC9782Db005Ba9b6C052"
DataCommitteeAddr = "0x246EcFCae4423631c9eE3A86DE37F77BCF27FAaE"
PolygonMaticAddress = "0xe223519d64C0A49e7C08303c2220251be6b70e1d"
SeqPrivateKey = {Path = "../../test/sequencer.keystore", Password = "testonly"}
AggPrivateKey = {Path = "../../test/aggregator.keystore", Password = "testonly"}

[Log]
Environment = "development"
Level = "debug"
Outputs = ["stdout"]
`
