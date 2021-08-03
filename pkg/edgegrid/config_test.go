package edgegrid

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
	"os"
	"testing"
)

func TestConfig_FromFile(t *testing.T) {
	tests := map[string]struct {
		fileName  string
		section   string
		expected  Config
		withError error
	}{
		"valid file and section": {
			fileName: "edgerc",
			section:  "test",
			expected: Config{
				Host:         "xxxx-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx.luna.akamaiapis.net",
				ClientToken:  "xxxx-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx",
				ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=",
				AccessToken:  "xxxx-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx",
				MaxBody:      131072,
			},
		},
		"file does not exist": {
			fileName:  "test",
			section:   "test",
			withError: ErrLoadingFile,
		},
		"section does not exist": {
			fileName:  "edgerc",
			section:   "abc",
			withError: ErrSectionDoesNotExist,
		},
		"missing host": {
			fileName:  "edgerc",
			section:   "missing-host",
			withError: ErrRequiredOptionEdgerc,
		},
		"missing client secret": {
			fileName:  "edgerc",
			section:   "missing-client-secret",
			withError: ErrRequiredOptionEdgerc,
		},
		"missing client token": {
			fileName:  "edgerc",
			section:   "missing-client-token",
			withError: ErrRequiredOptionEdgerc,
		},
		"missing access token": {
			fileName:  "edgerc",
			section:   "missing-access-token",
			withError: ErrRequiredOptionEdgerc,
		},
		"invalid host with slash at the end": {
			fileName:  "edgerc",
			section:   "slash-at-the-end-of-host-value",
			withError: ErrHostContainsSlashAtTheEnd,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cfg := Config{}
			err := cfg.FromFile(fmt.Sprintf("test/%s", test.fileName), test.section)
			if test.withError != nil {
				assert.True(t, errors.Is(err, test.withError), "want: %v; got: %v", test.withError, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.expected, cfg)
		})
	}
}

func TestConfig_FromEnv(t *testing.T) {
	tests := map[string]struct {
		section   string
		envs      map[string]string
		expected  Config
		withError error
	}{
		"default section, valid envs, default max body": {
			section: "default",
			envs: map[string]string{
				"AKAMAI_HOST":          "test-host",
				"AKAMAI_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_ACCESS_TOKEN":  "test-access-token",
			},
			expected: Config{
				Host:         "test-host",
				ClientToken:  "test-client-token",
				ClientSecret: "test-client-secret",
				AccessToken:  "test-access-token",
				MaxBody:      131072,
			},
		},
		"default section, valid envs": {
			section: "default",
			envs: map[string]string{
				"AKAMAI_HOST":          "test-host",
				"AKAMAI_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_ACCESS_TOKEN":  "test-access-token",
				"AKAMAI_MAX_BODY":      "123",
			},
			expected: Config{
				Host:         "test-host",
				ClientToken:  "test-client-token",
				ClientSecret: "test-client-secret",
				AccessToken:  "test-access-token",
				MaxBody:      123,
			},
		},
		"custom section, valid envs": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_HOST":          "test-host",
				"AKAMAI_TEST_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_TEST_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_TEST_ACCESS_TOKEN":  "test-access-token",
			},
			expected: Config{
				Host:         "test-host",
				ClientToken:  "test-client-token",
				ClientSecret: "test-client-secret",
				AccessToken:  "test-access-token",
				MaxBody:      131072,
			},
		},
		"custom section, missing host": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_TEST_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_TEST_ACCESS_TOKEN":  "test-access-token",
			},
			withError: ErrRequiredOptionEnv,
		},
		"custom section, missing client secret": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_HOST":         "test-host",
				"AKAMAI_TEST_CLIENT_TOKEN": "test-client-token",
				"AKAMAI_TEST_ACCESS_TOKEN": "test-access-token",
			},
			withError: ErrRequiredOptionEnv,
		},
		"custom section, missing client token": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_HOST":          "test-host",
				"AKAMAI_TEST_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_TEST_ACCESS_TOKEN":  "test-access-token",
			},
			withError: ErrRequiredOptionEnv,
		},
		"custom section, missing access token": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_HOST":          "test-host",
				"AKAMAI_TEST_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_TEST_CLIENT_SECRET": "test-client-secret",
			},
			withError: ErrRequiredOptionEnv,
		},
		"custom section, slash at the end of host value": {
			section: "test",
			envs: map[string]string{
				"AKAMAI_TEST_HOST":          "test-host/",
				"AKAMAI_TEST_CLIENT_TOKEN":  "test-client-token",
				"AKAMAI_TEST_CLIENT_SECRET": "test-client-secret",
				"AKAMAI_TEST_ACCESS_TOKEN":  "test-access-token",
			},
			withError: ErrHostContainsSlashAtTheEnd,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for k, v := range test.envs {
				require.NoError(t, os.Setenv(k, v))
			}
			defer func() {
				for k := range test.envs {
					require.NoError(t, os.Unsetenv(k))
				}
			}()
			cfg := Config{}
			err := cfg.FromEnv(test.section)
			if test.withError != nil {
				assert.True(t, errors.Is(err, test.withError), "want: %v; got: %v", test.withError, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.expected, cfg)
		})
	}
}
