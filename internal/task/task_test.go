package task_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/atc/internal/task"
)

func TestGetExamples(t *testing.T) {
	type args struct {
		contest  string
		question string
	}
	tests := []struct {
		name string
		args args
		want task.Examples
	}{
		// TODO できるようにする
		// {"Success", args{"abc001", "a"}, &task.Example{"15\n10", "5", "0\n0", "0", "5 20", "-15"}, false},
		// TODO できるようにする
		// {"Success", args{"abc020", "a"}, &task.Example{"1", "ABC", "2", "chokudai", "", ""}, false},
		{"ABC042", args{"abc042", "a"}, task.Examples{&task.Example{"5 5 7\n", "YES\n"}, &task.Example{"7 7 5\n", "NO\n"}}},
		{"ABC145", args{"abc145", "f"}, task.Examples{&task.Example{"4 1\n2 3 4 1\n", "3\n"}, &task.Example{"6 2\n8 6 9 1 2 1\n", "7\n"}, &task.Example{"10 0\n1 1000000000 1 1000000000 1 1000000000 1 1000000000 1 1000000000\n", "4999999996\n"}}},
		{"ARC072", args{"arc072", "b"}, task.Examples{&task.Example{"2 1\n", "Brown\n"}, &task.Example{"5 0\n", "Alice\n"}, &task.Example{"0 0\n", "Brown\n"}, &task.Example{"4 8\n", "Alice\n"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := task.GetExamples(tt.args.contest, tt.args.question)
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
