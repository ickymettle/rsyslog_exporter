package main

import (
	"errors"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type pointType int

const (
	counter pointType = iota
	gauge
)

var (
	ErrIncompatiblePointType = errors.New("incompatible point type")
	ErrUnknownPointType      = errors.New("unknown point type")
)

type point struct {
	Name        string
	Description string
	Type        pointType
	Value       int64
	LabelNames  []string
	LabelValues []string
}

func (p *point) add(newPoint *point) error {
	switch newPoint.Type {
	case gauge:
		if p.Type != gauge {
			return ErrIncompatiblePointType
		}
		p.Value = newPoint.Value
	case counter:
		if p.Type != counter {
			return ErrIncompatiblePointType
		}
		p.Value = p.Value + newPoint.Value
	default:
		return ErrUnknownPointType
	}
	return nil
}

func (p *point) promDescription() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName("", "rsyslog", p.Name),
		p.Description,
		p.LabelNames, nil,
	)
}

func (p *point) promType() prometheus.ValueType {
	if p.Type == counter {
		return prometheus.CounterValue
	}
	return prometheus.GaugeValue
}

func (p *point) promValue() float64 {
	return float64(p.Value)
}

func (p *point) key() string {
	return p.Name + ":" + strings.Join(p.LabelValues, ",")
}
