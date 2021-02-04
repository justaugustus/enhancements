/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kepval

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/blang/semver/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"k8s.io/enhancements/api"
)

func ValidatePRR(kep *api.Proposal, h *api.PRRHandler, prrDir string) error {
	requiredPRRApproval, _, err := isPRRRequired(kep)
	if err != nil {
		return errors.Wrap(err, "checking if PRR is required")
	}

	if !requiredPRRApproval {
		logrus.Debugf("PRR review is not required for %s", kep.Number)
		return nil
	}

	prrFilename := kep.Number + ".yaml"
	prrFilename = filepath.Join(prrsDir, kep.OwningSIG, prrFilename)

	logrus.Infof("PRR file: %s", prrFilename)

	prrFile, err := os.Open(prrFilename)
	if os.IsNotExist(err) {
		// TODO: Is this actually the error we want to return here?
		return err //needsPRRApproval(stageMilestone, kep.Stage, prrFilename)
	}

	if err != nil {
		return errors.Wrapf(err, "could not open file %s", prrFilename)
	}

	// TODO: Create a context to hold the parsers
	prr, prrParseErr := h.Parse(prrFile)
	if prrParseErr != nil {
		return errors.Wrap(prrParseErr, "parsing PRR approval file")
	}

	// TODO: This shouldn't be required once we push the errors into the
	//       parser struct
	if prr.Error != nil {
		return errors.Wrapf(prr.Error, "%v has an error", prrFilename)
	}

	// TODO: Check for edge cases
	var stageMilestone string
	switch kep.Stage {
	case "alpha":
		stageMilestone = kep.Milestone.Alpha
	case "beta":
		stageMilestone = kep.Milestone.Beta
	case "stable":
		stageMilestone = kep.Milestone.Stable
	}

	stagePRRApprover := prr.ApproverForStage(stageMilestone)
	validApprover := api.IsOneOf(stagePRRApprover, h.PRRApprovers)
	if !validApprover {
		return errors.New(
			fmt.Sprintf(
				"this contributor (%s) is not a PRR approver (%v)",
				stagePRRApprover,
				h.PRRApprovers,
			),
		)
	}

	return nil
}

func isPRRRequired(kep *api.Proposal) (required, missingMilestone bool, err error) {
	logrus.Infof("checking if PRR is required")

	required = true
	missingMilestone = kep.IsMissingMilestone()

	if kep.Status != "implementable" {
		required = false
		return required, missingMilestone, nil
	}

	if missingMilestone {
		required = false
		logrus.Warnf("KEP %s is missing the latest milestone field. This will become a validation error in future releases.", kep.Number)

		return required, missingMilestone, nil
	}

	// TODO: Consider making this a function
	prrRequiredAtSemVer, err := semver.ParseTolerant("v1.21")
	if err != nil {
		return required, missingMilestone, errors.Wrap(err, "creating a SemVer object for PRRs")
	}

	latestSemVer, err := semver.ParseTolerant(kep.LatestMilestone)
	if err != nil {
		return required, missingMilestone, errors.Wrap(err, "creating a SemVer object for latest milestone")
	}

	if latestSemVer.LT(prrRequiredAtSemVer) {
		required = false
		return required, missingMilestone, nil
	}

	return required, missingMilestone, nil
}
