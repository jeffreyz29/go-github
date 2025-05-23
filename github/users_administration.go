// Copyright 2014 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// PromoteSiteAdmin promotes a user to a site administrator of a GitHub Enterprise instance.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/users#promote-a-user-to-be-a-site-administrator
//
//meta:operation PUT /users/{username}/site_admin
func (s *UsersService) PromoteSiteAdmin(ctx context.Context, user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/site_admin", user)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DemoteSiteAdmin demotes a user from site administrator of a GitHub Enterprise instance.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/users#demote-a-site-administrator
//
//meta:operation DELETE /users/{username}/site_admin
func (s *UsersService) DemoteSiteAdmin(ctx context.Context, user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/site_admin", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// UserSuspendOptions represents the reason a user is being suspended.
type UserSuspendOptions struct {
	Reason *string `json:"reason,omitempty"`
}

// Suspend a user on a GitHub Enterprise instance.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/users#suspend-a-user
//
//meta:operation PUT /users/{username}/suspended
func (s *UsersService) Suspend(ctx context.Context, user string, opts *UserSuspendOptions) (*Response, error) {
	u := fmt.Sprintf("users/%v/suspended", user)

	req, err := s.client.NewRequest("PUT", u, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Unsuspend a user on a GitHub Enterprise instance.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/users#unsuspend-a-user
//
//meta:operation DELETE /users/{username}/suspended
func (s *UsersService) Unsuspend(ctx context.Context, user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/suspended", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
