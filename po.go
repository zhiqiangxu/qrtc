package qrtc

type (
	// UintRange for uint range
	UintRange struct {
		Min, Max uint
	}
	// Float32Range for float32 range
	Float32Range struct {
		Min, Max float32
	}
	// MediaTrackCapabilities represents the Capabilities of a MediaStreamTrack object
	MediaTrackCapabilities struct {
		Width        UintRange
		Height       UintRange
		AspectRatio  Float32Range
		FrameRate    Float32Range
		SampleRate   UintRange
		SampleSize   UintRange
		Latency      Float32Range
		ChannelCount UintRange
	}
)
