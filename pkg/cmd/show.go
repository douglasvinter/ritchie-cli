/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "github.com/spf13/cobra"

// NewShowCmd creates new cmd instance of show command
func NewShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:       "show SUB_COMMAND",
		Short:     "Show context and formula-runnner default",
		Long:      "Show current context and formula-runnner default",
		Example:   "rit show context",
		ValidArgs: []string{"context", "formula-runner"},
		Args:      cobra.OnlyValidArgs,
	}
}
