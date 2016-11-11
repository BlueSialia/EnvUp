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
	"path/filepath"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addDotfileCmd = &cobra.Command{
	Use:   "dotfile [file/dir path]",
	Short: "Add a file or directory under your home to your EnvUp repository",
	Long:  `Add a file or directory under your home to your EnvUp repository.
	This moves the real file into the git repository and creates a symlink in
	the original location`,
	Run: func(cmd *cobra.Command, args []string) {
		relPath, err := filepath.Rel(os.Getenv("HOME"), args[0])
		if (err != nil) {
		executeCmd("mv", args[0], viper.GetString("EnvUpRepo") + "/dotfiles/" + relPath)
		executeCmd("ln", "-s", viper.GetString("EnvUpRepo") + "/dotfiles/" + relPath + "/*", args[0])
		}
	},
}

func init() {
	addCmd.AddCommand(addDotfileCmd)
}
