// Code generated by github.com/jewzaam/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	pkg "github.com/Azure/ARO-RP/pkg/api"
)

type maintenanceManifestDocumentClient struct {
	*databaseClient
	path string
}

// MaintenanceManifestDocumentClient is a maintenanceManifestDocument client
type MaintenanceManifestDocumentClient interface {
	Create(context.Context, string, *pkg.MaintenanceManifestDocument, *Options) (*pkg.MaintenanceManifestDocument, error)
	List(*Options) MaintenanceManifestDocumentIterator
	ListAll(context.Context, *Options) (*pkg.MaintenanceManifestDocuments, error)
	Get(context.Context, string, string, *Options) (*pkg.MaintenanceManifestDocument, error)
	Replace(context.Context, string, *pkg.MaintenanceManifestDocument, *Options) (*pkg.MaintenanceManifestDocument, error)
	Delete(context.Context, string, *pkg.MaintenanceManifestDocument, *Options) error
	Query(string, *Query, *Options) MaintenanceManifestDocumentRawIterator
	QueryAll(context.Context, string, *Query, *Options) (*pkg.MaintenanceManifestDocuments, error)
	ChangeFeed(*Options) MaintenanceManifestDocumentIterator
}

type maintenanceManifestDocumentChangeFeedIterator struct {
	*maintenanceManifestDocumentClient
	continuation string
	options      *Options
}

type maintenanceManifestDocumentListIterator struct {
	*maintenanceManifestDocumentClient
	continuation string
	done         bool
	options      *Options
}

type maintenanceManifestDocumentQueryIterator struct {
	*maintenanceManifestDocumentClient
	partitionkey string
	query        *Query
	continuation string
	done         bool
	options      *Options
}

// MaintenanceManifestDocumentIterator is a maintenanceManifestDocument iterator
type MaintenanceManifestDocumentIterator interface {
	Next(context.Context, int) (*pkg.MaintenanceManifestDocuments, error)
	Continuation() string
}

// MaintenanceManifestDocumentRawIterator is a maintenanceManifestDocument raw iterator
type MaintenanceManifestDocumentRawIterator interface {
	MaintenanceManifestDocumentIterator
	NextRaw(context.Context, int, interface{}) error
}

// NewMaintenanceManifestDocumentClient returns a new maintenanceManifestDocument client
func NewMaintenanceManifestDocumentClient(collc CollectionClient, collid string) MaintenanceManifestDocumentClient {
	return &maintenanceManifestDocumentClient{
		databaseClient: collc.(*collectionClient).databaseClient,
		path:           collc.(*collectionClient).path + "/colls/" + collid,
	}
}

func (c *maintenanceManifestDocumentClient) all(ctx context.Context, i MaintenanceManifestDocumentIterator) (*pkg.MaintenanceManifestDocuments, error) {
	allmaintenanceManifestDocuments := &pkg.MaintenanceManifestDocuments{}

	for {
		maintenanceManifestDocuments, err := i.Next(ctx, -1)
		if err != nil {
			return nil, err
		}
		if maintenanceManifestDocuments == nil {
			break
		}

		allmaintenanceManifestDocuments.Count += maintenanceManifestDocuments.Count
		allmaintenanceManifestDocuments.ResourceID = maintenanceManifestDocuments.ResourceID
		allmaintenanceManifestDocuments.MaintenanceManifestDocuments = append(allmaintenanceManifestDocuments.MaintenanceManifestDocuments, maintenanceManifestDocuments.MaintenanceManifestDocuments...)
	}

	return allmaintenanceManifestDocuments, nil
}

func (c *maintenanceManifestDocumentClient) Create(ctx context.Context, partitionkey string, newmaintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) (maintenanceManifestDocument *pkg.MaintenanceManifestDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	if options == nil {
		options = &Options{}
	}
	options.NoETag = true

	err = c.setOptions(options, newmaintenanceManifestDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodPost, c.path+"/docs", "docs", c.path, http.StatusCreated, &newmaintenanceManifestDocument, &maintenanceManifestDocument, headers)
	return
}

func (c *maintenanceManifestDocumentClient) List(options *Options) MaintenanceManifestDocumentIterator {
	continuation := ""
	if options != nil {
		continuation = options.Continuation
	}

	return &maintenanceManifestDocumentListIterator{maintenanceManifestDocumentClient: c, options: options, continuation: continuation}
}

func (c *maintenanceManifestDocumentClient) ListAll(ctx context.Context, options *Options) (*pkg.MaintenanceManifestDocuments, error) {
	return c.all(ctx, c.List(options))
}

func (c *maintenanceManifestDocumentClient) Get(ctx context.Context, partitionkey, maintenanceManifestDocumentid string, options *Options) (maintenanceManifestDocument *pkg.MaintenanceManifestDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, nil, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodGet, c.path+"/docs/"+maintenanceManifestDocumentid, "docs", c.path+"/docs/"+maintenanceManifestDocumentid, http.StatusOK, nil, &maintenanceManifestDocument, headers)
	return
}

