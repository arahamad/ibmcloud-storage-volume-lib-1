// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/metrics"
)

type Factory struct {
	GetSummaryMetricStub        func(properties metrics.MetricProperties) (metrics.SummaryMetric, error)
	getSummaryMetricMutex       sync.RWMutex
	getSummaryMetricArgsForCall []struct {
		properties metrics.MetricProperties
	}
	getSummaryMetricReturns struct {
		result1 metrics.SummaryMetric
		result2 error
	}
	getSummaryMetricReturnsOnCall map[int]struct {
		result1 metrics.SummaryMetric
		result2 error
	}
	GetGaugeMetricStub        func(properties metrics.MetricProperties) (metrics.GaugeMetric, error)
	getGaugeMetricMutex       sync.RWMutex
	getGaugeMetricArgsForCall []struct {
		properties metrics.MetricProperties
	}
	getGaugeMetricReturns struct {
		result1 metrics.GaugeMetric
		result2 error
	}
	getGaugeMetricReturnsOnCall map[int]struct {
		result1 metrics.GaugeMetric
		result2 error
	}
	GetHistogramMetricStub        func(properties metrics.MetricProperties, buckets []float64) (metrics.HistogramMetric, error)
	getHistogramMetricMutex       sync.RWMutex
	getHistogramMetricArgsForCall []struct {
		properties metrics.MetricProperties
		buckets    []float64
	}
	getHistogramMetricReturns struct {
		result1 metrics.HistogramMetric
		result2 error
	}
	getHistogramMetricReturnsOnCall map[int]struct {
		result1 metrics.HistogramMetric
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Factory) GetSummaryMetric(properties metrics.MetricProperties) (metrics.SummaryMetric, error) {
	fake.getSummaryMetricMutex.Lock()
	ret, specificReturn := fake.getSummaryMetricReturnsOnCall[len(fake.getSummaryMetricArgsForCall)]
	fake.getSummaryMetricArgsForCall = append(fake.getSummaryMetricArgsForCall, struct {
		properties metrics.MetricProperties
	}{properties})
	fake.recordInvocation("GetSummaryMetric", []interface{}{properties})
	fake.getSummaryMetricMutex.Unlock()
	if fake.GetSummaryMetricStub != nil {
		return fake.GetSummaryMetricStub(properties)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSummaryMetricReturns.result1, fake.getSummaryMetricReturns.result2
}

func (fake *Factory) GetSummaryMetricCallCount() int {
	fake.getSummaryMetricMutex.RLock()
	defer fake.getSummaryMetricMutex.RUnlock()
	return len(fake.getSummaryMetricArgsForCall)
}

func (fake *Factory) GetSummaryMetricArgsForCall(i int) metrics.MetricProperties {
	fake.getSummaryMetricMutex.RLock()
	defer fake.getSummaryMetricMutex.RUnlock()
	return fake.getSummaryMetricArgsForCall[i].properties
}

func (fake *Factory) GetSummaryMetricReturns(result1 metrics.SummaryMetric, result2 error) {
	fake.GetSummaryMetricStub = nil
	fake.getSummaryMetricReturns = struct {
		result1 metrics.SummaryMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) GetSummaryMetricReturnsOnCall(i int, result1 metrics.SummaryMetric, result2 error) {
	fake.GetSummaryMetricStub = nil
	if fake.getSummaryMetricReturnsOnCall == nil {
		fake.getSummaryMetricReturnsOnCall = make(map[int]struct {
			result1 metrics.SummaryMetric
			result2 error
		})
	}
	fake.getSummaryMetricReturnsOnCall[i] = struct {
		result1 metrics.SummaryMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) GetGaugeMetric(properties metrics.MetricProperties) (metrics.GaugeMetric, error) {
	fake.getGaugeMetricMutex.Lock()
	ret, specificReturn := fake.getGaugeMetricReturnsOnCall[len(fake.getGaugeMetricArgsForCall)]
	fake.getGaugeMetricArgsForCall = append(fake.getGaugeMetricArgsForCall, struct {
		properties metrics.MetricProperties
	}{properties})
	fake.recordInvocation("GetGaugeMetric", []interface{}{properties})
	fake.getGaugeMetricMutex.Unlock()
	if fake.GetGaugeMetricStub != nil {
		return fake.GetGaugeMetricStub(properties)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getGaugeMetricReturns.result1, fake.getGaugeMetricReturns.result2
}

func (fake *Factory) GetGaugeMetricCallCount() int {
	fake.getGaugeMetricMutex.RLock()
	defer fake.getGaugeMetricMutex.RUnlock()
	return len(fake.getGaugeMetricArgsForCall)
}

func (fake *Factory) GetGaugeMetricArgsForCall(i int) metrics.MetricProperties {
	fake.getGaugeMetricMutex.RLock()
	defer fake.getGaugeMetricMutex.RUnlock()
	return fake.getGaugeMetricArgsForCall[i].properties
}

func (fake *Factory) GetGaugeMetricReturns(result1 metrics.GaugeMetric, result2 error) {
	fake.GetGaugeMetricStub = nil
	fake.getGaugeMetricReturns = struct {
		result1 metrics.GaugeMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) GetGaugeMetricReturnsOnCall(i int, result1 metrics.GaugeMetric, result2 error) {
	fake.GetGaugeMetricStub = nil
	if fake.getGaugeMetricReturnsOnCall == nil {
		fake.getGaugeMetricReturnsOnCall = make(map[int]struct {
			result1 metrics.GaugeMetric
			result2 error
		})
	}
	fake.getGaugeMetricReturnsOnCall[i] = struct {
		result1 metrics.GaugeMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) GetHistogramMetric(properties metrics.MetricProperties, buckets []float64) (metrics.HistogramMetric, error) {
	var bucketsCopy []float64
	if buckets != nil {
		bucketsCopy = make([]float64, len(buckets))
		copy(bucketsCopy, buckets)
	}
	fake.getHistogramMetricMutex.Lock()
	ret, specificReturn := fake.getHistogramMetricReturnsOnCall[len(fake.getHistogramMetricArgsForCall)]
	fake.getHistogramMetricArgsForCall = append(fake.getHistogramMetricArgsForCall, struct {
		properties metrics.MetricProperties
		buckets    []float64
	}{properties, bucketsCopy})
	fake.recordInvocation("GetHistogramMetric", []interface{}{properties, bucketsCopy})
	fake.getHistogramMetricMutex.Unlock()
	if fake.GetHistogramMetricStub != nil {
		return fake.GetHistogramMetricStub(properties, buckets)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getHistogramMetricReturns.result1, fake.getHistogramMetricReturns.result2
}

func (fake *Factory) GetHistogramMetricCallCount() int {
	fake.getHistogramMetricMutex.RLock()
	defer fake.getHistogramMetricMutex.RUnlock()
	return len(fake.getHistogramMetricArgsForCall)
}

func (fake *Factory) GetHistogramMetricArgsForCall(i int) (metrics.MetricProperties, []float64) {
	fake.getHistogramMetricMutex.RLock()
	defer fake.getHistogramMetricMutex.RUnlock()
	return fake.getHistogramMetricArgsForCall[i].properties, fake.getHistogramMetricArgsForCall[i].buckets
}

func (fake *Factory) GetHistogramMetricReturns(result1 metrics.HistogramMetric, result2 error) {
	fake.GetHistogramMetricStub = nil
	fake.getHistogramMetricReturns = struct {
		result1 metrics.HistogramMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) GetHistogramMetricReturnsOnCall(i int, result1 metrics.HistogramMetric, result2 error) {
	fake.GetHistogramMetricStub = nil
	if fake.getHistogramMetricReturnsOnCall == nil {
		fake.getHistogramMetricReturnsOnCall = make(map[int]struct {
			result1 metrics.HistogramMetric
			result2 error
		})
	}
	fake.getHistogramMetricReturnsOnCall[i] = struct {
		result1 metrics.HistogramMetric
		result2 error
	}{result1, result2}
}

func (fake *Factory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSummaryMetricMutex.RLock()
	defer fake.getSummaryMetricMutex.RUnlock()
	fake.getGaugeMetricMutex.RLock()
	defer fake.getGaugeMetricMutex.RUnlock()
	fake.getHistogramMetricMutex.RLock()
	defer fake.getHistogramMetricMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Factory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.Factory = new(Factory)