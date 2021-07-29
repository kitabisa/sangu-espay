package espay

import (
	"reflect"
	"testing"
)

func TestCoreGateway_CreateVA(t *testing.T) {
	type fields struct {
		Client EspayClient
	}
	type args struct {
		req CreateVaRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes CreateVaResponse
		wantErr bool
	}{
		{
			name:    "Test invalid url",
			fields:  fields{
				Client: EspayClient{
					BaseUrl:      "https://espay",
					SignatureKey: "",
					Timeout:      0,
					IsProduction: false,
				},
			},
			args:    args{
				req: CreateVaRequest{},
			},
			wantRes: CreateVaResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := tt.fields.Client
			gotRes, err := gateway.CreateVA(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CreateVA() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
