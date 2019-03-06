// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.ibm.com/narkarum/ibmcloud-storage-volume-lib/volume-providers/vpc/metrics"
)

type GaugeMetric struct {
	RegisterMetricStub        func(registerer prometheus.Registerer)
	registerMetricMutex       sync.RWMutex
	registerMetricArgsForCall []struct {
		registerer prometheus.Registerer
	}
	AddStub        func(i float64, labels metrics.MetricLabels)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		i      float64
		labels metrics.MetricLabels
	}
	SubStub        func(i float64, labels metrics.MetricLabels)
	subMutex       sync.RWMutex
	subArgsForCall []struct {
		i      float64
		labels metrics.MetricLabels
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *GaugeMetric) RegisterMetric(registerer prometheus.Registerer) {
	fake.registerMetricMutex.Lock()
	fake.registerMetricArgsForCall = append(fake.registerMetricArgsForCall, struct {
		registerer prometheus.Registerer
	}{registerer})
	fake.recordInvocation("RegisterMetric", []interface{}{registerer})
	fake.registerMetricMutex.Unlock()
	if fake.RegisterMetricStub != nil {
		fake.RegisterMetricStub(registerer)
	}
}

func (fake *GaugeMetric) RegisterMetricCallCount() int {
	fake.registerMetricMutex.RLock()
	defer fake.registerMetricMutex.RUnlock()
	return len(fake.registerMetricArgsForCall)
}

func (fake *GaugeMetric) RegisterMetricArgsForCall(i int) prometheus.Registerer {
	fake.registerMetricMutex.RLock()
	defer fake.registerMetricMutex.RUnlock()
	return fake.registerMetricArgsForCall[i].registerer
}

func (fake *GaugeMetric) Add(i float64, labels metrics.MetricLabels) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		i      float64
		labels metrics.MetricLabels
	}{i, labels})
	fake.recordInvocation("Add", []interface{}{i, labels})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(i, labels)
	}
}

func (fake *GaugeMetric) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *GaugeMetric) AddArgsForCall(i int) (float64, metrics.MetricLabels) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].i, fake.addArgsForCall[i].labels
}

func (fake *GaugeMetric) Sub(i float64, labels metrics.MetricLabels) {
	fake.subMutex.Lock()
	fake.subArgsForCall = append(fake.subArgsForCall, struct {
		i      float64
		labels metrics.MetricLabels
	}{i, labels})
	fake.recordInvocation("Sub", []interface{}{i, labels})
	fake.subMutex.Unlock()
	if fake.SubStub != nil {
		fake.SubStub(i, labels)
	}
}

func (fake *GaugeMetric) SubCallCount() int {
	fake.subMutex.RLock()
	defer fake.subMutex.RUnlock()
	return len(fake.subArgsForCall)
}

func (fake *GaugeMetric) SubArgsForCall(i int) (float64, metrics.MetricLabels) {
	fake.subMutex.RLock()
	defer fake.subMutex.RUnlock()
	return fake.subArgsForCall[i].i, fake.subArgsForCall[i].labels
}

func (fake *GaugeMetric) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMetricMutex.RLock()
	defer fake.registerMetricMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.subMutex.RLock()
	defer fake.subMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *GaugeMetric) recordInvocation(key string, args []interface{}) {
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

var _ metrics.GaugeMetric = new(GaugeMetric)