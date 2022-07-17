/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thzoid/trivial/ternary"
	"github.com/thzoid/trivial/tryte"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trivial",
	Short: "trivial is a virtualized ternary computer",
	Long: `trivial is a virtualized ternary computer and
CLI and a library for Go that allows the use of ternary
data types and operations.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		p := []ternary.Instruction{
			{
				Op:  ternary.STORE,
				Arg: tryte.UIntToUnb(1),
			},
			{
				Op:  ternary.LOAD,
				Arg: tryte.UIntToUnb(1),
			},
		}

		cpu := ternary.NewCPU(10)
		cpu.Load(p)

		// providedMemory := os.Args[1]
		// providedTryte, _ := strconv.ParseUint(os.Args[2], 10, 64)

		// m := ternary.NewMemory(uint64((len(providedMemory) + tryte.TRYTE_TRIT - 1) / tryte.TRYTE_TRIT))
		// for i := uint64(0); i < m.Size(); i++ {
		// 	if i+1 < m.Size() {
		// 		m.Set(i, tryte.MustRead(providedMemory[i*tryte.TRYTE_TRIT:i*tryte.TRYTE_TRIT+tryte.TRYTE_TRIT]))
		// 	} else {
		// 		m.Set(i, tryte.MustRead(providedMemory[i*tryte.TRYTE_TRIT:]))
		// 	}
		// }

		// fmt.Println(m.Get(providedTryte))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trivial.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
