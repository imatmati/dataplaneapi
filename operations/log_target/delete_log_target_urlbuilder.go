// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteLogTargetURL generates an URL for the delete log target operation
type DeleteLogTargetURL struct {
	ID int64

	ForceReload   *bool
	ParentName    string
	ParentType    string
	TransactionID *string
	Version       *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteLogTargetURL) WithBasePath(bp string) *DeleteLogTargetURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteLogTargetURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteLogTargetURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/services/haproxy/configuration/log_targets/{id}"

	id := swag.FormatInt64(o.ID)
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on DeleteLogTargetURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var forceReload string
	if o.ForceReload != nil {
		forceReload = swag.FormatBool(*o.ForceReload)
	}
	if forceReload != "" {
		qs.Set("force_reload", forceReload)
	}

	parentName := o.ParentName
	if parentName != "" {
		qs.Set("parent_name", parentName)
	}

	parentType := o.ParentType
	if parentType != "" {
		qs.Set("parent_type", parentType)
	}

	var transactionID string
	if o.TransactionID != nil {
		transactionID = *o.TransactionID
	}
	if transactionID != "" {
		qs.Set("transaction_id", transactionID)
	}

	var version string
	if o.Version != nil {
		version = swag.FormatInt64(*o.Version)
	}
	if version != "" {
		qs.Set("version", version)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteLogTargetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteLogTargetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteLogTargetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteLogTargetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteLogTargetURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteLogTargetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
