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
	"net/url"
	"github.com/spf13/viper"
	"os"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [repo URL]",
	Short: "Installs your dotfiles from a remote repository",
	Long: `Clone your EnvUp repository, symlink every dotfile in it, install APT packages and install NPM packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := url.ParseRequestURI(args[0])
		if err != nil {
			panic(err)
		}

		clone(u.String())
		symlink()
		installApt()
		installNpm()
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
}

func clone(gitUrl string) {
	executeCmd("git", "clone", gitUrl, viper.GetString("EnvUpRepo"))
}

func symlink() {
	dotfileDir := viper.GetString("EnvUpRepo") + "/dotfiles"
	executeCmd("ln", "-s", dotfileDir + "/*", os.Getenv("HOME"))
}

func installApt() {
	aptFile := viper.GetString("EnvUpRepo") + "/aptPkgs"
	if _, err := os.Stat(aptFile); os.IsNotExist(err) {
		executeCmd("sudo", "apt", "install", "<", aptFile)
	}
}

func installNpm() {
	npmFile := viper.GetString("EnvUpRepo") + "/npmPkgs"
	if _, err := os.Stat(npmFile); os.IsNotExist(err) {
		executeCmd("npm", "install", "<", npmFile)
	}
}
