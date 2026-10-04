package github

import (
	"context"

	"github.com/google/go-github/v92/github"
)

// IssueService is an interface that wraps the GitHub Issues API client.
//
//nolint:lll
type IssueService interface {
	CreateComment(ctx context.Context, owner, repo string, number int, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error)
	UpdateComment(ctx context.Context, owner, repo string, commentID int64, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error)
	ListComments(ctx context.Context, owner, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)
}
