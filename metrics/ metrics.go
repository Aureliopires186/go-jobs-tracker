GO Code "

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var JobsScraped = promauto.NewCounter(prometheus.CounterOpts{
	Name: "jobs_scraped_total",
	Help: "Total jobs scraped",
})
