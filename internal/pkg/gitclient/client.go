package gitclient

import (
	"os/exec"
	"strings"
	"time"
)

type GitClient interface {
	GetFirstCommit() (string, error)
	GetLastCommit() (string, error)
	GetDateOfHash(hash string) (time.Time, error)
}

type execOptions struct {
	args []string
}

type git struct {
}

func (g git) exec(opts execOptions) (string, error) {
	// TODO: Consider not using a private exec function and hardcode
	// each call to git in the respective command.
	// For now, the lint check is disabled.
	output, err := exec.Command("git", opts.args...).Output() // #nosec G204
	if err != nil {
		return "", err
	}

	return strings.Trim(string(output), "\n"), nil
}

func (g git) GetFirstCommit() (string, error) {
	return g.exec(execOptions{
		args: []string{"rev-list", "--max-parents=0", "HEAD"},
	})
}

func (g git) GetLastCommit() (string, error) {
	return g.exec(execOptions{
		args: []string{"log", "-1", "--pretty=format:%H"},
	})
}

func (g git) GetDateOfHash(hash string) (time.Time, error) {
	date, err := g.exec(execOptions{
		args: []string{"log", "-1", "--format=%cI", hash, "--date=local"},
	})

	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(time.RFC3339, date, time.Local)
}

func NewGitClient() GitClient {
	return git{}
}