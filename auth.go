package auth0_authorization_extension

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.ClientId == "" || c.Auth.ClientSecret == "" {
		return nil, fmt.Errorf("Define Auth0 ClientId and ClientSecret")
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.Auth.ClientId)
	data.Set("client_secret", c.Auth.ClientSecret)
	data.Set("audience", c.Auth.Audience)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", c.HostURL), strings.NewReader(string(rb)))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
