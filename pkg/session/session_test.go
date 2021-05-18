package session

import (
	"fdk-extension-golang/pkg/models"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		id    string
		isNew bool
	}
	tests := []struct {
		name string
		args args
		want *Session
	}{
		{
			name: "TestNew Session", args: args{"e8e4a6b2ded6a5252dd6ddd1182b0fe8e16e77cc28736a5df3613ae330cc06ad", true}, want: &Session{ID: "e8e4a6b2ded6a5252dd6ddd1182b0fe8e16e77cc28736a5df3613ae330cc06ad", IsNew: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.id, tt.args.isNew); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_CloneSession(t *testing.T) {
	type fields struct {
		ID           string
		CompanyID    string
		State        string
		Scope        []string
		Expires      string
		ExpiresIn    int
		AccessMode   string
		AccessToken  string
		CurrentUser  interface{}
		RefreshToken string
		IsNew        bool
	}
	type args struct {
		id    string
		isNew bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Session
	}{
		{name: "TestSession_CloneSession", fields: fields{ID: "e8e4a6b2ded6a5252dd6ddd1182b0fe8e16e77cc28736a5df3613ae330cc06ad", IsNew: true, CompanyID: "1"}, args: args{id: "e8e4a6b2ded6a5252dd6ddd1182b0fe8e16e77cc28736a5df3613ae330cc06ad", isNew: false}, want: Session{ID: "e8e4a6b2ded6a5252dd6ddd1182b0fe8e16e77cc28736a5df3613ae330cc06ad", CompanyID: "1", IsNew: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				ID:           tt.fields.ID,
				CompanyID:    tt.fields.CompanyID,
				State:        tt.fields.State,
				Scope:        tt.fields.Scope,
				Expires:      tt.fields.Expires,
				ExpiresIn:    tt.fields.ExpiresIn,
				AccessMode:   tt.fields.AccessMode,
				AccessToken:  tt.fields.AccessToken,
				CurrentUser:  tt.fields.CurrentUser,
				RefreshToken: tt.fields.RefreshToken,
				IsNew:        tt.fields.IsNew,
			}
			if got := s.CloneSession(tt.args.id, tt.args.isNew); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.CloneSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSessionID(t *testing.T) {
	type args struct {
		isOnline bool
		options  models.Option
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestGenerateSessionID", args: args{false, models.Option{CompanyID: "1", Cluster: "https://nice-turtle-9.loca.lt"}}, want: "91e0ae01f629dee5ea12674dcd9721ac2e445fc709c405e312d6b74ea034b655", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSessionID(tt.args.isOnline, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSessionID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateSessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}
