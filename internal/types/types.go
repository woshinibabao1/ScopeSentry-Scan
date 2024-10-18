// Package types -----------------------------
// @file      : type.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2023/12/11 10:05
// -------------------------------------------
package types

import (
	"encoding/json"
	"github.com/projectdiscovery/katana/pkg/navigation"
	"github.com/projectdiscovery/tlsx/pkg/tlsx/clients"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"
)

type SubdomainResult struct {
	Host       string
	Type       string
	Value      []string
	IP         []string
	Time       string
	Project    string
	TaskName   string `bson:"taskName"`
	RootDomain string `bson:"rootDomain"`
}

type AssetHttp struct {
	Time          string                 `bson:"time" csv:"time"`
	LastScanTime  string                 `bson:"lastScanTime"`
	TLSData       *clients.Response      `bson:"tls" csv:"tls"`
	Hashes        map[string]interface{} `bson:"hash" csv:"hash"`
	CDNName       string                 `bson:"cdn_name" csv:"cdn_name"`
	Port          string                 `bson:"port" csv:"port"`
	URL           string                 `bson:"url" csv:"url"`
	Title         string                 `bson:"title" csv:"title"`
	Type          string                 `bson:"Type" csv:"Type"`
	Error         string                 `bson:"error" csv:"error"`
	ResponseBody  string                 `bson:"body" csv:"body"`
	Host          string                 `bson:"host" csv:"host"`
	IP            string                 `bson:"ip"`
	FavIconMMH3   string                 `bson:"favicon" csv:"favicon"`
	FaviconPath   string                 `bson:"favicon_path" csv:"favicon_path"`
	RawHeaders    string                 `bson:"raw_header" csv:"raw_header"`
	Jarm          string                 `bson:"jarm" csv:"jarm"`
	Technologies  []string               `bson:"tech" csv:"tech"`
	StatusCode    int                    `bson:"status_code" csv:"status_code"`
	ContentLength int                    `bson:"content_length" csv:"content_length"`
	CDN           bool                   `bson:"cdn" csv:"cdn"`
	Webcheck      bool                   `bson:"webcheck" csv:"webcheck"`
	Project       string                 `bson:"project" csv:"project"`
	WebFinger     []string               `bson:"web_finger" csv:"web_finger"`
	IconContent   string                 `bson:"iconContent"`
	Domain        string                 `bson:"domain"`
	TaskName      string                 `bson:"taskName"`
	WebServer     string                 `bson:"webServer"`
	Service       string                 `bson:"service"`
	RootDomain    string                 `bson:"rootDomain"`
}

type PortAlive struct {
	Host string `bson:"host"`
	IP   string `bson:"ip"`
	Port string `bson:"port"`
}
type Project struct {
	ID     string   `bson:"id"`
	Target []string `bson:"target"`
}
type AssetOther struct {
	Time         string          `bson:"time" csv:"time"`
	LastScanTime string          `bson:"lastScanTime"`
	Host         string          `bson:"host"`
	IP           string          `bson:"ip"`
	Port         string          `bson:"port"`
	Service      string          `bson:"service"`
	TLS          bool            `bson:"tls"`
	Transport    string          `bson:"transport"`
	Version      string          `bson:"version"`
	Raw          json.RawMessage `bson:"metadata"`
	Project      string          `bson:"project"`
	Type         string
	TaskName     string `bson:"taskName"`
	RootDomain   string `bson:"rootDomain"`
}

type ChangeLog struct {
	FieldName string `json:"fieldName"`
	Old       string `json:"old"`
	New       string `json:"new"`
}

type AssetChangeLog struct {
	AssetId   string `json:"assetId"`
	Timestamp string `json:"timestamp" csv:"timestamp"`
	Change    []ChangeLog
}

type UrlResult struct {
	Input      string `json:"input"`
	Source     string `json:"source"`
	OutputType string `json:"type"`
	Output     string `json:"output"`
	Status     int    `json:"status"`
	Length     int    `json:"length"`
	Time       string `json:"time"`
	Body       string `bson:"body"`
	Project    string
	TaskName   string `bson:"taskName"`
	ResultId   string
	RootDomain string
}

type SecretResults struct {
	Url   string
	Kind  string
	Key   string
	Value string
}

type CrawlerResult struct {
	Url        string
	Method     string
	Body       string
	Project    string
	TaskName   string `bson:"taskName"`
	ResultId   string
	RootDomain string
}

type PortDict struct {
	ID    string `bson:"id"`
	Value string `bson:"value"`
}

type SubdomainTakerFinger struct {
	Name     string
	Cname    []string
	Response []string
}

type SubTakeResult struct {
	Input      string
	Value      string
	Cname      string
	Response   string
	Project    string
	TaskName   string `bson:"taskName"`
	RootDomain string
}

type DirResult struct {
	Url        string
	Status     int
	Msg        string
	Project    string
	Length     int
	TaskName   string `bson:"taskName"`
	RootDomain string
}

type SensitiveRule struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	State   bool   `json:"enabled"`
	Regular string `json:"pattern"`
	Color   string `bson:"color"`
}

type SensitiveResult struct {
	Url        string
	SID        string
	Match      []string
	Project    string
	Color      string
	Time       string
	Md5        string
	TaskName   string `bson:"taskName"`
	RootDomain string
}

type VulnResult struct {
	Url      string
	VulnId   string
	VulName  string
	Matched  string
	Project  string
	Level    string
	Time     string
	Request  string
	Response string
	TaskName string `bson:"taskName"`
}
type TmpPageMonitResult struct {
	Url      string
	Content  string
	TaskName string `bson:"taskName"`
}
type PageMonitResult struct {
	ID       primitive.ObjectID `bson:"_id"`
	Url      string             `bson:"url"`
	Content  []string           `bson:"content"`
	Hash     []string           `bson:"hash"`
	Diff     []string           `bson:"diff"`
	State    int                `bson:"state"`
	Project  string             `bson:"project"`
	Time     string             `bson:"time"`
	TaskName string             `bson:"taskName"`
}
type WebFinger struct {
	ID      string
	Express []string
	Name    string
	State   bool
}

type HttpSample struct {
	Url        string
	StatusCode int
	Body       string
	Msg        string
}

type HttpResponse struct {
	Url           string
	StatusCode    int
	Body          string
	ContentLength int
	Redirect      string
	Title         string
}

type NotificationConfig struct {
	SubdomainScan                 bool `bson:"subdomainScan"`
	DirScanNotification           bool `bson:"dirScanNotification"`
	PortScanNotification          bool `bson:"portScanNotification"`
	SensitiveNotification         bool `bson:"sensitiveNotification"`
	SubdomainNotification         bool `bson:"subdomainNotification"`
	SubdomainTakeoverNotification bool `bson:"subdomainTakeoverNotification"`
	PageMonNotification           bool `bson:"pageMonNotification"`
	VulNotification               bool `bson:"vulNotification"`
}

type NotificationApi struct {
	Url         string `bson:"url"`
	Method      string `bson:"method"`
	ContentType string `bson:"contentType"`
	Data        string `bson:"data"`
	State       bool   `bson:"state"`
}

type PocData struct {
	Name  string
	Level string
}

type CrawlerTask struct {
	Target []string
	Host   string
	Id     string
	Wg     *sync.WaitGroup
}

type DomainSkip struct {
	Domain string
	Skip   bool
	IP     []string
}
type DomainResolve struct {
	Domain string
	IP     []string
}

type KatanaResult struct {
	Timestamp        time.Time
	Request          *navigation.Request
	Response         *navigation.Response
	PassiveReference *navigation.PassiveReference
	Error            string
}
