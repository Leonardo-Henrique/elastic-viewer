package models

// ClusterStats represents the response from _cluster/stats endpoint
type ClusterStats struct {

	// Nodes brings high level of health of ES Instances
	Nodes struct {
		Total      int16 `json:"total"`
		Successful int16 `json:"successful"`
		Failed     int16 `json:"failed"`
	} `json:"_nodes"`

	ClusterName string `json:"cluster_name"`
	ClusterUUID string `json:"cluster_uuid"`
	Timestamp   int64  `json:"timestamp"`
	Status      string `json:"status"`

	Indices struct {
		Count  int `json:"count"`
		Shards struct {
			Total       int16       `json:"total"`
			Primaries   int16       `json:"primaries"`
			Replication interface{} `json:"replication"`
		} `json:"shards"`

		Index struct {
			Shards struct {
				Min int16 `json:"min"`
				Max int16 `json:"max"`
				Avg int16 `json:"avg"`
			} `json:"shards"`

			Primaries struct {
				Min int16 `json:"min"`
				Max int16 `json:"max"`
				Avg int16 `json:"avg"`
			} `json:"primaries"`
		} `json:"index"`
	} `json:"indices"`
}
