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

package minimatch

import (
	"github.com/sirupsen/logrus"
	"open-match.dev/open-match/internal/config"
	"open-match.dev/open-match/internal/future/app/backend"
	"open-match.dev/open-match/internal/future/app/frontend"
	"open-match.dev/open-match/internal/future/app/mmlogic"
	"open-match.dev/open-match/internal/future/serving"
)

var (
	minimatchLogger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "minimatch",
	})
)

// RunApplication creates a server.
func RunApplication() {
	cfg, err := config.Read()
	if err != nil {
		minimatchLogger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalf("cannot read configuration.")
	}
	p, err := serving.NewParamsFromConfig(cfg, "api.frontend")
	if err != nil {
		minimatchLogger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalf("cannot construct server.")
	}

	BindService(p)
	serving.MustServeForever(p)
}

// BindService creates the minimatch service to the server Params.
func BindService(p *serving.Params) {
	backend.BindService(p)
	frontend.BindService(p)
	mmlogic.BindService(p)
}
