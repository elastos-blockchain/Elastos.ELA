package config

var configTemplate = Configuration{
	Magic:               7630401,
	Version:             23,
	SeedList:            []string{"127.0.0.1:30338"},
	HttpInfoPort:        20333,
	HttpInfoStart:       true,
	HttpRestPort:        20334,
	HttpWsPort:          20335,
	WsHeartbeatInterval: 60,
	HttpJsonPort:        20336,
	NodePort:            20338,
	NodeOpenPort:        20866,
	OpenService:         true,
	PrintLevel:          0,
	MaxLogsSize:         0,
	MaxPerLogSize:       0,
	IsTLS:               false,
	CertPath:            "./sample-cert.pem",
	KeyPath:             "./sample-cert-key.pem",
	CAPath:              "./sample-ca.pem",
	MultiCoreNum:        4,
	MaxTxsInBlock:       10000,
	MaxBlockSize:        8000000,
	MinCrossChainTxFee:  10000,
	PowConfiguration: PowConfiguration{
		PayToAddr:  "8VYXVxKKSAxkmRrfmGpQR2Kc66XhG6m3ta",
		AutoMining: false,
		MinerInfo:  "ELA",
		MinTxFee:   100,
		ActiveNet:  "RegNet",
	},
	EnableArbiter: false,
	ArbiterConfiguration: ArbiterConfiguration{
		Name:  "test",
		Magic: 7630403,
		NodePort:         30338,
		ProtocolVersion:  0,
		Services:         0,
		PrintLevel:       1,
		SignTolerance:    5,
		MaxLogsSize:      0,
		MaxPerLogSize:    0,
		MaxConnections:   100,
		MajorityCount:    3,
		ArbitratorsCount: 5,
		CandidatesCount:  0,
	},
}