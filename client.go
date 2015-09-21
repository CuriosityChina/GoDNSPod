package dnspod

import (
	"time"

	"github.com/h2object/rpc"
	// "log"
)

type Cache interface {
	Set(string, interface{}, time.Duration)
	Add(string, interface{}, time.Duration) error
	Get(string) (interface{}, bool)
	Delete(string)
	DeleteExpired()
}

type Status struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	CreateAt string `json:"created_at"`
}

const DNSPODTOKEN = "DNSPODxTOKENx"

type DNSPodClient struct {
	loginToken string
	addr       string
	conn       *rpc.Client
}

func NewDNSPodClient(addr string, apiToken string) *DNSPodClient {
	connection := rpc.NewClient(NewDNSPODAnalyzer())
	return &DNSPodClient{
		loginToken: apiToken,
		addr:       addr,
		conn:       connection,
	}
}
