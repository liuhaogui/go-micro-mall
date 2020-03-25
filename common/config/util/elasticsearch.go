package util

type ElasticSearch struct {
	Enabled  bool   `json:"enabled"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}
