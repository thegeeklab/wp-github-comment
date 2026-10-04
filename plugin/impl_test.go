package plugin

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		event    string
		baseURL  string
		key      string
		issueNum int
		wantErr  error
	}{
		{
			name:     "valid settings with explicit key",
			event:    "pull_request",
			baseURL:  "https://api.github.com",
			key:      "my-key",
			issueNum: 42,
		},
		{
			name:     "empty key derives hash from metadata",
			event:    "pull_request",
			baseURL:  "https://api.github.com",
			issueNum: 42,
		},
		{
			name:     "unsupported pipeline event",
			event:    "push",
			baseURL:  "https://api.github.com",
			issueNum: 42,
			wantErr:  ErrPluginEventNotSupported,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CI_PIPELINE_EVENT", tt.event)
			t.Setenv("CI_REPO_OWNER", "octocat")
			t.Setenv("CI_REPO_NAME", "Hello-World")
			t.Setenv("PLUGIN_API_KEY", "test-token")
			t.Setenv("PLUGIN_MESSAGE", "hello")

			p := New(func(_ context.Context) error { return nil })
			_ = p.App.Run(t.Context(), []string{"wp-github-comment"})

			p.Settings = &Settings{
				Message:  "hello",
				BaseURL:  tt.baseURL,
				Key:      tt.key,
				IssueNum: tt.issueNum,
			}

			err := p.Validate()
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.True(t, strings.HasSuffix(p.Settings.BaseURL, "/"))
			assert.NotNil(t, p.Settings.baseURL)
			assert.NotEmpty(t, p.Settings.Key)
		})
	}
}
