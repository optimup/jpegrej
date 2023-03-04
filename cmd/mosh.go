/*
Copyright © 2020 Joel Curtis <joel@curti.se>

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
	"fmt"
	"time"

	"github.com/optimup/jpegrej/pkg"

	"github.com/spf13/cobra"
)

// moshCmd represents the mosh command
var moshCmd = &cobra.Command{
	Use:   "mosh",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: moshRun,
}

var (
	seed   int64
	amount int64
)

func init() {
	rootCmd.AddCommand(moshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	moshCmd.Flags().Int64VarP(&seed, "seed", "s", 0, "Seed for random mosh of jpeg. (int)")
	moshCmd.Flags().Int64VarP(&amount, "amount", "a", 10, "Amount of bytes to replace. (int)")
}

func moshRun(cmd *cobra.Command, args []string) error {
	fmt.Println("mosh called")

	filepath := args[0]
	fileout := args[1]

	jpeg, err := pkg.Jpegload(filepath)
	if err != nil {
		return err
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	jpeg.Seed(seed, amount)
	jpeg.Mosh(fileout)

	fmt.Printf("Filename: %s\n\tSeed: %d\n\tSize: %d\n\tStart & End: %d, %d\n", jpeg.Path, seed, len(jpeg.Data), jpeg.Start, jpeg.End)

	return nil
}
