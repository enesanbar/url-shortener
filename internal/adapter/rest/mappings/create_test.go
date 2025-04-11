package mappings

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enesanbar/go-service/utils"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/router"
	"github.com/enesanbar/url-shortener/internal/adapter/rest/mappings/mocks"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUrl(t *testing.T) {
	zapLogger, _ := log.NewZapLogger()
	logger := log.NewFactory(zapLogger)

	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		inModelMock    func() *mocks.CreateMappingUseCase
		expectedStatus int
		expectedError  string
	}{
		{
			desc: "Should create short url for a given valid URL",
			inRequest: func() *http.Request {
				validRequest := buildValidRequest()
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", validRequest)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.CreateMappingUseCase {
				output := &domain.Mapping{}

				mockGetModel := &mocks.CreateMappingUseCase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(output, nil).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusCreated,
			expectedError:  "",
		},
		{
			desc: "Should return error given invalid request body",
			inRequest: func() *http.Request {
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", bytes.NewBufferString(`url: dsfasdf}`))
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.CreateMappingUseCase {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "unable to serialize JSON body.",
		},
		{
			desc: "Should return error when request body validation fails",
			inRequest: func() *http.Request {
				validRequest := buildInvalidRequest()
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", validRequest)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.CreateMappingUseCase {
				mockGetModel := &mocks.CreateMappingUseCase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.Error{
					Code:    errors.EINVALID,
					Message: "validation failed",
				}).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "validation failed",
		},
		{
			desc: "Should return conflict given code already exists",
			inRequest: func() *http.Request {
				validRequest := buildValidRequest()
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", validRequest)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.CreateMappingUseCase {
				mockGetModel := &mocks.CreateMappingUseCase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.Error{
					Code:    errors.ECONFLICT,
					Message: "code already exists in the database",
				}).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusConflict,
			expectedError:  "code already exists in the database",
		},
		{
			desc: "Should return internal server error on database error",
			inRequest: func() *http.Request {
				validRequest := buildValidRequest()
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", validRequest)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.CreateMappingUseCase {
				mockGetModel := &mocks.CreateMappingUseCase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.Error{
					Code:    errors.EINTERNAL,
					Message: "unable to create mapping",
				}).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "unable to create mapping",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(scenario.inRequest(), rec)

			subCtx := context.WithValue(c.Request().Context(), utils.ContextKeyRequestID, "someid")
			c.SetRequest(c.Request().WithContext(subCtx))

			// build handler
			handler := NewCreateMappingAdapter(router.NewBaseHandler(logger), scenario.inModelMock(), response.NewMappingPresenter(), logger)
			err := handler.Handle(c)

			if assert.NoError(t, err) {
				assert.Equal(t, scenario.expectedStatus, rec.Code)

				var resp router.ApiResponse
				err := json.NewDecoder(rec.Body).Decode(&resp)
				assert.NoError(t, err)
				assert.Equal(t, scenario.expectedError, resp.Error)
			}
		})
	}
}

func buildValidRequest() io.Reader {
	requestData := &create.Request{
		URL:  "http://www.google.com/",
		Code: "SADASFF",
	}

	requestJson, _ := json.Marshal(requestData)
	return bytes.NewBuffer(requestJson)
}

func buildInvalidRequest() io.Reader {
	requestData := &create.Request{
		URL: "htt://www.google.com",
	}

	requestJson, _ := json.Marshal(requestData)
	return bytes.NewBuffer(requestJson)
}
