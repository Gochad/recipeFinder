package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ingrediens string
var numberOfRecipes int

var rootCmd = &cobra.Command{
	Use:   "recipeFinder",
	Short: "CLI app for searching recipes",
	Long: `CLI app for listing meals that can be prepered with minimal number of missing ingredients from given ingrediens
	You can use 2 commands:
	'--ingredients', which will be followed by a comma-separated list of ingredients. (eg. '--ingredients=tomatoes,eggs,pasta')
	'--numberOfRecipes', which will let the user specify the maximum number of recipes they'd like to get (eg. '--numberOfRecipes=5')
	An example call of this CLI app:
	'./recipeFinder --ingredients=tomatoes,eggs,pasta --numberOfRecipes=5'
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() (string, int) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ingrediens, numberOfRecipes
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&ingrediens, "ingrediens", "", "ingrediens string")
	rootCmd.Flags().IntVar(&numberOfRecipes, "numberOfRecipes", 0, "number of recipes int")
}

func initConfig() {
	ingrediens, _ = rootCmd.Flags().GetString("ingrediens")
	numberOfRecipes, _ = rootCmd.Flags().GetInt("numberOfRecipes")
}
