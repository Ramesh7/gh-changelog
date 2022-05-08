package writer_test

import (
	"bytes"
	"regexp"
	"testing"
	"time"

	"github.com/chelnak/gh-changelog/internal/pkg/changelog"
	"github.com/chelnak/gh-changelog/internal/pkg/writer"
	"github.com/chelnak/gh-changelog/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_ItWritesOutAChangelogInTheCorrectFormat(t *testing.T) {
	mockChangelog := &mocks.Changelog{}

	mockChangelog.On("GetRepoName").Return("TestRepo")
	mockChangelog.On("GetRepoOwner").Return("TestOwner")
	mockChangelog.On("GetEntries").Return([]changelog.Entry{
		{
			CurrentTag:  "v1.0.0",
			PreviousTag: "v0.0.0",
			Date:        time.Now(),
			Added:       []string{"Added 1", "Added 2"},
			Changed:     []string{"Changed 1", "Changed 2"},
			Deprecated:  []string{"Deprecated 1", "Deprecated 2"},
			Removed:     []string{"Removed 1", "Removed 2"},
			Fixed:       []string{"Fixed 1", "Fixed 2"},
			Security:    []string{"Security 1", "Security 2"},
			Other:       []string{"Other 1", "Other 2"},
		},
	})

	mockChangelog.On("GetUnreleased").Return([]string{"Unreleased 1", "Unreleased 2"})

	var buf bytes.Buffer
	err := writer.Write(&buf, mockChangelog)

	assert.NoError(t, err)
	mockChangelog.AssertExpectations(t)

	assert.Regexp(t, "## Unreleased", buf.String())
	assert.Regexp(t, "- Unreleased 1", buf.String())
	assert.Regexp(t, "- Unreleased 2", buf.String())

	assert.Regexp(t, regexp.MustCompile(`## \[v1.0.0\]\(https:\/\/github.com\/TestOwner\/TestRepo\/tree\/v1.0.0\)`), buf.String())
	assert.Regexp(t, "### Added", buf.String())
	assert.Regexp(t, "- Added 1", buf.String())
	assert.Regexp(t, "- Added 2", buf.String())

	assert.Regexp(t, "### Other", buf.String())
	assert.Regexp(t, "- Other 1", buf.String())
	assert.Regexp(t, "- Other 2", buf.String())
}