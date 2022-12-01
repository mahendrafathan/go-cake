package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mahendrafathan/go-cake/controllers/mock"
	"github.com/mahendrafathan/go-cake/models"
)

func TestCakeList(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		preTest  func(mock *mock.MockCakeRepository)
		wantCode int
	}{
		{
			name: "success 200",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://test.com/account/123",
					nil,
				),
			},
			preTest: func(mock *mock.MockCakeRepository) {
				mock.EXPECT().
					FindAll().
					Return([]models.Cake{
						models.Cake{
							Title: "cake",
						},
					}, nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "success 500",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://test.com/account/123",
					nil,
				),
			},
			preTest: func(mock *mock.MockCakeRepository) {
				mock.EXPECT().
					FindAll().
					Return([]models.Cake{
						models.Cake{
							Title: "cake",
						},
					}, fmt.Errorf("error"))
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock.NewMockCakeRepository(ctrl)
			cakeCtrl := CakeUsecase{
				cakeRepo: mock,
			}
			tt.preTest(mock)
			cakeCtrl.CakeList(tt.args.w, tt.args.r)
		})
	}
}
