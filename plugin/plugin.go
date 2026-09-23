package plugin

import (
	"fmt"
	"net/url"
	"slices"

	gh "github.com/thegeeklab/wp-github-comment/github"
	plugin_base "github.com/thegeeklab/wp-plugin-go/v7/plugin"
	"github.com/urfave/cli/v3"
)

//go:generate go run ../hack/docs-gen/main.go -output=../docs/data/data.yaml

// Plugin implements provide the plugin.
type Plugin struct {
	*plugin_base.Plugin
	Settings *Settings
}

// Settings for the Plugin.
type Settings struct {
	BaseURL     string
	IssueNum    int
	Key         string
	Message     string
	Update      bool
	APIKey      string
	SkipMissing bool
	IsFile      bool

	baseURL *url.URL
}

func New(e plugin_base.ExecuteFunc, build ...string) *Plugin {
	p := &Plugin{
		Settings: &Settings{},
	}

	options := plugin_base.Options{
		Name:        "wp-github-comment",
		Description: "Add comments to GitHub Issues and Pull Requests",
		Flags: slices.Concat(
			plugin_base.LoggingFlags(plugin_base.FlagsPluginCategory),
			plugin_base.NetworkFlags(plugin_base.FlagsPluginCategory),
			Flags(p.Settings, plugin_base.FlagsPluginCategory),
		),
		Execute:             p.run,
		HideWoodpeckerFlags: true,
	}

	if len(build) > 0 {
		options.Version = build[0]
	}

	if len(build) > 1 {
		options.VersionMetadata = fmt.Sprintf("date=%s", build[1])
	}

	if e != nil {
		options.Execute = e
	}

	p.Plugin = plugin_base.New(options)

	return p
}

// Flags returns a slice of CLI flags for the plugin.
func Flags(settings *Settings, category string) []cli.Flag {
	return []cli.Flag{
		// Personal access token to access the GitHub API.
		&cli.StringFlag{
			Name:        "api-key",
			Sources:     cli.EnvVars("PLUGIN_API_KEY", "GITHUB_COMMENT_API_KEY"),
			Usage:       "personal access token to access the GitHub API",
			Destination: &settings.APIKey,
			Category:    category,
			Required:    true,
		},
		// Api url.
		//
		// Only need to be changed for GitHub enterprise in most cases.
		&cli.StringFlag{
			Name:        "base-url",
			Sources:     cli.EnvVars("PLUGIN_BASE_URL", "GITHUB_COMMENT_BASE_URL"),
			Usage:       "API URL",
			Value:       gh.DefaultBaseURL,
			Destination: &settings.BaseURL,
			Category:    category,
		},
		// Unique identifier to assign to a comment.
		//
		// The identifier is used to update an existing comment.
		&cli.StringFlag{
			Name:        "key",
			Sources:     cli.EnvVars("PLUGIN_KEY", "GITHUB_COMMENT_KEY"),
			Usage:       "unique identifier to assign to a comment",
			Destination: &settings.Key,
			Category:    category,
		},
		// Path to file or string that contains the comment text.
		&cli.StringFlag{
			Name:        "message",
			Sources:     cli.EnvVars("PLUGIN_MESSAGE", "GITHUB_COMMENT_MESSAGE"),
			Usage:       "path to file or string that contains the comment text",
			Destination: &settings.Message,
			Category:    category,
			Required:    true,
		},
		// Enable update of an existing comment that matches the key.
		&cli.BoolFlag{
			Name:        "update",
			Sources:     cli.EnvVars("PLUGIN_UPDATE", "GITHUB_COMMENT_UPDATE"),
			Usage:       "enable update of an existing comment that matches the key",
			Value:       false,
			Destination: &settings.Update,
			Category:    category,
		},
		// Skip comment creation if the given message file does not exist.
		&cli.BoolFlag{
			Name:        "skip-missing",
			Sources:     cli.EnvVars("PLUGIN_SKIP_MISSING", "GITHUB_COMMENT_SKIP_MISSING"),
			Usage:       "skip comment creation if the given message file does not exist",
			Value:       false,
			Destination: &settings.SkipMissing,
			Category:    category,
		},
	}
}
