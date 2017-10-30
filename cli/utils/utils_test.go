package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * githubリポジトリ名がパースできることを確認する
 */
func TestSplitGithubRepository(t *testing.T) {

	{
		user, repo := SplitGithubRepository("git@github.com:user/repo.git")
		assert.Equal(t, user, "user")
		assert.Equal(t, repo, "repo")
	}

	{
		user, repo := SplitGithubRepository("https://github.com/user/repo.git")
		assert.Equal(t, user, "user")
		assert.Equal(t, repo, "repo")
	}

	{
		user, repo := SplitGithubRepository("https://GITHUB_API_KEY:x-oauth-basic@github.com/user/repo.git")
		assert.Equal(t, user, "user")
		assert.Equal(t, repo, "repo")
	}
}
