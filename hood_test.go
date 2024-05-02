package hood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	A string `confidential:""`
	B int
	C string `confidential:"1,1"`
	D Inner
}

type Inner struct {
	E string `confidential:"20,20"`
	F string `confidential:"0,5"`
}

func TestPrintConfidentialData(t *testing.T) {
	type args struct {
		binding interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				TestStruct{
					A: "FieldA",
					B: 1,
					C: "FieldC",
					D: Inner{
						E: "WordLessThen20",
						F: "012345678901234567890",
					},
				},
			},
			want:    "{A:****** B:1 C:F****C {E:WordLessThen20 F:****************67890}}",
			wantErr: false,
		},
		{
			name: "inner struct only",
			args: args{
				struct {
					A struct {
						B string `confidential:"3,3"`
					}
				}{
					A: struct {
						B string `confidential:"3,3"`
					}{
						B: "ABCDEFG1234567",
					},
				},
			},
			want:    "{{B:ABC********567}}",
			wantErr: false,
		},
		{
			name: "confidential tag on wrong type",
			args: args{
				struct {
					A string
					B int `confidential:"1,1"`
				}{
					"TestOne",
					1,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "confidential tag with wrong head value",
			args: args{
				struct {
					A string `confidential:"n"`
					B int
				}{
					"TestOne",
					1,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "confidential tag with wrong tail value",
			args: args{
				struct {
					A string `confidential:"1,r"`
					B int
				}{
					"TestOne",
					1,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "inner struct have wrong tag parameter",
			args: args{
				struct {
					A string `confidential:"1,1"`
					B struct {
						C string `confidential:"n,r"`
					}
				}{
					A: "TestOne",
					B: struct {
						C string `confidential:"n,r"`
					}{
						C: "TestTwo",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintConfidentialData(tt.args.binding)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintConfidentialData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equalf(t, tt.want, got, "PrintConfidentialData(%v)", tt.args.binding)
		})
	}
}

func TestHoodString(t *testing.T) {
	type args struct {
		target    string
		keepFirst int
		keepTail  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				target:    "TestString",
				keepFirst: 2,
				keepTail:  2,
			},
			want: "Actual   :Te******ng",
		},
		{
			name: "keepFirst > len(target)",
			args: args{
				target:    "TestString",
				keepFirst: 20,
				keepTail:  2,
			},
			want: "TestString",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HoodString(tt.args.target, tt.args.keepFirst, tt.args.keepTail), "HoodString(%v, %v, %v)", tt.args.target, tt.args.keepFirst, tt.args.keepTail)
		})
	}
}
