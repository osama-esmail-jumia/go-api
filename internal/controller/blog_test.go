package controller_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	controllerI "go-api/definitions/controller"
	"go-api/definitions/request"
	"go-api/definitions/response"
	"go-api/internal/controller"
	"go-api/internal/handler"
	serviceMocks "go-api/mocks/definitions/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type controllerMocks struct {
	serviceMock *serviceMocks.Blog
	routerMock  *gin.Engine
}

func setup() (controllerI.Blog, controllerMocks) {
	routerMock := gin.Default()
	gin.SetMode(gin.TestMode)
	
	serviceMock := new(serviceMocks.Blog)
	
	return controller.NewBlog(
			handler.NewBlog(),
			serviceMock,
		), controllerMocks{
			serviceMock,
			routerMock,
		}
}

func TestBlog_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			ctrl, cMocks := setup()
			
			// test data
			req := request.BlogCreate{
				Title: "foo",
			}
			resp := response.Blog{
				ID:    1,
				Title: "foo",
			}
			
			// mocks
			cMocks.serviceMock.On("Create", req).Return(resp, nil)
			cMocks.routerMock.POST("/blog", ctrl.Create)
			
			// http request
			body, err := json.Marshal(&req)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}
			httpReq, err := http.NewRequest(http.MethodPost, "/blog", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}
			reqRecorder := httptest.NewRecorder()
			cMocks.routerMock.ServeHTTP(reqRecorder, httpReq)
			
			// assert
			assert.Equal(t, http.StatusCreated, reqRecorder.Code)
		},
	)
}
