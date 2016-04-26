package config

import "errors"

// Errors
var (
	// Error
	ErrInterval                        = errors.New("interval should be an integer between 1s~10min")
	ErrPeriod                          = errors.New("period should be an integer greater than interval")
	ErrPeriodTooSmall                  = errors.New("period at least 1 hour")
	ErrExpiration                      = errors.New("expiration should be an integer greater than 5 * period")
	ErrExpirationDivPeriodClean        = errors.New("expiration should be divided by period cleanly")
	ErrDetectorPort                    = errors.New("invalid detector.port")
	ErrDetectorTrendingFactor          = errors.New("detector.trendingFactor should be a float between 0 and 1")
	ErrDetectorFilterOffset            = errors.New("detector.filterOffset should be a float between 0 and 1")
	ErrDetectorFilterTimes             = errors.New("detector.filterTimes should be smaller")
	ErrDetectorDefaultThresholdMaxsLen = errors.New("detector.defaultThresholdMaxs should have up to 8 items")
	ErrDetectorDefaultThresholdMinsLen = errors.New("detector.defaultThresholdMins should have up to 8 items")
	ErrDetectorDefaultThresholdMaxZero = errors.New("detector.defaultThresholdMaxs should not contain zeros")
	ErrDetectorDefaultThresholdMinZero = errors.New("detector.defaultThresholdMins should not contain zeros")
	ErrDetectorFillBlankZerosLen       = errors.New("detector.fillBlankZeros should have up to 8 items")
	ErrWebappPort                      = errors.New("invalid webapp.port")
	ErrWebappLanguage                  = errors.New("invalid webapp language")
	ErrAlerterInterval                 = errors.New("alerter.interval should be greater than 0")
	ErrAlerterOneDayLimit              = errors.New("alerter.oneDayLimit should be greater than 0")
	ErrAlerterDefaultSilentTimeRange   = errors.New("alerter.defaultTimeRange should be 2 numbers between 0~24")
	// Warn
	ErrAlerterCommandEmpty = errors.New("alerter.command is empty")
)
