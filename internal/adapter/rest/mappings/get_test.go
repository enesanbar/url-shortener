package mappings

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enesanbar/go-service/utils"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/router"
	"github.com/enesanbar/url-shortener/internal/adapter/rest/mappings/mocks"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetMapping(t *testing.T) {
	zapLogger, _ := log.NewZapLogger()
	logger := log.NewFactory(zapLogger)

	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		inModelMock    func() *mocks.GetMappingUsecase
		expectedStatus int
		expectedError  string
	}{
		{
			desc: "Should return success when mapping exists",
			inRequest: func() *http.Request {
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings/test", nil)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.GetMappingUsecase {
				output := &domain.Mapping{}

				mockGetModel := &mocks.GetMappingUsecase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(output, nil).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
		{
			desc: "Should return error when mapping doesn't exist",
			inRequest: func() *http.Request {
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings/test", nil)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.GetMappingUsecase {
				output := &domain.Mapping{}

				mockGetModel := &mocks.GetMappingUsecase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(output, errors.Error{
					Code:    errors.ENOTFOUND,
					Message: "mapping not found",
				}).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  "mapping not found",
		},
		{
			desc: "Should return internal server error on database error",
			inRequest: func() *http.Request {
				request, err := http.NewRequest(http.MethodPost, "/url-shortener/api/mappings", nil)
				require.NoError(t, err)
				return request
			},
			inModelMock: func() *mocks.GetMappingUsecase {
				mockGetModel := &mocks.GetMappingUsecase{}
				mockGetModel.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.Error{
					Code:    errors.EINTERNAL,
					Message: "unable to get mapping",
				}).Once()

				return mockGetModel
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "unable to get mapping",
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
			handler := NewGetMappingAdapter(GetMappingAdapterParams{
				BaseHandler: router.NewBaseHandler(logger),
				Presenter:   response.NewMappingPresenter(),
				Interactor:  scenario.inModelMock(),
			})
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
