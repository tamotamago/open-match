// Copyright 2019 Google LLC
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

package apisrv

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// OpenCensus Measures. These are exported as metrics to your monitoring system
// https://godoc.org/go.opencensus.io/stats
//
// When making opencensus stats, the 'name' param, with forward slashes changed
// to underscores, is appended to the 'namespace' value passed to the
// prometheus exporter to become the Prometheus metric name. You can also look
// into having Prometheus rewrite your metric names on scrape.
//
//  For example:
//   - defining the promethus export namespace "open_match" when instanciating the exporter:
//			pe, err := promethus.NewExporter(promethus.Options{Namespace: "open_match"})
//   - and naming the request counter "backend/requests_total":
//			MGrpcRequests := stats.Int64("backendapi/requests_total", ...
//   - results in the prometheus metric name:
//			open_match_backendapi_requests_total
//   - [note] when using opencensus views to aggregate the metrics into
//     distribution buckets and such, multiple metrics
//     will be generated with appended types ("<metric>_bucket",
//     "<metric>_count", "<metric>_sum", for example)
//
// In addition, OpenCensus stats propogated to Prometheus have the following
// auto-populated labels pulled from kubernetes, which we should avoid to
// prevent overloading and having to use the HonorLabels param in Prometheus.
//
// - Information about the k8s pod being monitored:
//		"pod" (name of the monitored k8s pod)
//		"namespace" (k8s namespace of the monitored pod)
// - Information about how promethus is gathering the metrics:
//		"instance" (IP and port number being scraped by prometheus)
//		"job" (name of the k8s service being scraped by prometheus)
//		"endpoint" (name of the k8s port in the k8s service being scraped by prometheus)
//
var (
	// Logging instrumentation
	// There's no need to record this measurement directly if you use
	// the logrus hook provided in metrics/helper.go after instantiating the
	// logrus instance in your application code.
	// https://godoc.org/github.com/sirupsen/logrus#LevelHooks
	EvaluatorLogLines = stats.Int64("evaluator/logs_total", "Number of evaluator lines logged", "1")

	// Instrumentation for evaluator execution.
	Evaluations        = stats.Int64("evaluator/evaluations_total", "Number of times all existing proposals have been evaluated", "1")
	EvaluationFailures = stats.Int64("evaluator/failures_total", "Number of failures attempting to evaluate all existing proposals", "1")
)

var (
	// KeyEvalTrigger is used to tag which code path caused the evaluator to run.
	KeyEvalTrigger, _ = tag.NewKey("evalTrigger")
	// KeySeverity is used to tag a the severity of a log message.
	KeySeverity, _ = tag.NewKey("severity")
)

// Package metrics provides some convience views.
// You need to register the views for the data to actually be collected.
// Note: The OpenCensus View 'Description' is exported to Prometheus as the HELP string.
// Note: If you get a "Failed to export to Prometheus: inconsistent label
// cardinality" error, chances are you forgot to set the tags specified in the
// view for a given measure when you tried to do a stats.Record()
var (
	EvaluatorCountView = &view.View{
		Name:        "evaluator/evaluations_total",
		Measure:     Evaluations,
		Description: "Number of times all existing proposals have been evaluated",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{KeyEvalTrigger},
	}

	EvaluatorFailuresCountView = &view.View{
		Name:        "evaluator/failures_total",
		Measure:     EvaluationFailures,
		Description: "Number of failures attempting to evaluate all existing proposals",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{KeyEvalTrigger},
	}
)

// DefaultEvaluatorViews are the default matchmaker orchestrator OpenCensus measure views.
var DefaultEvaluatorViews = []*view.View{
	EvaluatorCountView,
	EvaluatorFailuresCountView,
}
