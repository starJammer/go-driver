//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package driver

import "context"

// CollectionDocuments provides access to the documents in a single collection.
type CollectionDocuments interface {
	// ReadDocument reads a single document with given key from the collection.
	// The document data is stored into result, the document meta data is returned.
	// If no document exists with given key, a NotFoundError is returned.
	ReadDocument(ctx context.Context, key string, result interface{}) (DocumentMeta, error)

	// CreateDocument creates a single document in the collection.
	// The document data is loaded from the given document, the document meta data is returned.
	// If the document data already contains a `_key` field, this will be used as key of the new document,
	// otherwise a unique key is created.
	// A ConflictError is returned when a `_key` field contains a duplicate key, other any other field violates an index constraint.
	// To return the NEW document, prepare a context with `WithReturnNew`.
	// To wait until document has been synced to disk, prepare a context with `WithWaitForSync`.
	CreateDocument(ctx context.Context, document interface{}) (DocumentMeta, error)

	// CreateDocuments creates multiple documents in the collection.
	// The document data is loaded from the given documents slice, the documents meta data is returned.
	// If a documents element already contains a `_key` field, this will be used as key of the new document,
	// otherwise a unique key is created.
	// If a documents element contains a `_key` field with a duplicate key, other any other field violates an index constraint,
	// a ConflictError is returned in its inded in the errors slice.
	// To return the NEW documents, prepare a context with `WithReturnNew`. The data argument passed to `WithReturnNew` must be
	// a slice with the same number of entries as the `documents` slice.
	// To wait until document has been synced to disk, prepare a context with `WithWaitForSync`.
	// If the create request itself fails or one of the arguments is invalid, an error is returned.
	CreateDocuments(ctx context.Context, documents interface{}) (DocumentMetaSlice, ErrorSlice, error)

	// UpdateDocument updates a single document with given key in the collection.
	// The document meta data is returned.
	// To return the NEW document, prepare a context with `WithReturnNew`.
	// To return the OLD document, prepare a context with `WithReturnOld`.
	// To wait until document has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with given key, a NotFoundError is returned.
	UpdateDocument(ctx context.Context, key string, update interface{}) (DocumentMeta, error)

	// UpdateDocuments updates multiple document with given keys in the collection.
	// The updates are loaded from the given updates slice, the documents meta data are returned.
	// To return the NEW documents, prepare a context with `WithReturnNew` with a slice of documents.
	// To return the OLD documents, prepare a context with `WithReturnOld` with a slice of documents.
	// To wait until documents has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with a given key, a NotFoundError is returned at its errors index.
	// If keys is nil, each element in the updates slice must contain a `_key` field.
	UpdateDocuments(ctx context.Context, keys []string, updates interface{}) (DocumentMetaSlice, ErrorSlice, error)

	// ReplaceDocument replaces a single document with given key in the collection with the document given in the document argument.
	// The document meta data is returned.
	// To return the NEW document, prepare a context with `WithReturnNew`.
	// To return the OLD document, prepare a context with `WithReturnOld`.
	// To wait until document has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with given key, a NotFoundError is returned.
	ReplaceDocument(ctx context.Context, key string, document interface{}) (DocumentMeta, error)

	// ReplaceDocuments replaces multiple documents with given keys in the collection with the documents given in the documents argument.
	// The replacements are loaded from the given documents slice, the documents meta data are returned.
	// To return the NEW documents, prepare a context with `WithReturnNew` with a slice of documents.
	// To return the OLD documents, prepare a context with `WithReturnOld` with a slice of documents.
	// To wait until documents has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with a given key, a NotFoundError is returned at its errors index.
	// If keys is nil, each element in the documents slice must contain a `_key` field.
	ReplaceDocuments(ctx context.Context, keys []string, documents interface{}) (DocumentMetaSlice, ErrorSlice, error)

	// RemoveDocument removes a single document with given key from the collection.
	// The document meta data is returned.
	// To return the OLD document, prepare a context with `WithReturnOld`.
	// To wait until removal has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with given key, a NotFoundError is returned.
	RemoveDocument(ctx context.Context, key string) (DocumentMeta, error)

	// RemoveDocuments removes multiple documents with given keys from the collection.
	// The document meta data are returned.
	// To return the OLD documents, prepare a context with `WithReturnOld` with a slice of documents.
	// To wait until removal has been synced to disk, prepare a context with `WithWaitForSync`.
	// If no document exists with a given key, a NotFoundError is returned at its errors index.
	RemoveDocuments(ctx context.Context, keys []string) (DocumentMetaSlice, ErrorSlice, error)
}
