package main

import (
	"os"
	"path"
	"testing"
)

func prepareTestFiles(t *testing.T) string {
	tempdir := t.TempDir()
	filename1 := "test1.log"
	filename2 := "test2.log"

	tempfile1, _ := os.Create(path.Join(tempdir, filename1))
	tempfile2, _ := os.Create(path.Join(tempdir, filename2))

	_, _ = tempfile1.Write([]byte("log1 line1\nlog1 line2"))
	_, _ = tempfile2.Write([]byte("log2 line2\n"))

	tempfile1.Close()
	tempfile2.Close()

	return tempdir
}

func TestLocalDiskFilePoolHandler_NextLine(t *testing.T) {

	dir := prepareTestFiles(t)

	type fields struct {
		files map[string]*os.File
	}
	type args struct {
		fileNames []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr []bool
	}{
		{
			name: "should read lines from files",
			args: args{
				fileNames: []string{"test1.log", "test2.log"},
			},
			want:    []string{"log1 line1\n", "log2 line1\n", "log1 line2\n"},
			wantErr: []bool{false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := NewLocalDiskFilePoolHandler(dir)
			for i, filename := range tt.args.fileNames {
				got, err := fp.NextLine(filename)
				if (err != nil) != tt.wantErr[i] {
					t.Errorf("NextLine() error = %v, wantErr %v", err, tt.wantErr[i])
					return
				}
				if got != tt.want[i] {
					t.Errorf("NextLine() got = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}