/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/michaelcoll/version-badge/internal/back"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long: `
Starts the server`,
	Run: serve,
}

func serve(_ *cobra.Command, _ []string) {
	printBanner(version, Serve)

	module := back.New()

	module.GetApiController().Serve()
}

func init() {
	serveCmd.Flags().StringP("conf-location", "c", "config.yml", "The path where the configuration will be loaded.")

	_ = viper.BindPFlag("conf-location", serveCmd.Flags().Lookup("conf-location"))

	viper.SetDefault("conf-location", "config.yml")

	rootCmd.AddCommand(serveCmd)
}
