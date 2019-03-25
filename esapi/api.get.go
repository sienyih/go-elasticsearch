// Code generated from specification version 6.7.0: DO NOT EDIT

package esapi

import (
	"context"
	"strconv"
	"strings"
)

func newGetFunc(t Transport) Get {
	return func(index string, id string, o ...func(*GetRequest)) (*Response, error) {
		var r = GetRequest{Index: index, DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// Get returns a document.
//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html.
//
type Get func(index string, id string, o ...func(*GetRequest)) (*Response, error)

// GetRequest configures the Get API request.
//
type GetRequest struct {
	Index        string
	DocumentType string
	DocumentID   string

	Parent         string
	Preference     string
	Realtime       *bool
	Refresh        *bool
	Routing        string
	Source         []string
	SourceExclude  []string
	SourceExcludes []string
	SourceInclude  []string
	SourceIncludes []string
	StoredFields   []string
	Version        *int
	VersionType    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r GetRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	if r.DocumentType == "" {
		r.DocumentType = "_doc"
	}

	path.Grow(1 + len(r.Index) + 1 + len(r.DocumentType) + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString(r.DocumentType)
	path.WriteString("/")
	path.WriteString(r.DocumentID)

	params = make(map[string]string)

	if r.Parent != "" {
		params["parent"] = r.Parent
	}

	if r.Preference != "" {
		params["preference"] = r.Preference
	}

	if r.Realtime != nil {
		params["realtime"] = strconv.FormatBool(*r.Realtime)
	}

	if r.Refresh != nil {
		params["refresh"] = strconv.FormatBool(*r.Refresh)
	}

	if r.Routing != "" {
		params["routing"] = r.Routing
	}

	if len(r.Source) > 0 {
		params["_source"] = strings.Join(r.Source, ",")
	}

	if len(r.SourceExclude) > 0 {
		params["_source_exclude"] = strings.Join(r.SourceExclude, ",")
	}

	if len(r.SourceExcludes) > 0 {
		params["_source_excludes"] = strings.Join(r.SourceExcludes, ",")
	}

	if len(r.SourceInclude) > 0 {
		params["_source_include"] = strings.Join(r.SourceInclude, ",")
	}

	if len(r.SourceIncludes) > 0 {
		params["_source_includes"] = strings.Join(r.SourceIncludes, ",")
	}

	if len(r.StoredFields) > 0 {
		params["stored_fields"] = strings.Join(r.StoredFields, ",")
	}

	if r.Version != nil {
		params["version"] = strconv.FormatInt(int64(*r.Version), 10)
	}

	if r.VersionType != "" {
		params["version_type"] = r.VersionType
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f Get) WithContext(v context.Context) func(*GetRequest) {
	return func(r *GetRequest) {
		r.ctx = v
	}
}

// WithDocumentType - the type of the document (use `_all` to fetch the first document matching the ID across all types).
//
func (f Get) WithDocumentType(v string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.DocumentType = v
	}
}

// WithParent - the ID of the parent document.
//
func (f Get) WithParent(v string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Parent = v
	}
}

// WithPreference - specify the node or shard the operation should be performed on (default: random).
//
func (f Get) WithPreference(v string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Preference = v
	}
}

// WithRealtime - specify whether to perform the operation in realtime or search mode.
//
func (f Get) WithRealtime(v bool) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Realtime = &v
	}
}

// WithRefresh - refresh the shard containing the document before performing the operation.
//
func (f Get) WithRefresh(v bool) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Refresh = &v
	}
}

// WithRouting - specific routing value.
//
func (f Get) WithRouting(v string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Routing = v
	}
}

// WithSource - true or false to return the _source field or not, or a list of fields to return.
//
func (f Get) WithSource(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Source = v
	}
}

// WithSourceExclude - a list of fields to exclude from the returned _source field.
//
func (f Get) WithSourceExclude(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.SourceExclude = v
	}
}

// WithSourceExcludes - a list of fields to exclude from the returned _source field.
//
func (f Get) WithSourceExcludes(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.SourceExcludes = v
	}
}

// WithSourceInclude - a list of fields to extract and return from the _source field.
//
func (f Get) WithSourceInclude(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.SourceInclude = v
	}
}

// WithSourceIncludes - a list of fields to extract and return from the _source field.
//
func (f Get) WithSourceIncludes(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.SourceIncludes = v
	}
}

// WithStoredFields - a list of stored fields to return in the response.
//
func (f Get) WithStoredFields(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.StoredFields = v
	}
}

// WithVersion - explicit version number for concurrency control.
//
func (f Get) WithVersion(v int) func(*GetRequest) {
	return func(r *GetRequest) {
		r.Version = &v
	}
}

// WithVersionType - specific version type.
//
func (f Get) WithVersionType(v string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.VersionType = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f Get) WithPretty() func(*GetRequest) {
	return func(r *GetRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f Get) WithHuman() func(*GetRequest) {
	return func(r *GetRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f Get) WithErrorTrace() func(*GetRequest) {
	return func(r *GetRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f Get) WithFilterPath(v ...string) func(*GetRequest) {
	return func(r *GetRequest) {
		r.FilterPath = v
	}
}
