package module_traefik_package

import (
	"context"
	"fmt"
	"net/http"

	modulego "github.com/datadome/module-go-package/v2"
)

const (
	ModuleName    = "Traefik"
	ModuleVersion = "1.0.1"
)

type Config struct {
	ServerSideKey             string  `json:"serverSideKey"`
	EnableGraphQLSupport      *bool   `json:"enableGraphQLSupport,omitempty"`
	EnableReferrerRestoration *bool   `json:"enableReferrerRestoration,omitempty"`
	Endpoint                  *string `json:"endpoint,omitempty"`
	MaximumBodySize           *int    `json:"maximumBodySize,omitempty"`
	Timeout                   *int    `json:"timeout,omitempty"`
	UrlPatternExclusion       *string `json:"urlPatternExclusion,omitempty"`
	UrlPatternInclusion       *string `json:"urlPatternInclusion,omitempty"`
	UseXForwardedHost         *bool   `json:"useXForwardedHost,omitempty"`
}

func CreateConfig() *Config {
	return &Config{}
}

type DataDomePlugin struct {
	next           http.Handler
	name           string
	datadomeClient *modulego.Client
}

func loadOptionsFromConfig(config *Config) []modulego.Option {
	var options []modulego.Option

	if config.EnableGraphQLSupport != nil {
		options = append(options, modulego.WithGraphQLSupport(*config.EnableGraphQLSupport))
	}
	if config.EnableReferrerRestoration != nil {
		options = append(options, modulego.WithReferrerRestoration(*config.EnableReferrerRestoration))
	}
	if config.Endpoint != nil {
		options = append(options, modulego.WithEndpoint(*config.Endpoint))
	}
	if config.MaximumBodySize != nil {
		options = append(options, modulego.WithMaximumBodySize(*config.MaximumBodySize))
	}
	if config.Timeout != nil {
		options = append(options, modulego.WithTimeout(*config.Timeout))
	}
	if config.UrlPatternExclusion != nil {
		options = append(options, modulego.WithUrlPatternExclusion(*config.UrlPatternExclusion))
	}
	if config.UrlPatternInclusion != nil {
		options = append(options, modulego.WithUrlPatternInclusion(*config.UrlPatternInclusion))
	}
	if config.UseXForwardedHost != nil {
		options = append(options, modulego.WithXForwardedHost(*config.UseXForwardedHost))
	}

	return options
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	serverSideKey := config.ServerSideKey
	options := loadOptionsFromConfig(config)

	ddClient, err := modulego.NewClient(
		serverSideKey,
		options...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create DataDome client: %w", err)
	}
	ddClient.ModuleName = ModuleName
	ddClient.ModuleVersion = ModuleVersion

	return &DataDomePlugin{
		next:           next,
		name:           name,
		datadomeClient: ddClient,
	}, nil
}

func (m *DataDomePlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Example of payload, values should come from the req
	isBlocked, err := m.datadomeClient.DatadomeProtect(rw, req)
	if err != nil {
		fmt.Println("error when requesting DataDome", err)
	}

	if isBlocked {
		fmt.Println("request blocked by DataDome")
		return
	}

	m.next.ServeHTTP(rw, req)
}
