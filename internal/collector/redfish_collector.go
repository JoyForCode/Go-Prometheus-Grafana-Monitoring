package collector

import "github.com/prometheus/client_golang/prometheus"

type RedfishAPI interface {
	Get(endpoint string, target interface{}) (int, error)
}

type RedfishCollector struct {
	client RedfishAPI

	powerConsumedDesc *prometheus.Desc
	powerCapacityDesc *prometheus.Desc
	redfishUpDesc     *prometheus.Desc
}
