package gorhsm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const rhsmTokenURL = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"

// Token describes the response from RHSM token refresh requests
type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	SessionState     string `json:"session_state"`
	TokenType        string `json:"token_type"`
}

type tokenErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	ErrorMsg         string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// GenerateAccessToken generats a new API bearer token
func GenerateAccessToken(refreshToken string) (*Token, error) {
	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", "rhsm-api")
	form.Add("refresh_token", refreshToken)

	req, err := http.NewRequest(http.MethodPost, rhsmTokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkTokenResponse(resp)
	if err != nil {
		return nil, err
	}

	token := new(Token)

	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return nil, fmt.Errorf("error decoding token response body")
	}

	return token, nil

}

func checkTokenResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &tokenErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.ErrorMsg = string(data)
		}
	}

	return errorResponse
}

func (r *tokenErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.ErrorMsg+": "+r.ErrorDescription)
}
