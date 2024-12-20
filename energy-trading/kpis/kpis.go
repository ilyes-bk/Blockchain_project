package kpis

import (
	"time"
)

// Struct for KPI results
type KPIMetrics struct {
	TransactionThroughput float64
	Latency               float64
	Scalability           string
	Security              string
	CostEfficiency        string
}

// MeasureTransactionThroughput calculates transactions per second
func MeasureTransactionThroughput(totalTransactions int, duration time.Duration) float64 {
	return float64(totalTransactions) / duration.Seconds()
}

// MeasureLatency calculates average latency
func MeasureLatency(startTime time.Time, endTime time.Time) float64 {
	return endTime.Sub(startTime).Seconds()
}

// AssessScalability simulates scalability assessment
func AssessScalability(loadLevel int) string {
	if loadLevel < 1000 {
		return "High scalability"
	} else if loadLevel < 5000 {
		return "Moderate scalability"
	} else {
		return "Low scalability"
	}
}

// EvaluateSecurity checks basic conditions (simulated)
func EvaluateSecurity(isDataSecure bool) string {
	if isDataSecure {
		return "High security"
	}
	return "Low security"
}

// AnalyzeCostEfficiency simulates cost analysis
func AnalyzeCostEfficiency(cost float64, value float64) string {
	if value/cost > 1 {
		return "Cost-efficient"
	}
	return "Not cost-efficient"
}
