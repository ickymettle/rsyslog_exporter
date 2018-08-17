package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

type action struct {
	Name              string `json:"name"`
	Processed         int64  `json:"processed"`
	Failed            int64  `json:"failed"`
	Suspended         int64  `json:"suspended"`
	SuspendedDuration int64  `json:"suspended.duration"`
	Resumed           int64  `json:"resumed"`
}

func newActionFromJSON(b []byte) *action {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat action
	dec.Decode(&pstat)
	pstat.Name = strings.ToLower(pstat.Name)
	pstat.Name = strings.Replace(pstat.Name, " ", "_", -1)
	return &pstat
}

func (a *action) toPoints() []*point {
	points := make([]*point, 5)

	points[0] = &point{
		Name:        "processed_total",
		Type:        counter,
		Value:       a.Processed,
		Description: "messages processed",
		LabelNames:  []string{"name"},
		LabelValues: []string{a.Name},
	}

	points[1] = &point{
		Name:        "failed_total",
		Type:        counter,
		Value:       a.Failed,
		Description: "messages failed",
		LabelNames:  []string{"name"},
		LabelValues: []string{a.Name},
	}

	points[2] = &point{
		Name:        "suspended_total",
		Type:        counter,
		Value:       a.Suspended,
		Description: "times suspended",
		LabelNames:  []string{"name"},
		LabelValues: []string{a.Name},
	}

	points[3] = &point{
		Name:        "suspended_duration_seconds",
		Type:        counter,
		Value:       a.SuspendedDuration,
		Description: "time spent suspended",
		LabelNames:  []string{"name"},
		LabelValues: []string{a.Name},
	}

	points[4] = &point{
		Name:        "resumed_total",
		Type:        counter,
		Value:       a.Resumed,
		Description: "times resumed",
		LabelNames:  []string{"name"},
		LabelValues: []string{a.Name},
	}

	return points
}
