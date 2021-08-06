package espay

import (
	"reflect"
	"testing"
)

func Test_CreateVA_InvalidURL(t *testing.T) {
	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
		Pretty:          true,
	}

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
			name: "Test Create VA - Invalid URL",
			fields: fields{
				Client: EspayClient{
					BaseUrl:      "https://espay",
					SignatureKey: "",
					Timeout:      0,
					IsProduction: false,
					Logger:       *NewLogger(logOption),
				},
			},
			args: args{
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

func Test_CreateVA_Success(t *testing.T) {
	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
		Pretty:          true,
	}

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
			name: "Test Create VA - Success",
			fields: fields{
				Client: EspayClient{
					BaseUrl:      "https://geni.ktbs.xyz/sachie/api",
					SignatureKey: "",
					Timeout:      0,
					IsProduction: false,
					Logger:       *NewLogger(logOption),
				},
			},
			args: args{
				req: CreateVaRequest{
					Amount:       "10000",
					OrderId:      "12335",
					MerchantCode: "test",
				},
			},
			wantRes: CreateVaResponse{
				RequestUUID:     "espay88",
				ErrorCode:       "00",
				ErrorMessage:    "Success",
				VaNumber:        "1248391000000104",
				RequestDateTime: "2021-07-22 16:30:55",
				BankCode:        "002",
				Amount:          "100000",
				Fee:             "0.00",
			},
			wantErr: false,
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
			if reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CreateVA() gotRes = %v, want %v", gotRes.ErrorCode, tt.wantRes.ErrorCode)
			}
		})
	}
}
