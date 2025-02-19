package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type ClientTestSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	client *Client
}

func (suite *ClientTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.client = NewClient("http://example.com", 30*time.Second)
}

func (suite *ClientTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ClientTestSuite) TestPostReturnsResponseBodyOnSuccess() {
	body := `{"key":"value"}`
	suite.client.HTTPClient = &http.Client{
		Transport: &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
			},
		},
	}

	responseBody, err := suite.client.Post("/test", strings.NewReader(body))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []byte(body), responseBody)
}

func (suite *ClientTestSuite) TestPostReturnsErrorOnNonCreatedStatus() {
	suite.client.HTTPClient = &http.Client{
		Transport: &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader("")),
			},
		},
	}

	responseBody, err := suite.client.Post("/test", strings.NewReader(`{"key":"value"}`))
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), responseBody)
}

func (suite *ClientTestSuite) TestPostReturnsErrorOnClientFailure() {
	suite.client.HTTPClient = &http.Client{
		Transport: &mockTransport{
			err: errors.New("client error"),
		},
	}

	responseBody, err := suite.client.Post("/test", strings.NewReader(`{"key":"value"}`))
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), responseBody)
}

func (suite *ClientTestSuite) TestPostReturnsErrorOnReadFailure() {
	suite.client.HTTPClient = &http.Client{
		Transport: &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(&errorReader{}),
			},
		},
	}

	responseBody, err := suite.client.Post("/test", strings.NewReader(`{"key":"value"}`))
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), responseBody)
}

type mockTransport struct {
	response *http.Response
	err      error
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (e *errorReader) Close() error {
	return nil
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
