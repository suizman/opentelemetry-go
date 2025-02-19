// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncfloat64 // import "go.opentelemetry.io/otel/metric/instrument/syncfloat64"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument"
)

// Counter is an instrument that records increasing values.
//
// Warning: methods may be added to this interface in minor releases.
type Counter interface {
	// Add records a change to the counter.
	Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	instrument.Synchronous
}

// UpDownCounter is an instrument that records increasing or decreasing values.
//
// Warning: methods may be added to this interface in minor releases.
type UpDownCounter interface {
	// Add records a change to the counter.
	Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	instrument.Synchronous
}

// Histogram is an instrument that records a distribution of values.
//
// Warning: methods may be added to this interface in minor releases.
type Histogram interface {
	// Record adds an additional value to the distribution.
	Record(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	instrument.Synchronous
}
