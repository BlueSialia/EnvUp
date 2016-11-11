// Copyright © 2016 Jorge Domínguez Arnáez (BlueSialia) <bluesialia@gmail.com>
//
// This file is part of EnvUp.
//
// EnvUp is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// EnvUp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with EnvUp. If not, see <http://www.gnu.org/licenses/>.
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// remove_scriptCmd represents the remove_script command
var remove_scriptCmd = &cobra.Command{
	Use:   "remove_script",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("remove_script called")
	},
}

func init() {
	removeCmd.AddCommand(remove_scriptCmd)
}
