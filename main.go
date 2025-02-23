package goddit

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AccessToken struct {
	Token        string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RetrievedUTC int64  `json:"retrieved_utc"`
    AccessTokenFile string
}

type RedditCredentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientSecret string `json:"client_secret"`
	ClientId     string `json:"client_id"`
}

type Reddit struct {
	RedditCredentials
	AccessToken
	UserAgent         string
	requestsleft      int
	requestsCounted   int
	resetRequestCount int
}

func (a *AccessToken) ReadAccessTokenJson(jsonString []byte) {
	err := json.Unmarshal(jsonString, &a)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *AccessToken) ReadAccessTokenFile(filepath string) error {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	a.ReadAccessTokenJson(byteValue)
	return nil
}

func (a *AccessToken) WriteAccessTokenFile() {
	jsonFile, err := os.Create(a.AccessTokenFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonData, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	jsonFile.Write(jsonData)
}

func (r *RedditCredentials) ReadCredentialFile(credentialFile string) {
	jsonFile, err := os.Open(credentialFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &r)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Reddit) Init() {
	err := r.ReadAccessTokenFile(r.AccessTokenFile)
	if err == nil && !r.isExpired() {
		return
	}
	client := &http.Client{}
	parameters := fmt.Sprintf(`grant_type=password&username=%s&password=%s`, r.Username, r.Password)
	data := strings.NewReader(parameters)
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(r.ClientId, r.ClientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.RetrievedUTC = time.Now().Unix()
	r.ReadAccessTokenJson(bodyText)
	if err != nil {
		log.Fatal(err)
	}
    r.WriteAccessTokenFile()
}

// returns true if the access token is expired and false otherwise
func (a *AccessToken) isExpired() bool {
	if int64(a.ExpiresIn)+a.RetrievedUTC < time.Now().Unix() {
		return true
	}
	return false
}

func (r *Reddit) request(method, endpoint, parameters string) string {
	client := &http.Client{}
	data := strings.NewReader(parameters)
	url := fmt.Sprintf("https://oauth.reddit.com/%s", endpoint)
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Fatal(err)
	}

	if r.isExpired() {
		r.Init()
	}
	authorization := fmt.Sprintf("%s %s", r.TokenType, r.Token)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("User-Agent", r.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintf("%s", bodyText)
	return str
}

func (r *Reddit) NewAfter(subreddit, after string, count, limit int, show bool) string {
	parameters := fmt.Sprintf("after=%s&count=%d&limit=%d", after, count, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) NewBefore(subreddit, before string, count, limit int, show bool) string {
	parameters := fmt.Sprintf("before=%s&count=%d&limit=%d", before, count, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) New(subreddit string, limit int, show bool) string {
	parameters := fmt.Sprintf("limit=%d", limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) Info(id string) string {
    endpoint := fmt.Sprintf("api/info?id=%s", id)
	return r.request("GET", endpoint, "")
}

func (r *Reddit) new(subreddit, parameters string) string {
	endpoint := fmt.Sprintf("r/%s/new?%s", subreddit, parameters)
	return r.request("GET", endpoint, "")
}
