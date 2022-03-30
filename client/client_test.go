package client

import (
	"log"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var client = resty.New()

func setupTestSuite(tb testing.TB) func(tb testing.TB) {
	httpmock.ActivateNonDefault(client.GetClient())

	return func(tb testing.TB) {
		log.Println("Teardown suite test")
		httpmock.DeactivateAndReset()
	}
}

func Test_exampleClientImp_GetName(t *testing.T) {
	teardownSuite := setupTestSuite(t)
	defer teardownSuite(t)

	type fields struct {
		Client func() *resty.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *ExampleResponse
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "when_normal_should_success",
			fields: fields{
				Client: func() *resty.Client {
					// remove old mock
					httpmock.Reset()

					fakeUrl := "https://example.com/example"
					fixture := &ExampleResponse{Name: "name ja"}
					mockResponder, err := httpmock.NewJsonResponder(200, fixture)
					if err != nil {
						t.Fatalf("fixture is invalid")
					}

					httpmock.RegisterResponder("POST", fakeUrl, mockResponder)

					return client
				},
			},
			args: args{
				id: "1",
			},
			want:    &ExampleResponse{Name: "name ja"},
			wantErr: false,
		},
		{
			name: "when_invalid_model_should_error_unmarshal",
			fields: fields{
				Client: func() *resty.Client {
					// remove old mock
					httpmock.Reset()

					fakeUrl := "https://example.com/example"
					mockResponder, err := httpmock.NewJsonResponder(400, "invalid model")
					if err != nil {
						t.Fatalf("fixture is invalid")
					}

					httpmock.RegisterResponder("POST", fakeUrl, mockResponder)

					return client
				},
			},
			args: args{
				id: "1",
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "json: cannot unmarshal string into Go value of type client.ExampleError",
		},
		{
			name: "when_400_should_error_not_success",
			fields: fields{
				Client: func() *resty.Client {
					// remove old mock
					httpmock.Reset()

					fakeUrl := "https://example.com/example"
					mockResponder, err := httpmock.NewJsonResponder(400, nil)
					if err != nil {
						t.Fatalf("fixture is invalid")
					}

					httpmock.RegisterResponder("POST", fakeUrl, mockResponder)

					return client
				},
			},
			args: args{
				id: "1",
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "Not success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &exampleClientImp{
				Client: tt.fields.Client(),
			}
			got, err := c.GetName(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("exampleClientImp.GetName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				assert.Equal(t, tt.wantErrMessage, err.Error(), tt.name)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exampleClientImp.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
