package server

import (
	"github/kunhou/simple-backend/pkg/constant"

	"github.com/prometheus/client_golang/prometheus"
)

func initVersionInfo() {
	appVersion := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "go_app_version_info",
			Help: "app version",
		},
		[]string{"version", "commit_sha", "build_date"},
	)
	appVersion.WithLabelValues(constant.Version, constant.GitCommitSha, constant.BuildDate).Set(1)
	prometheus.MustRegister(appVersion)
}
