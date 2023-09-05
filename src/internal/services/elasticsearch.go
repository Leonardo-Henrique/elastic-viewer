package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/Leonardo-Henrique/elastic-viewer/src/models"
)

type ec models.ESClient

func MakeESConn(esCredentials models.ESCredentials) *ec {
	return &ec{
		ESCredentials: esCredentials,
		ESClient: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (e *ec) PingESConn() error {
	req, err := http.NewRequest(http.MethodGet, e.ESHost, nil)
	if err != nil {
		return err
	}

	auth := e.ESUsername + ":" + e.ESPassword

	encodedCred := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Add(
		"Authorization",
		"Basic "+encodedCred,
	)

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
