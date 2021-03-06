package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckExistsUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &CheckExistsUserParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",

		Username: "SYSOP2",
	}

	expect0 := &CheckExistsUserResult{IsExists: false}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, params: params0},
			expectedResult:     expect0,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := CheckExistsUser(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckExistsUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("CheckExistsUser() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CheckExistsUser() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
