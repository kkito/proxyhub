package main

// IHostDestClassifier check host type
type IHostDestClassifier interface {
	// try using  "geoip:private", "geoip:cn" etc
	isInternal() bool
	isCN() bool
	isWallBlock() bool
}

// IProxyChannel a proxy channel can use
type IProxyChannel interface {
	// how to invoke?
}

// IProxyChannelBenchmark can choose a best proxy channel
type IProxyChannelBenchmark interface {
}
