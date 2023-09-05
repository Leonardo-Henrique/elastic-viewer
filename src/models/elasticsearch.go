package models

import (
	"fmt"
	"net/http"
)

// ElasticsearchCredentials defines the structure of Elasticsearch Sign In information
type ESCredentials struct {
	//ElasticsearchHost main host to connect, choose a non-master host if possible
	ESHost string

	// ElasticsearchUsername is the username that will be used to be logged into Elasticsearch
	ESUsername string

	// ElasticsearchPassword is the user's password that will be used to be logged into Elasticsearch
	ESPassword string
}

// ElasticsearcEndpoint defines the endpoint structure to query Elasticsearch
type ESEndpoint struct {
	ESBaseURL   string
	ESPort      string
	Method      string
	QueryParams []interface{}
}

// ElasticsearchClient represents a instance of a Elasticsearch client connection
type ESClient struct {
	ESClient *http.Client
	ESCredentials
}

type ESError struct {
	Message string
}

func (er *ESError) Error() string {
	return fmt.Sprintf(er.Message)
}
