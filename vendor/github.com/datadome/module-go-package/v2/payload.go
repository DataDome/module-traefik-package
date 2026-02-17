package modulego

import (
	"net/url"
	"strings"
)

// OperationType describes the expected operations values for a GraphQL query.
type OperationType string

const (
	Mutation     OperationType = "mutation"
	Query        OperationType = "query"
	Subscription OperationType = "subscription"
)

// GraphQLData describes the informations extracted from the GraphQL query
type GraphQLData struct {
	Type  OperationType
	Name  string
	Count int
}

// ProtectionAPIRequestPayload is used to construct the payload that will be send to the Protection API
type ProtectionAPIRequestPayload struct {
	Key                    string
	RequestModuleName      string
	ModuleVersion          string
	ServerName             string
	APIConnectionState     string
	IP                     string
	Port                   string
	TimeRequest            string
	Protocol               string
	Method                 string
	ServerHostName         string
	Request                string
	HeadersList            string
	Host                   string
	UserAgent              string
	Referer                string
	Accept                 string
	AcceptEncoding         string
	AcceptLanguage         string
	AcceptCharset          string
	Origin                 string
	XForwardedForIP        string
	XRequestedWith         string
	Connection             string
	Pragma                 string
	CacheControl           string
	CookiesLen             string
	CookiesList            string
	AuthorizationLen       string
	PostParamLen           string
	XRealIP                string
	ClientID               string
	SecChDeviceMemory      string
	SecChUA                string
	SecChUAArch            string
	SecChUAFullVersionList string
	SecChUAMobile          string
	SecChUAModel           string
	SecChUAPlatform        string
	SecFetchDest           string
	SecFetchMode           string
	SecFetchSite           string
	SecFetchUser           string
	Via                    string
	From                   string
	ContentType            string
	TrueClientIP           string
	GraphQLOperationCount  string
	GraphQLOperationName   string
	GraphQLOperationType   OperationType
}

func (p *ProtectionAPIRequestPayload) Encode() string {
	var parts []string

	// Helper function to add key-value pairs
	addParam := func(key, value string) {
		if value != "" {
			parts = append(parts, url.QueryEscape(key)+"="+url.QueryEscape(value))
		}
	}

	// Required fields - always included first
	addParam("Key", p.Key)
	addParam("IP", p.IP)
	addParam("Request", p.Request)
	addParam("RequestModuleName", p.RequestModuleName)

	// Optional fields - only added if not empty, in order
	addParam("ModuleVersion", p.ModuleVersion)
	addParam("ServerName", p.ServerName)
	addParam("APIConnectionState", p.APIConnectionState)
	addParam("Port", p.Port)
	addParam("TimeRequest", p.TimeRequest)
	addParam("Protocol", p.Protocol)
	addParam("Method", p.Method)
	addParam("ServerHostname", p.ServerHostName)
	addParam("HeadersList", p.HeadersList)
	addParam("Host", p.Host)
	addParam("UserAgent", p.UserAgent)
	addParam("Referer", p.Referer)
	addParam("Accept", p.Accept)
	addParam("AcceptEncoding", p.AcceptEncoding)
	addParam("AcceptLanguage", p.AcceptLanguage)
	addParam("AcceptCharset", p.AcceptCharset)
	addParam("Origin", p.Origin)
	addParam("XForwardedForIP", p.XForwardedForIP)
	addParam("X-Requested-With", p.XRequestedWith)
	addParam("Connection", p.Connection)
	addParam("Pragma", p.Pragma)
	addParam("CacheControl", p.CacheControl)
	addParam("CookiesLen", p.CookiesLen)
	addParam("CookiesList", p.CookiesList)
	addParam("AuthorizationLen", p.AuthorizationLen)
	addParam("PostParamLen", p.PostParamLen)
	addParam("X-Real-IP", p.XRealIP)
	addParam("ClientID", p.ClientID)
	addParam("SecCHDeviceMemory", p.SecChDeviceMemory)
	addParam("SecCHUA", p.SecChUA)
	addParam("SecCHUAArch", p.SecChUAArch)
	addParam("SecCHUAFullVersionList", p.SecChUAFullVersionList)
	addParam("SecCHUAMobile", p.SecChUAMobile)
	addParam("SecCHUAModel", p.SecChUAModel)
	addParam("SecCHUAPlatform", p.SecChUAPlatform)
	addParam("SecFetchDest", p.SecFetchDest)
	addParam("SecFetchMode", p.SecFetchMode)
	addParam("SecFetchSite", p.SecFetchSite)
	addParam("SecFetchUser", p.SecFetchUser)
	addParam("Via", p.Via)
	addParam("From", p.From)
	addParam("ContentType", p.ContentType)
	addParam("TrueClientIP", p.TrueClientIP)
	addParam("GraphQLOperationCount", p.GraphQLOperationCount)
	addParam("GraphQLOperationName", p.GraphQLOperationName)
	addParam("GraphQLOperationType", string(p.GraphQLOperationType))

	return strings.Join(parts, "&")
}
