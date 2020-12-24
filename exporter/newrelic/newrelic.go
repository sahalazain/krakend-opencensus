package newrelic

import (
	"context"
	"errors"

	opencensus "github.com/devopsfaith/krakend-opencensus"
	"github.com/newrelic/newrelic-opencensus-exporter-go/nrcensus"
)

func init() {
	opencensus.RegisterExporterFactories(func(ctx context.Context, cfg opencensus.Config) (interface{}, error) {
		return Exporter(ctx, cfg)
	})
}

func Exporter(ctx context.Context, cfg opencensus.Config) (*nrcensus.Exporter, error) {
	if cfg.Exporters.NewRelic == nil {
		return nil, errDisabled
	}
	e, err := nrcensus.NewExporter(cfg.Exporters.NewRelic.ServiceName, cfg.Exporters.NewRelic.APIKey)
	if err != nil {
		return e, err
	}

	return e, nil

}

var errDisabled = errors.New("opencensus newrelic exporter disabled")
