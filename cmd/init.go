/*
Copyright © 2020 Juan Ezquerro LLanes <arrase@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var useTemplate bool

var initCmd = &cobra.Command{
	Use:   "init [url]",
	Short: "Init a workspace",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 && !useTemplate {
			return errors.New("requires a git repo or set the --template flag to start from scratch")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&useTemplate, "template", "t", false, "Init workspace from default config template")
}
