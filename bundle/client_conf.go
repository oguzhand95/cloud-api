// Copyright 2021-2022 Zenauth Ltd.

package bundle

import (
	"errors"
	"fmt"
	"os"
	"time"

	pdpv1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/pdp/v1"
	"github.com/go-logr/logr"
	"go.uber.org/multierr"
)

var (
	errEmptyServerURL    = errors.New("server URL must be defined")
	errMissingAPIKey     = errors.New("missing API key")
	errMissingIdentifier = errors.New("missing PDP identifier")
)

const (
	defaultRetryWaitMin     = 1 * time.Second //nolint:revive
	defaultRetryWaitMax     = 5 * time.Minute
	defaultRetryMaxAttempts = 10
)

type ClientConf struct {
	PDPIdentifier    *pdpv1.Identifier
	Logger           logr.Logger
	APIKey           string
	ServerURL        string
	CacheDir         string
	TempDir          string
	RetryWaitMin     time.Duration
	RetryWaitMax     time.Duration
	RetryMaxAttempts int
}

func (cc ClientConf) Validate() (outErr error) {
	if cc.APIKey == "" {
		outErr = multierr.Append(outErr, errMissingAPIKey)
	}

	if cc.ServerURL == "" {
		outErr = multierr.Append(outErr, errEmptyServerURL)
	}

	if cc.PDPIdentifier == nil {
		outErr = multierr.Append(outErr, errMissingIdentifier)
	} else if err := cc.PDPIdentifier.ValidateAll(); err != nil {
		outErr = multierr.Append(outErr, fmt.Errorf("invalid PDP identifier: %w", err))
	}

	if cc.CacheDir != "" {
		if err := validateDir(cc.CacheDir); err != nil {
			outErr = multierr.Append(outErr, err)
		}
	}

	if cc.TempDir != "" {
		if err := validateDir(cc.TempDir); err != nil {
			outErr = multierr.Append(outErr, err)
		}
	}

	return outErr
}

func (cc *ClientConf) SetDefaults() {
	if cc.RetryMaxAttempts == 0 {
		cc.RetryMaxAttempts = defaultRetryMaxAttempts
	}

	if cc.RetryWaitMin == 0 {
		cc.RetryWaitMin = defaultRetryWaitMin
	}

	if cc.RetryWaitMax == 0 {
		cc.RetryWaitMax = defaultRetryWaitMax
	}
}

func validateDir(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat %q: %w", path, err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("not a directory: %q", path)
	}

	return nil
}

type logWrapper struct {
	logr.Logger
}

func (lw logWrapper) Printf(msg string, args ...any) {
	if log := lw.V(1); log.Enabled() {
		log.Info(fmt.Sprintf(msg, args...))
	}
}