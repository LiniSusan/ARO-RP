// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ListOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ListOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAll set field All to given value
func (o *ListOptions) WithAll(value bool) *ListOptions {
	o.All = &value
	return o
}

// GetAll returns value of field All
func (o *ListOptions) GetAll() bool {
	if o.All == nil {
		var z bool
		return z
	}
	return *o.All
}

// WithExternal set field External to given value
func (o *ListOptions) WithExternal(value bool) *ListOptions {
	o.External = &value
	return o
}

// GetExternal returns value of field External
func (o *ListOptions) GetExternal() bool {
	if o.External == nil {
		var z bool
		return z
	}
	return *o.External
}

// WithFilters set field Filters to given value
func (o *ListOptions) WithFilters(value map[string][]string) *ListOptions {
	o.Filters = value
	return o
}

// GetFilters returns value of field Filters
func (o *ListOptions) GetFilters() map[string][]string {
	if o.Filters == nil {
		var z map[string][]string
		return z
	}
	return o.Filters
}

// WithLast set field Last to given value
func (o *ListOptions) WithLast(value int) *ListOptions {
	o.Last = &value
	return o
}

// GetLast returns value of field Last
func (o *ListOptions) GetLast() int {
	if o.Last == nil {
		var z int
		return z
	}
	return *o.Last
}

// WithNamespace set field Namespace to given value
func (o *ListOptions) WithNamespace(value bool) *ListOptions {
	o.Namespace = &value
	return o
}

// GetNamespace returns value of field Namespace
func (o *ListOptions) GetNamespace() bool {
	if o.Namespace == nil {
		var z bool
		return z
	}
	return *o.Namespace
}

// WithSize set field Size to given value
func (o *ListOptions) WithSize(value bool) *ListOptions {
	o.Size = &value
	return o
}

// GetSize returns value of field Size
func (o *ListOptions) GetSize() bool {
	if o.Size == nil {
		var z bool
		return z
	}
	return *o.Size
}

// WithSync set field Sync to given value
func (o *ListOptions) WithSync(value bool) *ListOptions {
	o.Sync = &value
	return o
}

// GetSync returns value of field Sync
func (o *ListOptions) GetSync() bool {
	if o.Sync == nil {
		var z bool
		return z
	}
	return *o.Sync
}