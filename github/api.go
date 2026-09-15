package github

import (
	"context"

	"github.com/google/go-github/v92/github"
)

// APIClient is an interface that wraps the GitHub API client.
//
//nolint:lll
type IssueService interface {
	CreateComment(ctx context.Context, owner, repo string, number int, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error)
	UpdateComment(ctx context.Context, owner, repo string, commentID int64, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error)
	ListComments(ctx context.Context, owner, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)
}

type IssueServiceImpl struct {
	client *github.Client
}

// CreateComment wraps the CreateComment method of the github.IssuesService.
//
//nolint:lll
func (s *IssueServiceImpl) CreateComment(ctx context.Context, owner, repo string, number int, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error) {
	return s.client.Issues.CreateComment(ctx, owner, repo, number, comment)
}

// UpdateComment wraps the UpdateComment method of the github.IssuesService.
//
//nolint:lll
func (s *IssueServiceImpl) UpdateComment(ctx context.Context, owner, repo string, commentID int64, comment github.IssueCommentRequest) (*github.IssueComment, *github.Response, error) {
	return s.client.Issues.UpdateComment(ctx, owner, repo, commentID, comment)
}

// ListComments wraps the ListComments method of the github.IssuesService.
//
//nolint:lll
func (s *IssueServiceImpl) ListComments(ctx context.Context, owner, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	return s.client.Issues.ListComments(ctx, owner, repo, number, opts)
}
