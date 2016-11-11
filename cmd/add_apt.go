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
	"io/ioutil"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// add_aptCmd represents the add_apt command
var addAptCmd = &cobra.Command{
	Use:   "apt [pkg name]",
	Short: "Add an APT package to the EnvUp list of packages",
	Long: `Add an APT package to the list of packages that EnvUp should install when running "envup install [repo URL]".`,
	Run: func(cmd *cobra.Command, args []string) {
		aptFile, err := ioutil.ReadFile(viper.GetString("EnvUpRepo") + "/aptPkgs")
		if os.IsNotExist(err) {
			lines := strings.Split(string(aptFile), "\\s")
		}

			fmt.Println("add_apt called")
	},
}

func init() {
	addCmd.AddCommand(addAptCmd)
}
