/*
Copyright 2020 The Kubernetes Authors.

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

package cmd

import (
	"github.com/spf13/cobra"

	"k8s.io/enhancements/pkg/kepctl"
	"k8s.io/enhancements/pkg/kepval"
)

// TODO: Struct literal instead?
//var validateOpt

var validateCmd = &cobra.Command{
	Use:           "validate [KEP]",
	Short:         "Validate a new KEP",
	Long:          "Validate a new KEP using the current KEP template for the given type",
	Example:       `  kepctl validate sig-architecture/000-mykep`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(*cobra.Command, []string) error {
		return runValidate(rootOpts)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func runValidate(opts *kepctl.CommonArgs) error {
	_, _, err := kepval.ValidateRepository(opts.RepoPath)
	if err != nil {
		return err
	}

	/*
		if len(warnings) > 0 {
			logrus.Infof("warnings: %v", warnings)
		}

		if len(valErrs) > 0 {
			logrus.Infof("validation errors: %v", valErrs)
		}
	*/

	return nil
}
