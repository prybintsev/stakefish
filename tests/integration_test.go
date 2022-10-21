package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/prybintsev/stakefish/internal/models"
)

func TestMain(m *testing.M) {
	os.Exit(initContainersAndRunTests(m))
}

func initContainersAndRunTests(m *testing.M) int {
	err := startDockerComposeEnv()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = stopDockerComposeEnv()
		if err != nil {
			panic(err)
		}
	}()

	// Wait for containers to start
	time.Sleep(5 * time.Second)

	return m.Run()
}

const DockerComposePath = "../docker-compose.yaml"

func startDockerComposeEnv() error {
	cmd := exec.Command("docker-compose", "-f", DockerComposePath, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func stopDockerComposeEnv() error {
	cmd := exec.Command("docker-compose", "-f", DockerComposePath, "down", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("docker-compose", "-f", DockerComposePath, "rm", "-f")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func TestApi(t *testing.T) {
	host := "http://localhost:3000"
	t.Run("Test app info", func(t *testing.T) {
		res, err := http.Get(fmt.Sprintf("%s/", host))
		require.NoError(t, err)

		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		var appInfo models.AppInfo
		err = json.Unmarshal(body, &appInfo)
		require.NoError(t, err)
		require.False(t, appInfo.Kubernetes)
	})

	t.Run("Test validate", func(t *testing.T) {
		req := models.ValidateRequest{
			IP: "1.1.1.1",
		}
		reqStr, err := json.Marshal(&req)
		require.NoError(t, err)
		res, err := http.Post(fmt.Sprintf("%s/v1/tools/validate", host), "application/json",
			strings.NewReader(string(reqStr)))
		require.NoError(t, err)

		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		var validateResponse models.ValidateResponse
		err = json.Unmarshal(body, &validateResponse)
		require.NoError(t, err)
		require.True(t, validateResponse.Status)
	})

	t.Run("Test lookup and history", func(t *testing.T) {
		res, err := http.Get(fmt.Sprintf("%s/v1/tools/lookup?domain=google.com", host))
		require.NoError(t, err)
		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		var lookupResponse models.Lookup
		err = json.Unmarshal(body, &lookupResponse)
		require.NoError(t, err)
		require.Equal(t, "google.com", lookupResponse.Domain)
		require.True(t, len(lookupResponse.Addresses) > 0)

		res, err = http.Get(fmt.Sprintf("%s/v1/history", host))
		require.NoError(t, err)
		body, err = io.ReadAll(res.Body)
		require.NoError(t, err)

		var historyResponse []models.Lookup
		err = json.Unmarshal(body, &historyResponse)
		require.Equal(t, 1, len(historyResponse))
		require.Equal(t, lookupResponse, historyResponse[0])
	})
}
