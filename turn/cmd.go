package turn

// Cmd for turn cmd
type Cmd int

const (
	// CmdCreateAllocatin for CreateAllocatin
	CmdCreateAllocatin Cmd = iota
	// CmdRefreshAllocation for RefreshAllocation
	CmdRefreshAllocation
	// CmdGetServerReflexiveAddr for GetServerReflexiveAddr
	CmdGetServerReflexiveAddr
	// CmdCreatePermission for CreatePermission
	CmdCreatePermission
	// CmdSendData for SendData
	CmdSendData
	// CmdChannelBind for ChannelBind
	CmdChannelBind
)

// Request for turn request
type Request struct {
	TransactionID uint64
	Cmd           Cmd
	Param         []byte
}

// Response for turn response
type Response struct {
	TransactionID uint64
	Result        []byte
}
