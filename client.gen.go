// Package telnyx provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package telnyx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for CallEventRecordType.
const (
	CallEventRecordTypeCallEvent CallEventRecordType = "call_event"
)

// Defines values for CallEventType.
const (
	CallEventTypeCommand CallEventType = "command"
	CallEventTypeWebhook CallEventType = "webhook"
)

// Defines values for FilterCallEventStatus.
const (
	FilterCallEventStatusDelivered FilterCallEventStatus = "delivered"
	FilterCallEventStatusFailed    FilterCallEventStatus = "failed"
)

// Defines values for FilterCallEventType.
const (
	FilterCallEventTypeCommand FilterCallEventType = "command"
	FilterCallEventTypeWebhook FilterCallEventType = "webhook"
)

// Defines values for ListCallEventsParamsFilterStatus.
const (
	ListCallEventsParamsFilterStatusDelivered ListCallEventsParamsFilterStatus = "delivered"
	ListCallEventsParamsFilterStatusFailed    ListCallEventsParamsFilterStatus = "failed"
)

// Defines values for ListCallEventsParamsFilterType.
const (
	Command ListCallEventsParamsFilterType = "command"
	Webhook ListCallEventsParamsFilterType = "webhook"
)

// CallEvent defines model for CallEvent.
type CallEvent struct {
	// CallLegId Uniquely identifies an individual call leg.
	CallLegId string `json:"call_leg_id"`

	// CallSessionId Uniquely identifies the call control session. A session may include multiple call leg events.
	CallSessionId string `json:"call_session_id"`

	// EventTimestamp Event timestamp
	EventTimestamp string `json:"event_timestamp"`

	// Metadata Event metadata, which includes raw event, and extra information based on event type
	Metadata map[string]interface{} `json:"metadata"`

	// Name Event name
	Name       string              `json:"name"`
	RecordType CallEventRecordType `json:"record_type"`

	// Type Event type
	Type CallEventType `json:"type"`
}

// CallEventRecordType defines model for CallEvent.RecordType.
type CallEventRecordType string

// CallEventType Event type
type CallEventType string

// Error defines model for Error.
type Error struct {
	Code   string                  `json:"code"`
	Detail *string                 `json:"detail,omitempty"`
	Meta   *map[string]interface{} `json:"meta,omitempty"`
	Source *struct {
		// Parameter Indicates which query parameter caused the error.
		Parameter *string `json:"parameter,omitempty"`

		// Pointer JSON pointer (RFC6901) to the offending entity.
		Pointer *string `json:"pointer,omitempty"`
	} `json:"source,omitempty"`
	Title string `json:"title"`
}

// Errors defines model for Errors.
type Errors struct {
	Errors *[]Error `json:"errors,omitempty"`
}

// PaginationMeta defines model for PaginationMeta.
type PaginationMeta struct {
	PageNumber   *int `json:"page_number,omitempty"`
	PageSize     *int `json:"page_size,omitempty"`
	TotalPages   *int `json:"total_pages,omitempty"`
	TotalResults *int `json:"total_results,omitempty"`
}

// FilterCallEventStatus defines model for FilterCallEventStatus.
type FilterCallEventStatus string

// FilterCallEventTimestampEqualTo defines model for FilterCallEventTimestampEqualTo.
type FilterCallEventTimestampEqualTo = string

// FilterCallEventTimestampGreaterThan defines model for FilterCallEventTimestampGreaterThan.
type FilterCallEventTimestampGreaterThan = string

// FilterCallEventTimestampGreaterThanOrEqualTo defines model for FilterCallEventTimestampGreaterThanOrEqualTo.
type FilterCallEventTimestampGreaterThanOrEqualTo = string

// FilterCallEventTimestampLessThan defines model for FilterCallEventTimestampLessThan.
type FilterCallEventTimestampLessThan = string

// FilterCallEventTimestampLessThanOrEqualTo defines model for FilterCallEventTimestampLessThanOrEqualTo.
type FilterCallEventTimestampLessThanOrEqualTo = string

// FilterCallEventType defines model for FilterCallEventType.
type FilterCallEventType string

// FilterCallLegId defines model for FilterCallLegId.
type FilterCallLegId = openapi_types.UUID

// FilterCallSessionId defines model for FilterCallSessionId.
type FilterCallSessionId = openapi_types.UUID

// PageNumber defines model for PageNumber.
type PageNumber = int

// PageSize defines model for PageSize.
type PageSize = int

// GenericErrorResponse defines model for GenericErrorResponse.
type GenericErrorResponse = Errors

