package integration

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/service"
	"github.com/enesanbar/url-shortener/internal/adapter"
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/validators"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
	testApp        *fx.App
	MongoContainer *MongoContainer
}

func (s *E2ETestSuite) SetupSuite() {
	s.MongoContainer = NewMongoContainer("root", "password")
	ctx := context.Background()
	err := s.MongoContainer.Start(ctx)
	s.NoError(err)

	_ = os.Setenv("GG_DATASOURCES_MONGO_DEFAULT_HOST", fmt.Sprintf("%s:%s", s.MongoContainer.Hostname(), s.MongoContainer.Port()))
	_ = os.Setenv("GG_DATASOURCES_MONGO_DEFAULT_USERNAME", s.MongoContainer.Username())
	_ = os.Setenv("GG_DATASOURCES_MONGO_DEFAULT_PASSWORD", s.MongoContainer.Password())

	s.startServer()
}

func (s *E2ETestSuite) startServer() {
	testApp := service.New("url-shortener").
		WithModules(
			adapter.Module,
			validators.Module,
		).
		WithConstructor(domain.NewAppConfig).
		Build()
	s.testApp = testApp
	go testApp.Run()

	serverReady := make(chan bool)
	go func(ready chan bool) {
		for {
			_, err := http.Get("http://localhost:9090/url-shortener/_hc")
			if err == nil {
				serverReady <- true
				break
			}
		}
	}(serverReady)

	select {
	case <-time.After(time.Second * 20):
		s.Fail("unable to spin up server in 20 seconds")
	case <-serverReady:
		s.T().Log("server is ready to serve connections")
	}
}

func (s *E2ETestSuite) SetupTest() {
}

func (s *E2ETestSuite) TestCreateWithURL() {
	request, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9090/url-shortener/api/mappings",
		bytes.NewBufferString(`
			{"url": "https://example.com"}
		`),
	)
	s.NoError(err)
	resp, err := http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.StatusCode)

	request, err = http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/api/mappings",
		nil,
	)
	s.NoError(err)

	resp, err = http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *E2ETestSuite) TestCreateWithCode() {
	request, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9090/url-shortener/api/mappings",
		bytes.NewBufferString(`
			{"url": "https://example.com", "code": "testing"}
		`),
	)
	s.NoError(err)

	resp, err := http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.StatusCode)

	request, err = http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/api/mappings/testing",
		nil,
	)
	s.NoError(err)

	resp, err = http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *E2ETestSuite) TestDelete() {
	request, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9090/url-shortener/api/mappings",
		bytes.NewBufferString(`
			{"url": "https://example.com", "code": "delete"}
		`),
	)
	s.NoError(err)

	resp, err := http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.StatusCode)

	request, err = http.NewRequest(
		http.MethodDelete,
		"http://localhost:9090/url-shortener/api/mappings/delete",
		nil,
	)
	s.NoError(err)

	// test delete
	resp, err = http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusNoContent, resp.StatusCode)

	// test deleting deleted item
	resp, err = http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusNotModified, resp.StatusCode)

	// check if item is deleted
	request, err = http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/api/mappings/delete",
		nil,
	)
	s.NoError(err)

	resp, err = http.DefaultClient.Do(request)
	s.NoError(err)
	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *E2ETestSuite) TestRedirectHappyPath() {
	request, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9090/url-shortener/api/mappings",
		bytes.NewBufferString(`
			{"url": "https://example.com", "code": "redirect"}
		`),
	)
	s.NoError(err)

	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.StatusCode)

	request, err = http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/redirect/redirect",
		nil,
	)
	s.NoError(err)

	resp, err = client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusMovedPermanently, resp.StatusCode)
	s.Equal("https://example.com", resp.Header.Get("Location"))
}

func (s *E2ETestSuite) TestRedirectNotFound() {
	client := http.DefaultClient
	request, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/redirect/non-existing",
		nil,
	)
	s.NoError(err)

	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *E2ETestSuite) TestRedirectExpired() {
	request, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9090/url-shortener/api/mappings",
		bytes.NewBufferString(`
			{"url": "https://example.com", "code": "expired", "expires_at": "2021-11-01 13:44:28"}
		`),
	)
	s.NoError(err)

	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.StatusCode)

	request, err = http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/url-shortener/redirect/expired",
		nil,
	)
	s.NoError(err)

	resp, err = client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusMovedPermanently, resp.StatusCode)
	s.NotEqual("https://example.com", resp.Header.Get("Location"))
}

func (s *E2ETestSuite) TearDownSuite() {
	ctx := context.Background()
	s.NoError(s.testApp.Stop(ctx))
	s.NoError(s.MongoContainer.Stop(ctx))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}
