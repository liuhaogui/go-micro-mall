package util

type Hystrix struct {
	DefaultTimeout               int `json:"defaultTimeout"`
	DefaultMaxConcurrent         int `json:"defaultMaxConcurrent"`
	DefaultVolumeThreshold       int `json:"defaultVolumeThreshold"`
	DefaultSleepWindow           int `json:"defaultSleepWindow"`
	DefaultErrorPercentThreshold int `json:"defaultErrorPercentThreshold"`
}
