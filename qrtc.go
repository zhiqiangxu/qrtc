package qrtc

// MediaDevices provides access to connected media input devices like cameras and microphones, as well as screen sharing
type MediaDevices interface {
	// GetUserMedia produces a MediaStream with tracks containing the requested types of media
	GetUserMedia() MediaStream
}

// MediaStream is used to group several MediaStreamTrack objects into one unit that can be recorded or rendered in a media element
type MediaStream interface {
	Active() bool
	GetTracks() []MediaStreamTrack
	GetAudioTracks() []MediaStreamTrack
	GetVideoTracks() []MediaStreamTrack
	AddTrack(MediaStreamTrack)
	RemoveTrack(MediaStreamTrack)
	Clone() MediaStream
}

// MediaStreamTrack object represents media of a single type that originates from one media source in the User Agent, e.g. video produced by a web camera
type MediaStreamTrack interface {
	Stop()
	GetCapabilities() MediaTrackCapabilities
	GetConstraints()
	Clone() MediaStreamTrack
}

// PeerConnection for peeer connection
type PeerConnection interface {
	SetLocalDescription(desc SessionDescription) error
	SetRemoteDescription(desc SessionDescription) error
	CreateAnswer() (SessionDescription, error)
	CreateOffer() (SessionDescription, error)
	AddTrack(MediaStreamTrack, MediaStream)
	OnTrack()
}

// TurnClient for turn client
type TurnClient interface {
	CreateAllocation(TurnCreateAllocationRequest) TurnCreateAllocationResponse
	RefreshAllocation(TurnRefreshAllocationRequest) TurnRefreshAllocationResponse
	GetServerReflexiveAddr() TurnGetServerReflexiveAddrResponse
	CreatePermission(TurnCreatePermissionRequest) TurnCreatePermissionResponse
	SendData(TurnSendDataRequest) error // error is generated locally
	OnData(func(TurnPushedData))
	ChannelBind(TurnChannelBindRequest) TurnChannelBindResponse
}
