package main

import (
	"bytes"
	"encoding/json"
)

type input struct {
	Name      string `json:"name"`
	Submitted int64  `json:"submitted"`
}

func newInputFromJSON(b []byte) *input {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat input
	dec.Decode(&pstat)
	return &pstat
}

func (i *input) toPoints() []*point {
	points := make([]*point, 1)

	points[0] = &point{
		Name:        "input_messages_submitted_total",
		Type:        counter,
		Value:       i.Submitted,
		LabelNames:  []string{"name"},
		LabelValues: []string{i.Name},
		Description: "messages submitted",
	}

	return points
}
