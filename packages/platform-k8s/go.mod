module code-code.internal/platform-k8s

go 1.26.2

require (
	code-code.internal/go-contract v0.0.0
	github.com/nats-io/nats.go v1.50.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/klauspost/compress v1.18.5 // indirect
	github.com/nats-io/nkeys v0.4.15 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.49.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
)

replace code-code.internal/go-contract => ../../code-code-contracts/packages/go-contract
