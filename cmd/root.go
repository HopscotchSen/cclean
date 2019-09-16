/*
 * Copyright 2019 Gozap, Inc.
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

import (
	"fmt"
	"os"
	"time"

	"github.com/gozap/cclean/cclean"
	"github.com/spf13/cobra"
)

var timeout time.Duration
var exclude string
var port int = 8500

var rootCmd = &cobra.Command{
	Use:   "cclean [CONSUL_ADDRESS] --port [CONSUL_PORT]",
	Short: "A simple service clean tool for Consul, support port.",
	Long: `
A simple service clean tool for Consul, support port`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Consul server port ", port)
		if len(args) == 1 {
			cclean.Clean(args[0], port, exclude, timeout)
		} else {
			cclean.Clean("", port, exclude, timeout)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().IntVar(&port, "port", 8500, "Consul Server port")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 3*time.Second, "http timeout")
	rootCmd.PersistentFlags().StringVar(&exclude, "exclude", "", "exclude consul node ip (eg: 10.20.0.0/16)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
