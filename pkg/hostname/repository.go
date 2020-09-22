package hostname

//import "github.com/activecm/rita/pkg/data"

// Repository for hostnames collection
type Repository interface {
	CreateIndexes() error
	Upsert(domainMap map[string]*Input)
}

//update ....
type update struct {
	selector interface{}
	query    interface{}
}

//Input ....
type Input struct {
	Host        string   //A hostname
	ResolvedIPs []string //Resolved IPs associated with a given hostname
	ClientIPs   []string //DNS Client IPs which issued queries for a given hostname
}

//TODO[AGENT]: Use UniqueIP/ NetworkID info in hostname ips/ clientIPs

//AnalysisView (for reporting)
type AnalysisView struct {
	Host              string   `bson:"host"`
	Connections       int      `bson:"conn_count"`
	UniqueConnections int      `bson:"uconn_count"`
	TotalBytes        int      `bson:"total_bytes"`
	ConnectedHosts    []string `bson:"ips,omitempty"`
}
