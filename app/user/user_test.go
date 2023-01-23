package user

import (
	"reflect"
	"testing"
)

func TestGenerateUser(t *testing.T) {
	type args struct {
		id   int
		name string
		age  int
	}
	tests := []struct {
		name  string
		args  args
		wantU User
	}{
		{
			name: "Success 1",
			args: args{
				id:   1,
				name: "hoge",
				age:  10,
			},
			wantU: User{
				Id:   1,
				Name: "hoge",
				Age:  10,
			},
		},
		{
			name: "Success 2",
			args: args{
				id:   2,
				name: "fuga",
				age:  100,
			},
			wantU: User{
				Id:   2,
				Name: "fuga",
				Age:  100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotU := GenerateUser(tt.args.id, tt.args.name, tt.args.age); !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("GenerateUser() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}
