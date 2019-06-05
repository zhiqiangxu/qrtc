package qrtc

// TurnError for turn error
type TurnError int

const (
	// TurnErrorOK for ok
	TurnErrorOK TurnError = iota
	// TurnTimeout when timeout
	TurnTimeout
)

type (
	// TransportAddr for transport addr
	TransportAddr struct {
		IP   string
		Port uint16
	}

	// TurnBaseResp for turn base resp
	TurnBaseResp struct {
		Error TurnError
	}
	// TurnCreateAllocationRequest for CreateAllocation
	TurnCreateAllocationRequest struct {
		Peers []TransportAddr
	}
	// TurnCreateAllocationResponse is resp for CreateAllocation
	TurnCreateAllocationResponse struct {
		TurnBaseResp
		*TransportAddr
	}
	// TurnRefreshAllocationRequest for RefreshAllocation
	TurnRefreshAllocationRequest struct {
		// lifetime -1 delete; 0 refresh with origin lifetime
		LifeTime int
	}
	// TurnRefreshAllocationResponse is resp for RefreshAllocation
	TurnRefreshAllocationResponse struct {
		TurnBaseResp
	}
	// TurnGetServerReflexiveAddrResponse is resp for GetServerReflexiveAddr
	TurnGetServerReflexiveAddrResponse struct {
		TurnBaseResp
		*TransportAddr
	}
	// TurnCreatePermissionRequest for CreatePermission
	TurnCreatePermissionRequest struct {
		Peers []TransportAddr
	}
	// TurnCreatePermissionResponse is resp for CreatePermission
	TurnCreatePermissionResponse struct {
		TurnBaseResp
	}
	// TurnSendDataRequest for turn send data
	TurnSendDataRequest struct {
		ToAddrs          []TransportAddr
		ToChannelNumbers []int
		Data             []byte
	}
	// TurnChannelBindRequest for ChannelBind
	TurnChannelBindRequest struct {
		Peers []TransportAddr
	}
	// TurnChannelBindResponse is resp for ChannelBind
	TurnChannelBindResponse struct {
		TurnBaseResp
		ChannelNumbers []int
	}
	// TurnPushedData for OnData
	TurnPushedData struct {
		Peer          TransportAddr
		ChannelNumber int
		Data          []byte
	}
)
