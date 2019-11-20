package config

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/log"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/utils"
)

// Proxy structure
type Proxy struct {
	Host    string
	Port    string
	Scheme  string
	Servers []Server
}

// URL returns server url
func (proxy Proxy) URL() string {
	return proxy.Scheme + "://" + proxy.Host + ":" + proxy.Port
}

// ChooseServer select server from Servers list and not in ignoreList
// TODO: This crashes if we define no servers in our config
func (proxy Proxy) ChooseServer(ignoreList []string) *Server {
	min := -1
	minIndex := 0
	for idx, server := range proxy.Servers {
		if utils.SliceContainsString(ignoreList, server.Name) {
			continue
		}

		conn := server.Connections
		if min == -1 || conn < min {
			min = conn
		}
		minIndex = idx
	}

	return &proxy.Servers[minIndex]
}

// ReverseProxy func
func (proxy Proxy) ReverseProxy(w http.ResponseWriter, r *http.Request, server Server) (int, error) {
	u, err := url.Parse(server.URL() + r.RequestURI)
	if err != nil {
		log.Errorln(err.Error())
	}

	r.URL = u
	r.Header.Set("X-Forwarded-Host", r.Host)
	r.Header.Set("Origin", proxy.origin())
	r.Host = server.Url()
	r.RequestURI = ""

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// TODO: If the server doesn't respond, try a new web server
	// We could return a status code from this function and let the handler try passing the request to a new server.
	resp, err := client.Do(r)
	if err != nil {
		// For now, this is a fatal error
		// When we can fail to another webserver, this should only be a warning.
		LogErr("connection refused")
		return 0, err
	}
	LogInfo("Recieved response: " + strconv.Itoa(resp.StatusCode))

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErr("Proxy: Failed to read response body")
		http.NotFound(w, r)
		return 0, err
	}

	buffer := bytes.NewBuffer(bodyBytes)
	for k, v := range resp.Header {
		w.Header().Set(k, strings.Join(v, ";"))
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, buffer)
	return resp.StatusCode, nil
}
