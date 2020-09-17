package parsetypes

import (
	"github.com/activecm/rita/config"
	"github.com/globalsign/mgo/bson"
)

// Conn provides a data structure for bro's connection data
type Conn struct {
	// ID is the id coming out of mongodb
	ID bson.ObjectId `bson:"_id,omitempty"`
	// TimeStamp of this connection
	TimeStamp int64 `bson:"ts" bro:"ts" brotype:"time" json:"-"`
	// TimeStampGeneric is used when reading from json files
	TimeStampGeneric interface{} `bson:"-" json:"ts"`
	// UID is the Unique Id for this connection (generated by Bro)
	UID string `bson:"uid" bro:"uid" brotype:"string" json:"uid"`
	// Source is the source address for this connection
	Source string `bson:"id_orig_h" bro:"id.orig_h" brotype:"addr" json:"id.orig_h"`
	// SourcePort is the source port of this connection
	SourcePort int `bson:"id_orig_p" bro:"id.orig_p" brotype:"port" json:"id.orig_p"`
	// Destination is the destination of the connection
	Destination string `bson:"id_resp_h" bro:"id.resp_h" brotype:"addr" json:"id.resp_h"`
	// DestinationPort is the port at the destination host
	DestinationPort int `bson:"id_resp_p" bro:"id.resp_p" brotype:"port" json:"id.resp_p"`

	// TODO[AGENT]: Add AgentUUID string
	// TODO[AGENT]: Add AgentHostname string

	// Proto is the string protocol identifier for this connection
	Proto string `bson:"proto" bro:"proto" brotype:"enum" json:"proto"`
	// Service describes the service of this connection if there was one
	Service string `bson:"service" bro:"service" brotype:"string" json:"service"`
	// Duration is the floating point representation of connection length
	Duration float64 `bson:"duration" bro:"duration" brotype:"interval" json:"duration"`
	// OrigBytes is the byte count coming from the origin
	OrigBytes int64 `bson:"orig_bytes" bro:"orig_bytes" brotype:"count" json:"orig_bytes"`
	// RespBytes is the byte count coming in on response
	RespBytes int64 `bson:"resp_bytes" bro:"resp_bytes" brotype:"count" json:"resp_bytes"`
	// ConnState has data describing the state of a connection
	ConnState string `bson:"conn_state" bro:"conn_state" brotype:"string" json:"conn_state"`
	// LocalOrigin denotes that the connection originated locally
	LocalOrigin bool `bson:"local_orig" bro:"local_orig" brotype:"bool" json:"local_orig"`
	// LocalResponse denote that the connection responded locally
	LocalResponse bool `bson:"local_resp" bro:"local_resp" brotype:"bool" json:"local_resp"`
	// MissedBytes keeps a count of bytes missed
	MissedBytes int64 `bson:"missed_bytes" bro:"missed_bytes" brotype:"count" json:"missed_bytes"`
	// History is a string containing historical information
	History string `bson:"history"  bro:"history" brotype:"string" json:"history"`
	// OrigPkts is a count of origin packets
	OrigPkts int64 `bson:"orig_pkts"  bro:"orig_pkts" brotype:"count" json:"orig_pkts"`
	// OrigIpBytes is another origin data count
	OrigIPBytes int64 `bson:"orig_ip_bytes" bro:"orig_ip_bytes" brotype:"count" json:"orig_ip_bytes"`
	// RespPkts counts response packets
	RespPkts int64 `bson:"resp_pkts" bro:"resp_pkts" brotype:"count" json:"resp_pkts"`
	// RespIpBytes gives the bytecount of response data
	RespIPBytes int64 `bson:"resp_ip_bytes" bro:"resp_ip_bytes" brotype:"count" json:"resp_ip_bytes"`
	// TunnelParents lists tunnel parents
	TunnelParents []string `bson:"tunnel_parents" bro:"tunnel_parents" brotype:"set[string]" json:"tunnel_parents"`
}

//TargetCollection returns the mongo collection this entry should be inserted
func (line *Conn) TargetCollection(config *config.StructureTableCfg) string {
	return config.ConnTable
}

//Indices gives MongoDB indices that should be used with the collection
func (line *Conn) Indices() []string {
	return []string{"$hashed:id_orig_h", "$hashed:id_resp_h", "-duration", "ts", "uid"}
}

//ConvertFromJSON performs any extra conversions necessary when reading from JSON
func (line *Conn) ConvertFromJSON() {
	line.TimeStamp = convertTimestamp(line.TimeStampGeneric)
}
