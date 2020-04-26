package csv_conv

import (
	"reflect"
	"testing"
)

func TestConverter_AppendColumn(t *testing.T) {
	type fields struct {
		original [][]string
	}
	type args struct {
		colName      string
		defaultValue string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				original: [][]string{
					{"Rob", "Pike", "rob"},
					{"Ken", "Thompson", "ken"},
					{"Robert", "Griesemer", "gri"},
				},
			},
			args: args{
				colName:      "Steve",
				defaultValue: "",
			},
			want: [][]string{
				{"Rob", "Pike", "rob", "Steve"},
				{"Ken", "Thompson", "ken", ""},
				{"Robert", "Griesemer", "gri", ""},
			},
		},
		{
			name:    "fail empty or nil source",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				original: tt.fields.original,
			}
			got, err := c.AppendColumn(tt.args.colName, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendColumn() got = %v, want %v", got, tt.want)
			}
		})
	}
}
