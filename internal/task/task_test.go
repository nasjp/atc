package task_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/atc/internal/task"
)

func TestGetExample(t *testing.T) {
	type args struct {
		contest  string
		question string
	}
	tests := []struct {
		name string
		args args
		want *task.Example
	}{
		// TODO できるようにする
		// {"Success", args{"abc001", "a"}, &task.Example{"15\n10", "5", "0\n0", "0", "5 20", "-15"}, false},
		// TODO できるようにする
		// {"Success", args{"abc020", "a"}, &task.Example{"1", "ABC", "2", "chokudai", "", ""}, false},
		{"ABC042", args{"abc042", "a"}, &task.Example{"5 5 7\n", "YES\n", "7 7 5\n", "NO\n", "", "", "", ""}},
		{"ABC145", args{"abc145", "f"}, &task.Example{"4 1\n2 3 4 1\n", "3\n", "6 2\n8 6 9 1 2 1\n", "7\n", "10 0\n1 1000000000 1 1000000000 1 1000000000 1 1000000000 1 1000000000\n", "4999999996\n", "", ""}},
		{"ARC072", args{"arc072", "b"}, &task.Example{"2 1\n", "Brown\n", "5 0\n", "Alice\n", "0 0\n", "Brown\n", "4 8\n", "Alice\n"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := task.GetExample(tt.args.contest, tt.args.question)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExamples() = %v, want %v", got, tt.want)
			}
		})
	}
}
