package service_test

import (
	"github.com/stretchr/testify/assert"
	"go-api/definitions/model"
	"go-api/definitions/request"
	"go-api/definitions/response"
	serviceI "go-api/definitions/service"
	"go-api/internal/service"
	mapperMocks "go-api/mocks/definitions/mapper"
	repoMocks "go-api/mocks/definitions/repository"
	"testing"
)

type serviceMocks struct {
	repoMock   *repoMocks.Blog
	mapperMock *mapperMocks.Blog
}

func setup() (serviceI.Blog, serviceMocks) {
	repoMock := new(repoMocks.Blog)
	mapperMock := new(mapperMocks.Blog)
	
	return service.NewBlog(
			repoMock,
			mapperMock,
		), serviceMocks{
			repoMock,
			mapperMock,
		}
}

func TestBlog_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			svc, mocks := setup()
			
			// test data
			req := request.BlogCreate{
				Title: "foo",
			}
			resp := response.Blog{
				ID:    1,
				Title: "foo",
			}
			title := "foo"
			m := model.Blog{
				ID:    1,
				Title: &title,
			}
			
			// mocks
			mocks.mapperMock.On("CreateRequestToModel", req).Return(m, nil)
			mocks.mapperMock.On("ModelToResponse", m).Return(resp)
			mocks.repoMock.On("Create", &m).Return(nil)
			
			// call
			svcResp, err := svc.Create(req)
			
			// assert
			assert.Nil(t, err)
			assert.Equal(t, svcResp, resp)
		},
	)
}
