package csv_conv

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestConverter_ChangeColumnName(t *testing.T) {
	type fields struct {
		original [][]string
	}
	type args struct {
		newNames map[string]string
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
				newNames: map[string]string{
					"Rob": "Bob",
					"rob": "Steve",
				},
			},
			want: [][]string{
				{"Bob", "Pike", "Steve"},
				{"Ken", "Thompson", "ken"},
				{"Robert", "Griesemer", "gri"},
			},
			wantErr: false,
		},
		{
			name:    "fail nil source",
			fields:  fields{original: nil},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fail empty source",
			fields:  fields{original: [][]string{}},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				original: tt.fields.original,
			}
			got, err := c.ChangeColumnName(tt.args.newNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeColumnName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeColumnName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConverter(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Converter
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				strings.NewReader(`
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`),
			},
			want: &Converter{original: [][]string{
				{"Rob", "Pike", "rob"},
				{"Ken", "Thompson", "ken"},
				{"Robert", "Griesemer", "gri"},
			}},
			wantErr: false,
		},
		{
			name:    "fail get an empty source error",
			args:    args{},
			wantErr: true,
		},
		{
			name: "faild get an invalid source error",
			args: args{
				strings.NewReader(`
"Rob","Pike",rob
Ken,Thompson
"Robert","Griesemer","gri"
`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConverter(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConverter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
