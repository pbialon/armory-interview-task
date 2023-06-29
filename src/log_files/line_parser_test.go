package log_files

import "testing"

func Test_logLineImpl_Timestamp(t *testing.T) {
	type fields struct {
		line string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			name: "should parse datetime",
			fields: fields{
				line: "2021-01-01T00:00:00Z,log1 line1",
			},
			want:    1609459200,
			wantErr: false,
		}, {
			name: "should return error if datetime is invalid",
			fields: fields{
				line: "2021-13-01T00:00:00,log1 line1",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logLineImpl{
				line: tt.fields.line,
			}
			got, err := l.Timestamp()
			if (err != nil) != tt.wantErr {
				t.Errorf("Timestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Timestamp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
