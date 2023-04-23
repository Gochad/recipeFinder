package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ingredients string
var numberOfRecipes int

var rootCmd = &cobra.Command{
	Use:   "recipeFinder",
	Short: "CLI app for searching recipes",
	Long: `CLI app for listing meals that can be prepered with minimal number of missing ingredients from given ingredients
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
	return ingredients, numberOfRecipes
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&ingredients, "ingredients", "", "ingredients string")
	rootCmd.Flags().IntVar(&numberOfRecipes, "numberOfRecipes", 0, "number of recipes int")
}

func initConfig() {
	ingredients, _ = rootCmd.Flags().GetString("ingredients")
	numberOfRecipes, _ = rootCmd.Flags().GetInt("numberOfRecipes")
}
