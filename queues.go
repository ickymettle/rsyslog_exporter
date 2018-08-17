package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

type queue struct {
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Enqueued      int64  `json:"enqueued"`
	Full          int64  `json:"full"`
	DiscardedFull int64  `json:"discarded.full"`
	DiscardedNf   int64  `json:"discarded.nf"`
	MaxQsize      int64  `json:"maxqsize"`
}

func newQueueFromJSON(b []byte) *queue {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat queue
	dec.Decode(&pstat)
	pstat.Name = strings.ToLower(pstat.Name)
	pstat.Name = strings.Replace(pstat.Name, " ", "_", -1)
	return &pstat
}

func (q *queue) toPoints() []*point {
	points := make([]*point, 6)

	points[0] = &point{
		Name:        "size",
		Type:        gauge,
		Value:       q.Size,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "messages currently in queue",
	}

	points[1] = &point{
		Name:        "enqueued_total",
		Type:        counter,
		Value:       q.Enqueued,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "total messages enqueued",
	}

	points[2] = &point{
		Name:        "full_total",
		Type:        counter,
		Value:       q.Full,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "times queue was full",
	}

	points[3] = &point{
		Name:        "discarded_full_total",
		Type:        counter,
		Value:       q.DiscardedFull,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "messages discarded due to queue being full",
	}

	points[4] = &point{
		Name:        "discarded_not_full_total",
		Type:        counter,
		Value:       q.DiscardedNf,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "messages discarded when queue not full",
	}

	points[5] = &point{
		Name:        "max_queue_size",
		Type:        gauge,
		Value:       q.MaxQsize,
		LabelNames:  []string{"name"},
		LabelValues: []string{q.Name},
		Description: "maximum size queue has reached",
	}

	return points
}
