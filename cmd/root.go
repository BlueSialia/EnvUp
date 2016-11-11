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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
	"bytes"
)

var defaultConfigPath string
var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "envup [command]",
	Short: "A simple CLI tool for managing your dotfiles repository",
	Long: `EnvUp is a simple command line application that tries to go a little bit further than your average dotfile
	manager. On top of managing your dotfiles, EnvUp can install packages from APT and run scripts. The goal is to get
	your customized environment up and running with a single command.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME == "" {
		XDG_CONFIG_HOME = os.Getenv("HOME") + "/.config"
	}
	defaultConfigPath = XDG_CONFIG_HOME + "/envup"
	defaultConfigFile := defaultConfigPath + "/envup.yaml"

	// Global Flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultConfigFile, "config file for EnvUp")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") //TODO: delete
}

func initConfig() {
	// Default Values
	viper.SetDefault("EnvUpRepo", defaultConfigPath + "/repository")


	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("envup") // name of config file (without extension)
	viper.AddConfigPath(defaultConfigPath)  // adding default directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func executeCmd(name string, args ...string) {
	command := exec.Command(name, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		panic(err)
	}

	fmt.Print(out.String())
}
