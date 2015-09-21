# GoDNSPod Golang Client Library
--------------------------------

DNSPOD API 简易客户端封装包。

DNSPod API client library for golang based on DNSPod.com API service


## Quick Start

````
import "github.com/lowstz/GoDNSPod"

apiToken := "ID,API_TOKEN"
client := dnspod.NewDNSPodClient("dnsapi.com", apiToken)

// domain info variable
var domain dnspod.DomainInfo

if err := client.GetDomainInfo("your.domain", &domain); err != nil {
	fmt.Println(err)
}

// domain record
var rr dnspod.RecordInfo

// configure record
rr.Name = "subdomain"
rr.Type = "A"
rr.Line = "default"
rr.Value = "x.x.x.x"

if err := client.AddRecord(domain.ID, &rr); err != nil {
	fmt.Println(err)
}

````

## Shit

我只想说 DNSPOD 提供的 APIs 很烂!!! 操作极度不便! 修改一条记录, 必须获取整个记录列表竟然没有单获取接口。
