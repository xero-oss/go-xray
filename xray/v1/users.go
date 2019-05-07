package v1

import (
	"context"
	"fmt"
	"net/http"
)

type UsersService Service

type User struct {
	Admin    *bool   `json:"admin,omitempty"`
	Blocked  *bool   `json:"blocked,omitempty"`
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

// Description: Gets a list of all users in the system
// Security:  Requires an admin user
// Usage: client.V1.Users.ListUsers(ctx)
func (s *UsersService) ListUsers(ctx context.Context) (*[]User, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("GET", "/api/v1/users", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	users := new([]User)
	resp, err := s.client.Do(ctx, req, &users)
	return users, resp, err
}

// Description: Gets a named user
// Security:  Requires an admin user
// Usage: client.V1.Users.GetUser(ctx, "name")
func (s *UsersService) GetUser(ctx context.Context, name string) (*User, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/users/%s", name)
	req, err := s.client.NewJSONEncodedRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	user := new(User)
	resp, err := s.client.Do(ctx, req, &user)
	return user, resp, err
}

// Description: Creates a new User
// Security:  Requires an admin user
// Usage: client.V1.Users.CreateUser(ctx, user)
func (s *UsersService) CreateUser(ctx context.Context, user *User) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/users", user)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Updates a user.
// Security:  Requires an admin user
// Usage: client.V1.Users.UpdateUser(ctx, "name", user)
func (s *UsersService) UpdateUser(ctx context.Context, name string, user *User) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/users/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, user)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Deletes a User
// Security:  Requires an admin users
// Usage: client.V1.Users.DeleteUser(ctx, "name")
func (s *UsersService) DeleteUser(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/users/%s", name)
	req, err := s.client.NewJSONEncodedRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
