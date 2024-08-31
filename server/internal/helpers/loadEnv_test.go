package helpers

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoadEnvTestSuite struct {
	suite.Suite
}

func (suite *LoadEnvTestSuite) SetupTest() {
	GetAppModeFunc = GetAppMode
}

func (suite *LoadEnvTestSuite) TestLoadEnv_Success() {
	GetAppModeFunc = func() string {
		return ".env.test"
	}

	os.WriteFile(".env.test", []byte("KEY=VALUE"), 0644)
	defer os.Remove(".env.test")

	LoadEnv()

	assert.Equal(suite.T(), "VALUE", os.Getenv("KEY"))
}

func (suite *LoadEnvTestSuite) TestLoadEnv_Failure() {
	GetAppModeFunc = func() string {
		return ".env.notexist"
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")
	cmd.Env = append(os.Environ(), "HELPER_PROCESS=1")

	err := cmd.Run()

	exitError, ok := err.(*exec.ExitError)
	assert.True(suite.T(), ok)
	assert.False(suite.T(), exitError.Success())
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("HELPER_PROCESS") == "1" {
		LoadEnv()
	}
}

func TestLoadEnvSuite(t *testing.T) {
	suite.Run(t, new(LoadEnvTestSuite))
}

func TestGetAppMode(t *testing.T) {
	testCases := []struct {
		envValue string
		expected string
	}{
		{"production", ".env"},
		{"docker", ".env.docker"},
		{"", ".env.dev"},
	}

	for _, tesCase := range testCases {
		os.Setenv("APP_ENV", tesCase.envValue)

		result := GetAppMode()

		assert.Equal(t, tesCase.expected, result)
	}

	os.Unsetenv("APP_ENV")
}
