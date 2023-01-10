package github_test

import (
	"testing"

	gh "github.com/qbarrand/gitstream/internal/github"
	"github.com/stretchr/testify/assert"
)

func TestGetAssignee(t *testing.T) {

	t.Run("known author", func(t *testing.T) {

		res := gh.GetAssignee("Yoni Bettan")

		assert.Equal(t, res, "ybettan")
	})

	t.Run("unknown author", func(t *testing.T) {

		res := gh.GetAssignee("John Doh")

		assert.NotEmpty(t, res)

		var isDefaultReviewer bool
		for _, dr := range gh.DefaultReviewers {
			if res == dr {
				isDefaultReviewer = true
				break
			}
		}
		assert.True(t, isDefaultReviewer)
	})
}
