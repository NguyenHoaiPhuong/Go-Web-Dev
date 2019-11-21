package config

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"math"
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
		log.Errorln(err)
	}

	r.URL = u
	r.Header.Set("X-Forwarded-Host", r.Host)
	r.Header.Set("Origin", proxy.URL())
	r.Host = server.URL()
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
		log.Errorln()
		return 0, errors.New("connection refused:" + err.Error())
	}
	log.Infoln("Received response: " + strconv.Itoa(resp.StatusCode))

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("Proxy: Failed to read response body")
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

func (proxy Proxy) attemptServers(w http.ResponseWriter, r *http.Request, ignoreList []string) {
	if float64(len(ignoreList)) >= math.Min(float64(3), float64(len(proxy.Servers))) {
		log.Errorln("Failed to find server for request")
		http.NotFound(w, r)
		return
	}

	server := proxy.ChooseServer(ignoreList)
	log.Infoln("Got request: " + r.RequestURI)
	log.Infoln("Sending to server: " + server.Name)

	server.Connections++
	_, err := proxy.ReverseProxy(w, r, *server)
	server.Connections--

	if err != nil && strings.Contains(err.Error(), "connection refused") {
		log.Warnln("Server did not respond: " + server.Name)

		proxy.attemptServers(w, r, append(ignoreList, server.Name))
		return
	}

	log.Infoln("Responded to request successfuly")
}

// Handler func
func (proxy Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	proxy.attemptServers(w, r, []string{})
}
