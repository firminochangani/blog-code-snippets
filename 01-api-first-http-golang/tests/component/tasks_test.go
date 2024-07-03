package component

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/blog-code-snippets/01-api-first-http-golang/tests/client"
)

func TestTasks(t *testing.T) {
	apiClient := createClient(t, "")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	resp, err := apiClient.GetAllTasksWithResponse(ctx)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}

func createClient(t *testing.T, token string) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses("http://localhost:8080", client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		if token != "" {
			req.Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		}
		return nil
	}))
	require.NoError(t, err)

	return cli
}
