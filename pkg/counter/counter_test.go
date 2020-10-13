package counter

import (
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		threads int
	}
	tests := []struct {
		name string
		args args
		want ThreadCounter
	}{
		{"1", args{5}, New(5)},
		{"2", args{7}, New(7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.threads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadCounter_Completed(t *testing.T) {
	type fields struct {
		threads int
		count   int
		mux     *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"1", fields{2, 0, &sync.Mutex{}}},
		{"2", fields{3, 0, &sync.Mutex{}}},
		{"3", fields{7, 0, &sync.Mutex{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &ThreadCounter{
				threads: tt.fields.threads,
				count:   tt.fields.count,
				mux:     tt.fields.mux,
			}
			mc.Completed()
		})
	}
}

func TestThreadCounter_AreCompleted(t *testing.T) {
	type fields struct {
		threads int
		count   int
		mux     *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1", fields{1, 1, &sync.Mutex{}}, true},
		{"2", fields{3, 1, &sync.Mutex{}}, false},
		{"3", fields{7, 0, &sync.Mutex{}}, false},
		{"3", fields{7, 6, &sync.Mutex{}}, true},
		{"3", fields{7, 8, &sync.Mutex{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &ThreadCounter{
				threads: tt.fields.threads,
				count:   tt.fields.count,
				mux:     tt.fields.mux,
			}

			mc.Completed()

			if got := mc.AreCompleted(); got != tt.want {
				t.Errorf("ThreadCounter.AreCompleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
