package main

import (
	"bytes"
	"encoding/json"
)

type resource struct {
	Name     string `json:"name"`
	Utime    int64  `json:"utime"`
	Stime    int64  `json:"stime"`
	Maxrss   int64  `json:"maxrss"`
	Minflt   int64  `json:"minflt"`
	Majflt   int64  `json:"majflt"`
	Inblock  int64  `json:"inblock"`
	Outblock int64  `json:"oublock"`
	Nvcsw    int64  `json:"nvcsw"`
	Nivcsw   int64  `json:"nivcsw"`
}

func newResourceFromJSON(b []byte) *resource {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat resource
	dec.Decode(&pstat)
	return &pstat
}

func (r *resource) toPoints() []*point {
	points := make([]*point, 9)

	points[0] = &point{
		Name:        "resource_usage_utime_microseconds_total",
		Type:        counter,
		Value:       r.Utime,
		Description: "user time used in microseconds",
	}

	points[1] = &point{
		Name:        "resource_usage_stime_microseconds_total",
		Type:        counter,
		Value:       r.Stime,
		Description: "system time used in microsends",
	}

	points[2] = &point{
		Name:        "resource_usage_maxrss_bytes",
		Type:        gauge,
		Value:       r.Maxrss,
		Description: "maximum resident set size",
	}

	points[3] = &point{
		Name:        "resource_usage_minflt_total",
		Type:        counter,
		Value:       r.Minflt,
		Description: "total minor faults",
	}

	points[4] = &point{
		Name:        "resource_usage_majflt_total",
		Type:        counter,
		Value:       r.Majflt,
		Description: "total major faults",
	}

	points[5] = &point{
		Name:        "resource_usage_inblock_total",
		Type:        counter,
		Value:       r.Inblock,
		Description: "filesystem input operations",
	}

	points[6] = &point{
		Name:        "resource_usage_oublock_total",
		Type:        counter,
		Value:       r.Outblock,
		Description: "filesystem output operations",
	}

	points[7] = &point{
		Name:        "resource_usage_nvcsw_total",
		Type:        counter,
		Value:       r.Nvcsw,
		Description: "voluntary context switches",
	}

	points[8] = &point{
		Name:        "resource_usage_nivcsw_total",
		Type:        counter,
		Value:       r.Nivcsw,
		Description: "involuntary context switches",
	}

	return points
}
