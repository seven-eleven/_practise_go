// command.go
package common

// Client与Server交互的消息类型
const (
	CmdUpdate     = "update"     // client -> server
	CmdQuery      = "query"      // client -> server
	CmdStop       = "stop"       // client -> server
	CmdAck4Update = "ack4update" // server -> client
	CmdAck4Query  = "ack4query"  // server -> client
)
