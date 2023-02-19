package userservice_test

import (
	"context"
	"reflect"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func TestUserGetFeedService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *user.FeedRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.FeedResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userservice.UserGetFeedService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserGetFeedService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserGetFeedService() = %v, want %v", got, tt.want)
			}
		})
	}
}