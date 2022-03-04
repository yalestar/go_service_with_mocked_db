package healthcheck

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

// a map of fake requests and responses
type StubDatastore struct {
    requests map[string]string
}

func (stub *StubDatastore) GetHealthCheck(path string) string {
    fakeResponse := stub.requests[path]
    return fakeResponse
}

func TestHealthCheck(t *testing.T) {
    
    stubDB := StubDatastore{
        map[string]string{
            "/health": "ok",
            "/ljsdf":  "BADBAD",
        },
    }
    server := &NcpdpServer{Store: &stubDB}
    
    t.Run(
        "gets health check if path is /health", func(t *testing.T) {
            request, _ := http.NewRequest(http.MethodGet, "/health", nil)
            response := httptest.NewRecorder()
            
            server.ServeHTTP(response, request)
            
            got := response.Body.String()
            want := "ok"
            
            assert.Equal(t, want, got)
        },
    )
    
    t.Run(
        "says BAD if path is not /health", func(t *testing.T) {
            request, _ := http.NewRequest(http.MethodGet, "/ljsdf", nil)
            response := httptest.NewRecorder()
            
            server.ServeHTTP(response, request)
            got := response.Body.String()
            want := "BADBAD"
            
            assert.Equal(t, want, got)
            
        },
    )
}
