package hyperdb

/**
 * HyperDB Go Client
 *
 * @author Afaan Bilal
 * @link   https://afaan.dev
 */

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HyperDB struct {
	address  string
	username string
	password string

	authEnabled bool
	token       string
}

const R_PONG string = "PONG"
const R_TRUE string = "YES"
const R_OK string = "OK"
const R_INVALID_CREDENTIALS string = "INVALID_CREDENTIALS"
const R_AUTH_FAILED string = "AUTH_FAILED"

func Create(address string, username string, password string) HyperDB {
	return HyperDB{address: address, username: username, password: password, authEnabled: username != "" && password != ""}
}

func (h *HyperDB) http(url string, method string, body string) string {
	if method == "" {
		method = "GET"
	}

	req, err := http.NewRequest(method, h.address+"/"+url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	if h.authEnabled {
		if h.token == "" {
			h.auth()
		}

		req.Header.Add("Auth", h.token)
	}

	response := sendRequest(req)
	if response == R_AUTH_FAILED {
		h.auth()
		req.Header.Set("Auth", h.token)
		response = sendRequest(req)
	}

	return response
}

func sendRequest(req *http.Request) string {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func (h *HyperDB) auth() {
	req, err := http.NewRequest(http.MethodPost, h.address+"/auth", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("username", h.username)
	req.Header.Add("password", h.password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	response := string(body)

	if response == R_INVALID_CREDENTIALS {
		panic("Invalid credentials")
	}

	h.token = response
}

func (h *HyperDB) Ping() bool {
	return h.http("ping", "", "") == R_PONG
}

func (h *HyperDB) Version() string {
	return h.http("", "", "")
}

func (h *HyperDB) Has(key string) bool {
	return h.http(fmt.Sprintf("has/%s", key), "", "") == R_TRUE
}

func (h *HyperDB) Get(key string) string {
	return h.http(fmt.Sprintf("data/%s", key), "", "")
}

func (h *HyperDB) Set(key string, value string) string {
	return h.http(fmt.Sprintf("data/%s", key), "POST", value)
}

func (h *HyperDB) Delete(key string) bool {
	return h.http(fmt.Sprintf("data/%s", key), "DELETE", "") == R_OK
}

func (h *HyperDB) All() string {
	return h.http("data", "", "")
}

func (h *HyperDB) Clear() bool {
	return h.http("data", "DELETE", "") == R_OK
}

func (h *HyperDB) Empty() bool {
	return h.http("empty", "", "") == R_TRUE
}

func (h *HyperDB) Save() bool {
	return h.http("save", "POST", "") == R_OK
}

func (h *HyperDB) Reload() bool {
	return h.http("reload", "POST", "") == R_OK
}

func (h *HyperDB) Reset() bool {
	return h.http("reset", "DELETE", "") == R_OK
}
