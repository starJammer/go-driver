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

import (
	"context"
	"io"
)

// Cursor is returned from a query, used to iterate over a list of documents.
// Note that a Cursor must always be closed to avoid holding on to resources in the server while they are no longer needed.
type Cursor interface {
	io.Closer

	// HasMore returns true if the next call to ReadDocument does not return a NoMoreDocuments error.
	HasMore() bool

	// ReadDocument reads the next document from the cursor.
	// The document data is stored into result, the document meta data is returned.
	// If the cursor has no more documents, a NoMoreDocuments error is returned.
	ReadDocument(ctx context.Context, result interface{}) (DocumentMeta, error)

	// Count returns the total number of result documents available.
	// A valid return value is only available when the cursor has been created with a context that was
	// prepare with `WithQueryCount`.
	Count() int64
}
