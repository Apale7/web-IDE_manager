package rpc

import (
	"context"
	"fmt"
	"testing"
	"web-IDE_manager/proto/docker_manager"
)

var ctx = context.Background()

func TestGetContainers(t *testing.T) {
	type args struct {
		ctx         context.Context
		userID      uint
		containerID string
	}
	tests := []struct {
		name           string
		args           args
		wantContainers []*docker_manager.Container
		wantErr        bool
	}{
		{
			name:    "test1",
			args:    args{ctx: ctx, userID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContainers, err := GetContainers(tt.args.ctx, tt.args.userID, tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContainers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotContainers)
			// if !reflect.DeepEqual(gotContainers, tt.wantContainers) {
			// 	t.Errorf("GetContainers() = %v, want %v", gotContainers, tt.wantContainers)
			// }
		})
	}
}
