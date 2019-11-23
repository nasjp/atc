package internal_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/atc/internal"
)

const testFile = "tmp.py"

func TestGetExamples(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want internal.Examples
	}{
		// TODO できるようにする
		// {"Success", args{"abc001", "a"}, &internal.Example{"15\n10", "5", "0\n0", "0", "5 20", "-15"}, false},
		// TODO できるようにする
		// {"Success", args{"abc020", "a"}, &internal.Example{"1", "ABC", "2", "chokudai", "", ""}, false},
		{"ABC042", args{"abc042_a.py"}, internal.Examples{&internal.Example{"5 5 7\n", "YES\n"}, &internal.Example{"7 7 5\n", "NO\n"}}},
		{"ABC145", args{"abc145_f.go"}, internal.Examples{&internal.Example{"4 1\n2 3 4 1\n", "3\n"}, &internal.Example{"6 2\n8 6 9 1 2 1\n", "7\n"}, &internal.Example{"10 0\n1 1000000000 1 1000000000 1 1000000000 1 1000000000 1 1000000000\n", "4999999996\n"}}},
		{"ARC072", args{"arc072_b.cpp"}, internal.Examples{&internal.Example{"2 1\n", "Brown\n"}, &internal.Example{"5 0\n", "Alice\n"}, &internal.Example{"0 0\n", "Brown\n"}, &internal.Example{"4 8\n", "Alice\n"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := internal.GetExamples(tt.args.fileName)
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

func setupTestExampleRun(t *testing.T) func(t *testing.T) {

	code := `def f():
    a, b = input().split()
    a, b = int(a), int(b)
    rem = int(a) * int(b) % 2
    if rem == 1:
        print('Odd')
        return
    print('Even')


f()`

	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	f.WriteString(code)

	return func(t *testing.T) {
		if err := os.Remove(testFile); err != nil {
			t.Fatal(err)
		}
	}
}

func TestExample_Run(t *testing.T) {
	defer setupTestExampleRun(t)(t)
	type fields struct {
		Input  string
		Output string
	}
	type args struct {
		command  string
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Success", fields{"3 4", "Even"}, args{"python3", testFile}, "Even\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &internal.Example{
				Input:  tt.fields.Input,
				Output: tt.fields.Output,
			}
			got, err := e.Run(tt.args.command, tt.args.fileName)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Example.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