func (c *maintenanceManifestDocumentClient) Replace(ctx context.Context, partitionkey string, newmaintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) (maintenanceManifestDocument *pkg.MaintenanceManifestDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, newmaintenanceManifestDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodPut, c.path+"/docs/"+newmaintenanceManifestDocument.ID, "docs", c.path+"/docs/"+newmaintenanceManifestDocument.ID, http.StatusOK, &newmaintenanceManifestDocument, &maintenanceManifestDocument, headers)
	return
}

func (c *maintenanceManifestDocumentClient) Delete(ctx context.Context, partitionkey string, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) (err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, maintenanceManifestDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodDelete, c.path+"/docs/"+maintenanceManifestDocument.ID, "docs", c.path+"/docs/"+maintenanceManifestDocument.ID, http.StatusNoContent, nil, nil, headers)
	return
}

func (c *maintenanceManifestDocumentClient) Query(partitionkey string, query *Query, options *Options) MaintenanceManifestDocumentRawIterator {
	continuation := ""
	if options != nil {
		continuation = options.Continuation
	}

	return &maintenanceManifestDocumentQueryIterator{maintenanceManifestDocumentClient: c, partitionkey: partitionkey, query: query, options: options, continuation: continuation}
}

func (c *maintenanceManifestDocumentClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.MaintenanceManifestDocuments, error) {
	return c.all(ctx, c.Query(partitionkey, query, options))
}

func (c *maintenanceManifestDocumentClient) ChangeFeed(options *Options) MaintenanceManifestDocumentIterator {
	continuation := ""
	if options != nil {
		continuation = options.Continuation
	}

	return &maintenanceManifestDocumentChangeFeedIterator{maintenanceManifestDocumentClient: c, options: options, continuation: continuation}
}

func (c *maintenanceManifestDocumentClient) setOptions(options *Options, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, headers http.Header) error {
	if options == nil {
		return nil
	}

	if maintenanceManifestDocument != nil && !options.NoETag {
		if maintenanceManifestDocument.ETag == "" {
			return ErrETagRequired
		}
		headers.Set("If-Match", maintenanceManifestDocument.ETag)
	}
	if len(options.PreTriggers) > 0 {
		headers.Set("X-Ms-Documentdb-Pre-Trigger-Include", strings.Join(options.PreTriggers, ","))
	}
	if len(options.PostTriggers) > 0 {
		headers.Set("X-Ms-Documentdb-Post-Trigger-Include", strings.Join(options.PostTriggers, ","))
	}
	if len(options.PartitionKeyRangeID) > 0 {
		headers.Set("X-Ms-Documentdb-PartitionKeyRangeID", options.PartitionKeyRangeID)
	}

	return nil
}

func (i *maintenanceManifestDocumentChangeFeedIterator) Next(ctx context.Context, maxItemCount int) (maintenanceManifestDocuments *pkg.MaintenanceManifestDocuments, err error) {
	headers := http.Header{}
	headers.Set("A-IM", "Incremental feed")

	headers.Set("X-Ms-Max-Item-Count", strconv.Itoa(maxItemCount))
	if i.continuation != "" {
		headers.Set("If-None-Match", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &maintenanceManifestDocuments, headers)
	if IsErrorStatusCode(err, http.StatusNotModified) {
		err = nil
	}
	if err != nil {
		return
	}

	i.continuation = headers.Get("Etag")

	return
}

func (i *maintenanceManifestDocumentChangeFeedIterator) Continuation() string {
	return i.continuation
}

func (i *maintenanceManifestDocumentListIterator) Next(ctx context.Context, maxItemCount int) (maintenanceManifestDocuments *pkg.MaintenanceManifestDocuments, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Max-Item-Count", strconv.Itoa(maxItemCount))
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &maintenanceManifestDocuments, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

func (i *maintenanceManifestDocumentListIterator) Continuation() string {
	return i.continuation
}

func (i *maintenanceManifestDocumentQueryIterator) Next(ctx context.Context, maxItemCount int) (maintenanceManifestDocuments *pkg.MaintenanceManifestDocuments, err error) {
	err = i.NextRaw(ctx, maxItemCount, &maintenanceManifestDocuments)
	return
}

func (i *maintenanceManifestDocumentQueryIterator) NextRaw(ctx context.Context, maxItemCount int, raw interface{}) (err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Max-Item-Count", strconv.Itoa(maxItemCount))
	headers.Set("X-Ms-Documentdb-Isquery", "True")
	headers.Set("Content-Type", "application/query+json")
	if i.partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+i.partitionkey+`"]`)
	} else {
		headers.Set("X-Ms-Documentdb-Query-Enablecrosspartition", "True")
	}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodPost, i.path+"/docs", "docs", i.path, http.StatusOK, &i.query, &raw, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

func (i *maintenanceManifestDocumentQueryIterator) Continuation() string {
	return i.continuation
}