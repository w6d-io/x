/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 07/10/2021
*/

package cmdx

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version returns a *cobra.Command that handles the application version
func Version(tag, commit, buildTime *string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the build version, build commit and build time",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(*tag) == 0 {
				_, _ = fmt.Fprintln(os.Stderr, "Unable to determine version because the build process did not properly configure it.")
			} else {
				fmt.Printf("Version:      %s\n", *tag)
			}

			if len(*commit) == 0 {
				_, _ = fmt.Fprintln(os.Stderr, "Unable to determine commit sha because the build process did not properly configure it.")
			} else {
				fmt.Printf("Build Commit: %s\n", *commit)
			}

			if len(*buildTime) == 0 {
				_, _ = fmt.Fprintln(os.Stderr, "Unable to determine build timestamp because the build process did not properly configure it.")
			} else {
				fmt.Printf("Build Commit: %s\n", *buildTime)
			}
			return nil
		},
	}
}
