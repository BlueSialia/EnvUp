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
	"github.com/spf13/cobra"
	"os"
	"github.com/spf13/viper"
	"path/filepath"
)

// removeCmd represents the remove command
var removeDotfileCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		relPath, err := filepath.Rel(os.Getenv("HOME"), args[0])
		if (err != nil) {
			executeCmd("rm", "-rf", args[0])
			executeCmd("mv", viper.GetString("EnvUpRepo") + "/dotfiles/" + relPath, args[0])
		}
	},
}

func init() {
	removeCmd.AddCommand(removeDotfileCmd)
}