// ListCallEventsResponse defines model for ListCallEventsResponse.
type ListCallEventsResponse struct {
	Data *[]CallEvent    `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// ListCallEventsParams defines parameters for ListCallEvents.
type ListCallEventsParams struct {
	// FilterCallLegId The unique identifier of an individual call leg.
	FilterCallLegId *FilterCallLegId `form:"filter[call_leg_id],omitempty" json:"filter[call_leg_id],omitempty"`

	// FilterCallSessionId The unique identifier of the call control session. A session may include multiple call leg events.
	FilterCallSessionId *FilterCallSessionId `form:"filter[call_session_id],omitempty" json:"filter[call_session_id],omitempty"`

	// FilterStatus Event status
	FilterStatus *ListCallEventsParamsFilterStatus `form:"filter[status],omitempty" json:"filter[status],omitempty"`

	// FilterType Event type
	FilterType *ListCallEventsParamsFilterType `form:"filter[type],omitempty" json:"filter[type],omitempty"`

	// FilterEventTimestampGt Event timestamp: greater than
	FilterEventTimestampGt *FilterCallEventTimestampGreaterThan `form:"filter[event_timestamp][gt],omitempty" json:"filter[event_timestamp][gt],omitempty"`

	// FilterEventTimestampGte Event timestamp: greater than or equal
	FilterEventTimestampGte *FilterCallEventTimestampGreaterThanOrEqualTo `form:"filter[event_timestamp][gte],omitempty" json:"filter[event_timestamp][gte],omitempty"`

	// FilterEventTimestampLt Event timestamp: lower than
	FilterEventTimestampLt *FilterCallEventTimestampLessThan `form:"filter[event_timestamp][lt],omitempty" json:"filter[event_timestamp][lt],omitempty"`

	// FilterEventTimestampLte Event timestamp: lower than or equal
	FilterEventTimestampLte *FilterCallEventTimestampLessThanOrEqualTo `form:"filter[event_timestamp][lte],omitempty" json:"filter[event_timestamp][lte],omitempty"`

	// FilterEventTimestampEq Event timestamp: equal
	FilterEventTimestampEq *FilterCallEventTimestampEqualTo `form:"filter[event_timestamp][eq],omitempty" json:"filter[event_timestamp][eq],omitempty"`

	// PageNumber The page number to load
	PageNumber *PageNumber `form:"page[number],omitempty" json:"page[number],omitempty"`

	// PageSize The size of the page
	PageSize *PageSize `form:"page[size],omitempty" json:"page[size],omitempty"`
}

// ListCallEventsParamsFilterStatus defines parameters for ListCallEvents.
type ListCallEventsParamsFilterStatus string

// ListCallEventsParamsFilterType defines parameters for ListCallEvents.
type ListCallEventsParamsFilterType string

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListCallEvents request
	ListCallEvents(ctx context.Context, params *ListCallEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListCallEvents(ctx context.Context, params *ListCallEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListCallEventsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListCallEventsRequest generates requests for ListCallEvents
func NewListCallEventsRequest(server string, params *ListCallEventsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/call_events")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.FilterCallLegId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[call_leg_id]", runtime.ParamLocationQuery, *params.FilterCallLegId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterCallSessionId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[call_session_id]", runtime.ParamLocationQuery, *params.FilterCallSessionId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterStatus != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[status]", runtime.ParamLocationQuery, *params.FilterStatus); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterType != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[type]", runtime.ParamLocationQuery, *params.FilterType); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterEventTimestampGt != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[event_timestamp][gt]", runtime.ParamLocationQuery, *params.FilterEventTimestampGt); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterEventTimestampGte != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[event_timestamp][gte]", runtime.ParamLocationQuery, *params.FilterEventTimestampGte); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterEventTimestampLt != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[event_timestamp][lt]", runtime.ParamLocationQuery, *params.FilterEventTimestampLt); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterEventTimestampLte != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[event_timestamp][lte]", runtime.ParamLocationQuery, *params.FilterEventTimestampLte); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FilterEventTimestampEq != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter[event_timestamp][eq]", runtime.ParamLocationQuery, *params.FilterEventTimestampEq); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PageNumber != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page[number]", runtime.ParamLocationQuery, *params.PageNumber); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PageSize != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page[size]", runtime.ParamLocationQuery, *params.PageSize); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListCallEventsWithResponse request
	ListCallEventsWithResponse(ctx context.Context, params *ListCallEventsParams, reqEditors ...RequestEditorFn) (*ListCallEventsResult, error)
}

type ListCallEventsResult struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListCallEventsResponse
	JSONDefault  *GenericErrorResponse
}

// Status returns HTTPResponse.Status
func (r ListCallEventsResult) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListCallEventsResult) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListCallEventsWithResponse request returning *ListCallEventsResult
func (c *ClientWithResponses) ListCallEventsWithResponse(ctx context.Context, params *ListCallEventsParams, reqEditors ...RequestEditorFn) (*ListCallEventsResult, error) {
	rsp, err := c.ListCallEvents(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListCallEventsResult(rsp)
}

// ParseListCallEventsResult parses an HTTP response from a ListCallEventsWithResponse call
func ParseListCallEventsResult(rsp *http.Response) (*ListCallEventsResult, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListCallEventsResult{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListCallEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GenericErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
