package dao

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
	"testing"
	"backend-record/pkg/model/dto"

)


func TestMain(m *testing.M){
	err:=Init()
	if err != nil {
		os.Exit(500)
	}
	status := m.Run()
	os.Exit(status)

}

func TestCheckNiceStatus(t *testing.T) {
	type args struct {
		userID    string
		ArticleID string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "CheckNiceStatusSuccess",
			args: args{"b1018085","1"},
			want: true,
			wantErr: false,
		},
		{
			name: "CheckNiceStatusFalse",
			args: args{"b1018085","4"},
			want: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckNiceStatus(tt.args.userID, tt.args.ArticleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckNiceStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckNiceStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upDateAddLike_Request(t *testing.T) {
	type fields struct {
		ArticleID string
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Nice
		wantErr bool
	}{
		{
			name: "updateADDLikeSuccess",
			fields: fields{ArticleID: "2"},
			args: args{"b1018085"},
			want: &dto.Nice{101},
			wantErr: false,
		},
		{
			name: "updateDeleteLikeSuccess",
			fields: fields{ArticleID: "2"},
			args: args{"b1018085"},
			want: &dto.Nice{100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &upDateAddLike{
				ArticleID: tt.fields.ArticleID,
			}
			got, err := info.Request(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request() got = %v, want %v", got, tt.want)
			}
		})
	}
}