/*
Copyright © 2023 Immanuel Klinkenberg <drieschel@yahoo.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drieschel/mailboxsync/internal/sync"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var concurrentSyncs int
var rootCmd = &cobra.Command{
	Use:   "mailbox-sync /file/to/mailboxes.json",
	Short: "Synchronize multiple mailboxes between different servers at the same time",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		if _, err := os.Stat(args[0]); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %q does not exist", args[0])
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}

		defer jsonFile.Close()

		var syncs []sync.Sync
		jsonBytes, _ := io.ReadAll(jsonFile)
		if json.Valid(jsonBytes) == false {
			log.Fatalf("json data in file %s is not valid", args[0])
		}

		err = json.Unmarshal(jsonBytes, &syncs)
		if err != nil {
			log.Fatalf("%v", err)
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(syncs)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); !ok {
				log.Fatal(err)
			}
		}

		sync.NewService(concurrentSyncs).SyncMailboxes(syncs)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	rootCmd.PersistentFlags().IntVar(&concurrentSyncs, "concurrent-syncs", 3, "max amount of concurrent mailbox syncs")
}
