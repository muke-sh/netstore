package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}
// Transport is anything that handles the communication
// between the nodes. This can be of the any form (TCP, UDP, Websockets)
type Transport interface {
	ListenAndAccept() error
}
