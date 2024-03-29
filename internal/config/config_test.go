package config

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	// This test checks default values.
	expected := Config{
		CommitMarkup: "Upstream-Commit",
		Downstream: Downstream{
			LocalRepoPath: ".",
			MainBranch:    "main",
			MaxOpenItems:  -1,
			OwnersFile:    "OWNERS",
		},
		Upstream: Upstream{Ref: "main"},
	}

	cfg, err := ReadConfig(strings.NewReader("---"))
	require.NoError(t, err)
	assert.Equal(t, &expected, cfg)
}

func TestReadConfigFile(t *testing.T) {
	// This test overrides default values
	since := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	expected := Config{
		CommitMarkup: "test",
		Downstream: Downstream{
			GitHubRepoName: "owner/repo",
			LocalRepoPath:  "some-path",
			MainBranch:     "some-branch",
			MaxOpenItems:   3,
			OwnersFile:     "some-dir/some-file",
		},
		Diff: Diff{
			CommitsSince: &since,
		},
		LogLevel: 1000,
		Sync: Sync{
			BeforeCommit: [][]string{
				{"command", "one"},
				{"command", "two"},
			},
		},
		Upstream: Upstream{
			Ref: "some-ref",
			URL: "https://url.to.some/git/repo",
		},
	}

	cfg, err := ReadConfigFile("testdata/config.yml")
	require.NoError(t, err)
	assert.Equal(t, &expected, cfg)
}
