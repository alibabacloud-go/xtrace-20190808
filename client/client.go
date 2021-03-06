// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTagKeyRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetTagKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyRequest) GoString() string {
	return s.String()
}

func (s *GetTagKeyRequest) SetRegionId(v string) *GetTagKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagKeyRequest) SetServiceName(v string) *GetTagKeyRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagKeyRequest) SetSpanName(v string) *GetTagKeyRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagKeyRequest) SetStartTime(v int64) *GetTagKeyRequest {
	s.StartTime = &v
	return s
}

func (s *GetTagKeyRequest) SetEndTime(v int64) *GetTagKeyRequest {
	s.EndTime = &v
	return s
}

type GetTagKeyResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   *GetTagKeyResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s GetTagKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBody) SetRequestId(v string) *GetTagKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagKeyResponseBody) SetTagKeys(v *GetTagKeyResponseBodyTagKeys) *GetTagKeyResponseBody {
	s.TagKeys = v
	return s
}

type GetTagKeyResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s GetTagKeyResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBodyTagKeys) SetTagKey(v []*string) *GetTagKeyResponseBodyTagKeys {
	s.TagKey = v
	return s
}

type GetTagKeyResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTagKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponse) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponse) SetHeaders(v map[string]*string) *GetTagKeyResponse {
	s.Headers = v
	return s
}

func (s *GetTagKeyResponse) SetBody(v *GetTagKeyResponseBody) *GetTagKeyResponse {
	s.Body = v
	return s
}

type GetTagValRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	TagKey      *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetTagValRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagValRequest) GoString() string {
	return s.String()
}

func (s *GetTagValRequest) SetRegionId(v string) *GetTagValRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagValRequest) SetServiceName(v string) *GetTagValRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagValRequest) SetSpanName(v string) *GetTagValRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagValRequest) SetTagKey(v string) *GetTagValRequest {
	s.TagKey = &v
	return s
}

func (s *GetTagValRequest) SetStartTime(v int64) *GetTagValRequest {
	s.StartTime = &v
	return s
}

func (s *GetTagValRequest) SetEndTime(v int64) *GetTagValRequest {
	s.EndTime = &v
	return s
}

type GetTagValResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues *GetTagValResponseBodyTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Struct"`
}

func (s GetTagValResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBody) SetRequestId(v string) *GetTagValResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagValResponseBody) SetTagValues(v *GetTagValResponseBodyTagValues) *GetTagValResponseBody {
	s.TagValues = v
	return s
}

type GetTagValResponseBodyTagValues struct {
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s GetTagValResponseBodyTagValues) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBodyTagValues) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBodyTagValues) SetTagValue(v []*string) *GetTagValResponseBodyTagValues {
	s.TagValue = v
	return s
}

type GetTagValResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTagValResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagValResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponse) GoString() string {
	return s.String()
}

func (s *GetTagValResponse) SetHeaders(v map[string]*string) *GetTagValResponse {
	s.Headers = v
	return s
}

func (s *GetTagValResponse) SetBody(v *GetTagValResponseBody) *GetTagValResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	IsForce     *bool   `json:"IsForce,omitempty" xml:"IsForce,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetRegionId(v string) *GetTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetTokenRequest) SetAppType(v string) *GetTokenRequest {
	s.AppType = &v
	return s
}

func (s *GetTokenRequest) SetProxyUserId(v string) *GetTokenRequest {
	s.ProxyUserId = &v
	return s
}

func (s *GetTokenRequest) SetIsForce(v bool) *GetTokenRequest {
	s.IsForce = &v
	return s
}

type GetTokenResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *GetTokenResponseBodyToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v *GetTokenResponseBodyToken) *GetTokenResponseBody {
	s.Token = v
	return s
}

type GetTokenResponseBodyToken struct {
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	LicenseKey     *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	InternalDomain *string `json:"InternalDomain,omitempty" xml:"InternalDomain,omitempty"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetTokenResponseBodyToken) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyToken) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyToken) SetDomain(v string) *GetTokenResponseBodyToken {
	s.Domain = &v
	return s
}

func (s *GetTokenResponseBodyToken) SetLicenseKey(v string) *GetTokenResponseBodyToken {
	s.LicenseKey = &v
	return s
}

func (s *GetTokenResponseBodyToken) SetInternalDomain(v string) *GetTokenResponseBodyToken {
	s.InternalDomain = &v
	return s
}

func (s *GetTokenResponseBodyToken) SetPid(v string) *GetTokenResponseBodyToken {
	s.Pid = &v
	return s
}

type GetTokenResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

func (s *GetTraceRequest) SetAppType(v string) *GetTraceRequest {
	s.AppType = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

type GetTraceResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spans     *GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Struct"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v *GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

type GetTraceResponseBodySpans struct {
	Span []*GetTraceResponseBodySpansSpan `json:"Span,omitempty" xml:"Span,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) SetSpan(v []*GetTraceResponseBodySpansSpan) *GetTraceResponseBodySpans {
	s.Span = v
	return s
}

type GetTraceResponseBodySpansSpan struct {
	SpanId        *string                                    `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	OperationName *string                                    `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Timestamp     *int64                                     `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TagEntryList  *GetTraceResponseBodySpansSpanTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	LogEventList  *GetTraceResponseBodySpansSpanLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Struct"`
	HaveStack     *bool                                      `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	ServiceIp     *string                                    `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ParentSpanId  *string                                    `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	Duration      *int64                                     `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RpcId         *string                                    `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName   *string                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string                                    `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpansSpan) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpan) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpan) SetSpanId(v string) *GetTraceResponseBodySpansSpan {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetOperationName(v string) *GetTraceResponseBodySpansSpan {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetResultCode(v string) *GetTraceResponseBodySpansSpan {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTimestamp(v int64) *GetTraceResponseBodySpansSpan {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTagEntryList(v *GetTraceResponseBodySpansSpanTagEntryList) *GetTraceResponseBodySpansSpan {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetLogEventList(v *GetTraceResponseBodySpansSpanLogEventList) *GetTraceResponseBodySpansSpan {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetHaveStack(v bool) *GetTraceResponseBodySpansSpan {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetServiceIp(v string) *GetTraceResponseBodySpansSpan {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetParentSpanId(v string) *GetTraceResponseBodySpansSpan {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetDuration(v int64) *GetTraceResponseBodySpansSpan {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetRpcId(v string) *GetTraceResponseBodySpansSpan {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetServiceName(v string) *GetTraceResponseBodySpansSpan {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTraceID(v string) *GetTraceResponseBodySpansSpan {
	s.TraceID = &v
	return s
}

type GetTraceResponseBodySpansSpanTagEntryList struct {
	TagEntry []*GetTraceResponseBodySpansSpanTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanTagEntryList) SetTagEntry(v []*GetTraceResponseBodySpansSpanTagEntryListTagEntry) *GetTraceResponseBodySpansSpanTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodySpansSpanTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansSpanTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodySpansSpanTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansSpanTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodySpansSpanTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponseBodySpansSpanLogEventList struct {
	LogEvent []*GetTraceResponseBodySpansSpanLogEventListLogEvent `json:"LogEvent,omitempty" xml:"LogEvent,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventList) SetLogEvent(v []*GetTraceResponseBodySpansSpanLogEventListLogEvent) *GetTraceResponseBodySpansSpanLogEventList {
	s.LogEvent = v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEvent struct {
	TagEntryList *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	Timestamp    *int64                                                         `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEvent) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEvent) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEvent) SetTagEntryList(v *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) *GetTraceResponseBodySpansSpanLogEventListLogEvent {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEvent) SetTimestamp(v int64) *GetTraceResponseBodySpansSpanLogEventListLogEvent {
	s.Timestamp = &v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList struct {
	TagEntry []*GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) SetTagEntry(v []*GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type GetTraceAnalysisRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Api         *string `json:"Api,omitempty" xml:"Api,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s GetTraceAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAnalysisRequest) SetRegionId(v string) *GetTraceAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceAnalysisRequest) SetApi(v string) *GetTraceAnalysisRequest {
	s.Api = &v
	return s
}

func (s *GetTraceAnalysisRequest) SetQuery(v string) *GetTraceAnalysisRequest {
	s.Query = &v
	return s
}

func (s *GetTraceAnalysisRequest) SetProxyUserId(v string) *GetTraceAnalysisRequest {
	s.ProxyUserId = &v
	return s
}

type GetTraceAnalysisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetTraceAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAnalysisResponseBody) SetRequestId(v string) *GetTraceAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAnalysisResponseBody) SetData(v string) *GetTraceAnalysisResponseBody {
	s.Data = &v
	return s
}

type GetTraceAnalysisResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTraceAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAnalysisResponse) SetHeaders(v map[string]*string) *GetTraceAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAnalysisResponse) SetBody(v *GetTraceAnalysisResponseBody) *GetTraceAnalysisResponse {
	s.Body = v
	return s
}

type ListIpOrHostsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListIpOrHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsRequest) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsRequest) SetRegionId(v string) *ListIpOrHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpOrHostsRequest) SetServiceName(v string) *ListIpOrHostsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListIpOrHostsRequest) SetStartTime(v int64) *ListIpOrHostsRequest {
	s.StartTime = &v
	return s
}

func (s *ListIpOrHostsRequest) SetEndTime(v int64) *ListIpOrHostsRequest {
	s.EndTime = &v
	return s
}

type ListIpOrHostsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpNames   *ListIpOrHostsResponseBodyIpNames `json:"IpNames,omitempty" xml:"IpNames,omitempty" type:"Struct"`
}

func (s ListIpOrHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponseBody) SetRequestId(v string) *ListIpOrHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpOrHostsResponseBody) SetIpNames(v *ListIpOrHostsResponseBodyIpNames) *ListIpOrHostsResponseBody {
	s.IpNames = v
	return s
}

type ListIpOrHostsResponseBodyIpNames struct {
	IpName []*string `json:"IpName,omitempty" xml:"IpName,omitempty" type:"Repeated"`
}

func (s ListIpOrHostsResponseBodyIpNames) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponseBodyIpNames) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponseBodyIpNames) SetIpName(v []*string) *ListIpOrHostsResponseBodyIpNames {
	s.IpName = v
	return s
}

type ListIpOrHostsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIpOrHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIpOrHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponse) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponse) SetHeaders(v map[string]*string) *ListIpOrHostsResponse {
	s.Headers = v
	return s
}

func (s *ListIpOrHostsResponse) SetBody(v *ListIpOrHostsResponseBody) *ListIpOrHostsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetAppType(v string) *ListServicesRequest {
	s.AppType = &v
	return s
}

type ListServicesResponseBody struct {
	Services  *ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetServices(v *ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListServicesResponseBodyServices struct {
	Service []*ListServicesResponseBodyServicesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetService(v []*ListServicesResponseBodyServicesService) *ListServicesResponseBodyServices {
	s.Service = v
	return s
}

type ListServicesResponseBodyServicesService struct {
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListServicesResponseBodyServicesService) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesService) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesService) SetPid(v string) *ListServicesResponseBodyServicesService {
	s.Pid = &v
	return s
}

func (s *ListServicesResponseBodyServicesService) SetServiceName(v string) *ListServicesResponseBodyServicesService {
	s.ServiceName = &v
	return s
}

func (s *ListServicesResponseBodyServicesService) SetRegionId(v string) *ListServicesResponseBodyServicesService {
	s.RegionId = &v
	return s
}

type ListServicesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSpanNamesRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListSpanNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesRequest) GoString() string {
	return s.String()
}

func (s *ListSpanNamesRequest) SetRegionId(v string) *ListSpanNamesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSpanNamesRequest) SetServiceName(v string) *ListSpanNamesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListSpanNamesRequest) SetStartTime(v int64) *ListSpanNamesRequest {
	s.StartTime = &v
	return s
}

func (s *ListSpanNamesRequest) SetEndTime(v int64) *ListSpanNamesRequest {
	s.EndTime = &v
	return s
}

type ListSpanNamesResponseBody struct {
	SpanNames *ListSpanNamesResponseBodySpanNames `json:"SpanNames,omitempty" xml:"SpanNames,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSpanNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponseBody) SetSpanNames(v *ListSpanNamesResponseBodySpanNames) *ListSpanNamesResponseBody {
	s.SpanNames = v
	return s
}

func (s *ListSpanNamesResponseBody) SetRequestId(v string) *ListSpanNamesResponseBody {
	s.RequestId = &v
	return s
}

type ListSpanNamesResponseBodySpanNames struct {
	SpanName []*string `json:"SpanName,omitempty" xml:"SpanName,omitempty" type:"Repeated"`
}

func (s ListSpanNamesResponseBodySpanNames) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponseBodySpanNames) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponseBodySpanNames) SetSpanName(v []*string) *ListSpanNamesResponseBodySpanNames {
	s.SpanName = v
	return s
}

type ListSpanNamesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSpanNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpanNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponse) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponse) SetHeaders(v map[string]*string) *ListSpanNamesResponse {
	s.Headers = v
	return s
}

func (s *ListSpanNamesResponse) SetBody(v *ListSpanNamesResponseBody) *ListSpanNamesResponse {
	s.Body = v
	return s
}

type QueryMetricRequest struct {
	IntervalInSec *int32                       `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	StartTime     *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy       *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Limit         *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Metric        *string                      `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order         *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	ProxyUserId   *string                      `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Filters       []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	Dimensions    []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Measures      []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
}

func (s QueryMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) SetIntervalInSec(v int32) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int32) *QueryMetricRequest {
	s.Limit = &v
	return s
}

func (s *QueryMetricRequest) SetMetric(v string) *QueryMetricRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricRequest) SetOrder(v string) *QueryMetricRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
	return s
}

type QueryMetricRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricRequestFilters) SetKey(v string) *QueryMetricRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricRequestFilters) SetValue(v string) *QueryMetricRequestFilters {
	s.Value = &v
	return s
}

type QueryMetricResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricResponseBody) SetRequestId(v string) *QueryMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricResponseBody) SetData(v string) *QueryMetricResponseBody {
	s.Data = &v
	return s
}

type QueryMetricResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricResponse) SetHeaders(v map[string]*string) *QueryMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricResponse) SetBody(v *QueryMetricResponseBody) *QueryMetricResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	StartTime     *int64                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId      *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName   *string                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	OperationName *string                   `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	MinDuration   *int64                    `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	AppType       *string                   `json:"AppType,omitempty" xml:"AppType,omitempty"`
	PageNumber    *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse       *bool                     `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp     *string                   `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Tag           []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetAppType(v string) *SearchTracesRequest {
	s.AppType = &v
	return s
}

func (s *SearchTracesRequest) SetPageNumber(v int32) *SearchTracesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesRequest) SetPageSize(v int32) *SearchTracesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

type SearchTracesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

type SearchTracesResponseBody struct {
	PageBean  *SearchTracesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) SetPageBean(v *SearchTracesResponseBodyPageBean) *SearchTracesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

type SearchTracesResponseBodyPageBean struct {
	TraceInfos *SearchTracesResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Struct"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchTracesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBean) SetTraceInfos(v *SearchTracesResponseBodyPageBeanTraceInfos) *SearchTracesResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetPageSize(v int32) *SearchTracesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetTotalCount(v int64) *SearchTracesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchTracesResponseBodyPageBeanTraceInfos struct {
	TraceInfo []*SearchTracesResponseBodyPageBeanTraceInfosTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBodyPageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBeanTraceInfos) SetTraceInfo(v []*SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) *SearchTracesResponseBodyPageBeanTraceInfos {
	s.TraceInfo = v
	return s
}

type SearchTracesResponseBodyPageBeanTraceInfosTraceInfo struct {
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetOperationName(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetServiceIp(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetDuration(v int64) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetTimestamp(v int64) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetServiceName(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetTraceID(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("xtrace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagKeyWithOptions(request *GetTagKeyRequest, runtime *util.RuntimeOptions) (_result *GetTagKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTagKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTagKey"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagKey(request *GetTagKeyRequest) (_result *GetTagKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagKeyResponse{}
	_body, _err := client.GetTagKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagValWithOptions(request *GetTagValRequest, runtime *util.RuntimeOptions) (_result *GetTagValResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTagValResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTagVal"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagVal(request *GetTagValRequest) (_result *GetTagValResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagValResponse{}
	_body, _err := client.GetTagValWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetToken"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrace"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrace(request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceAnalysisWithOptions(request *GetTraceAnalysisRequest, runtime *util.RuntimeOptions) (_result *GetTraceAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTraceAnalysisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTraceAnalysis"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTraceAnalysis(request *GetTraceAnalysisRequest) (_result *GetTraceAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceAnalysisResponse{}
	_body, _err := client.GetTraceAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIpOrHostsWithOptions(request *ListIpOrHostsRequest, runtime *util.RuntimeOptions) (_result *ListIpOrHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIpOrHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIpOrHosts"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIpOrHosts(request *ListIpOrHostsRequest) (_result *ListIpOrHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpOrHostsResponse{}
	_body, _err := client.ListIpOrHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServices"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpanNamesWithOptions(request *ListSpanNamesRequest, runtime *util.RuntimeOptions) (_result *ListSpanNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSpanNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSpanNames"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpanNames(request *ListSpanNamesRequest) (_result *ListSpanNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpanNamesResponse{}
	_body, _err := client.ListSpanNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMetricWithOptions(request *QueryMetricRequest, runtime *util.RuntimeOptions) (_result *QueryMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMetric"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMetric(request *QueryMetricRequest) (_result *QueryMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricResponse{}
	_body, _err := client.QueryMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTraces"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraces(request *SearchTracesRequest) (_result *SearchTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesResponse{}
	_body, _err := client.SearchTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
