package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Leonardo-Henrique/elastic-viewer/src/models"
)

type esClient models.ESClient

func MakeESConn(esCredentials models.ESCredentials) *esClient {
	return &esClient{
		ESCredentials: esCredentials,
		ESClient: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func formatBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func addBasicAuthIntoHeader(r *http.Request, username, password string) {
	r.Header.Add(
		"Authorization",
		"Basic "+formatBasicAuth(username, password),
	)
}

func (e *esClient) PingESConn() error {
	req, err := http.NewRequest(http.MethodGet, e.ESHost, nil)
	if err != nil {
		return err
	}

	addBasicAuthIntoHeader(req, e.ESUsername, e.ESPassword)

	resp, err := e.ESClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	statusCodeReceived := resp.StatusCode

	if statusCodeReceived != http.StatusOK {

		return &models.ESError{
			Message: (fmt.Sprintf("Failed to Ping Elasticsearch at host %s with statuscode %d", e.ESHost, statusCodeReceived)),
		}

	}

	return nil
}

func (e *esClient) GetClusterStats() (models.ClusterStats, error) {
	var clusterStats models.ClusterStats
	req, err := http.NewRequest(http.MethodGet, e.ESHost+"/_cluster/stats", nil)
	if err != nil {
		return clusterStats, err
	}

	addBasicAuthIntoHeader(req, e.ESUsername, e.ESPassword)

	resp, err := e.ESClient.Do(req)
	if err != nil {
		return clusterStats, err
	}
	defer resp.Body.Close()

	statusCodeReceived := resp.StatusCode

	if statusCodeReceived != http.StatusOK {

		return clusterStats, &models.ESError{
			Message: (fmt.Sprintf("Failed to Ping Elasticsearch at host %s with statuscode %d", e.ESHost, statusCodeReceived)),
		}

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clusterStats, err
	}

	if err := json.Unmarshal(body, &clusterStats); err != nil {
		return clusterStats, err
	}

	return clusterStats, nil

}
