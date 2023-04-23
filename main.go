package main

import "fmt"

func main() {

	// saveRecipe(4, "czwarty przepis", "10g", "5g", "3kcal")
	recipe := getRecipe(23224)
	fmt.Println(recipe.Title)
	//generateMeals()
}
