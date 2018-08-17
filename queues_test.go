package main

import (
	"reflect"
	"testing"
)

var (
	queueLog = []byte(`{"name":"main Q","size":10,"enqueued":20,"full":30,"discarded.full":40,"discarded.nf":50,"maxqsize":60}`)
)

func TestNewQueueFromJSON(t *testing.T) {
	logType := getStatType(queueLog)
	if logType != rsyslogQueue {
		t.Errorf("detected pstat type should be %d but is %d", rsyslogQueue, logType)
	}

	pstat := newQueueFromJSON([]byte(queueLog))

	if want, got := "main_q", pstat.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(10), pstat.Size; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(20), pstat.Enqueued; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(30), pstat.Full; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(40), pstat.DiscardedFull; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(50), pstat.DiscardedNf; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(60), pstat.MaxQsize; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestQueueToPoints(t *testing.T) {
	pstat := newQueueFromJSON([]byte(queueLog))
	points := pstat.toPoints()

	point := points[0]
	if want, got := "size", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(10), point.Value; want != got {
	}

	if want, got := gauge, point.Type; want != got {
	}

	point = points[1]
	if want, got := "enqueued_total", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(20), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := []string{"name"}, point.LabelNames; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	if want, got := []string{"main_q"}, point.LabelValues; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	point = points[2]
	if want, got := "full_total", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(30), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := []string{"name"}, point.LabelNames; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	if want, got := []string{"main_q"}, point.LabelValues; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	point = points[3]
	if want, got := "discarded_full_total", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(40), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := []string{"name"}, point.LabelNames; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	if want, got := []string{"main_q"}, point.LabelValues; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	point = points[4]
	if want, got := "discarded_not_full_total", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(50), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := []string{"name"}, point.LabelNames; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	if want, got := []string{"main_q"}, point.LabelValues; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	point = points[5]
	if want, got := "max_queue_size", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(60), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := gauge, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := []string{"name"}, point.LabelNames; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

	if want, got := []string{"main_q"}, point.LabelValues; !reflect.DeepEqual(want, got) {
		t.Errorf("wanted '%v', got '%v'", want, got)
	}

}
