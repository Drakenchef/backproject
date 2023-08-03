package handler

import (
	"bytes"
	"errors"
	"github.com/drakenchef/backproject"
	"github.com/drakenchef/backproject/pkg/service"
	mock_service "github.com/drakenchef/backproject/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorizarion, user backproject.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           backproject.User
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: backproject.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehaviour: func(s *mock_service.MockAuthorizarion, user backproject.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"username":"test","password":"qwerty"}`,
			mockBehaviour:       func(s *mock_service.MockAuthorizarion, user backproject.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: backproject.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehaviour: func(s *mock_service.MockAuthorizarion, user backproject.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorizarion(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorizarion: auth}
			handler := NewHandler(services)

			//test Server

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//test Request

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			//Perform Request

			r.ServeHTTP(w, req)

			// Assert

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
