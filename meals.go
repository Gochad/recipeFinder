package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"recipeFinder/cmd"
)

const apiKey string = "ba78847855ab4d6896c72ae7d7f3da39"

type Ingredient struct {
	Id     int     `json: "id"`
	Amount float32 `json: "amount"`
	Name   string  `json: "name"`
}
type Meal struct {
	Id                int          `json: "id`
	Title             string       `json: "title"`
	MissedIngredients []Ingredient `json: "missedIngredients"`
	UsedIngredients   []Ingredient `json: "usedIngredients"`
}

type Recipes []struct {
	Id                int          `json: "id`
	Title             string       `json: "title"`
	MissedIngredients []Ingredient `json: "missedIngredients"`
	UsedIngredients   []Ingredient `json: "usedIngredients"`
}

type Nutrition struct {
	Calories string `json: "calories"`
	Carbs    string `json: "carbs"`
	Protein  string `json: "protein"`
}

func generateMeals() {
	ingredients, numberOfRecipes := cmd.Execute()
	showRecipes(ingredients, numberOfRecipes)
}

func showRecipes(ingredients string, numberOfRecipes int) {
	recipes := getRecipes(ingredients, numberOfRecipes)
	if len(recipes) == 0 {
		showFromAPI(ingredients, numberOfRecipes)
	} else {
		printOutRecipes(recipes)
	}
}

func printOutRecipes(recipes []*Recipe) {
	for _, meal := range recipes {
		fmt.Printf("TITLE: %s\n", meal.Title)
		fmt.Printf("NUTRITION: Calories: %s\t Carbs: %s\t Protein: %s\n", meal.Calories, meal.Carbs, meal.Protein)
		fmt.Printf("MISSED INGREDIENTS: \n")
		for _, ingredient := range meal.MissedIngredients {
			fmt.Printf("id: %d  amount: %g  name: %s \n", ingredient.Id, ingredient.Amount, ingredient.Name)
		}
		fmt.Printf("USED INGREDIENTS: \n")
		for _, ingredient := range meal.UsedIngredients {
			fmt.Printf("id: %d  amount: %g  name: %s \n", ingredient.Id, ingredient.Amount, ingredient.Name)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
	}
}

func getApiRecipesLink(ingredients string, numberOfRecipes int) string {
	return fmt.Sprintf("https://api.spoonacular.com/recipes/findByIngredients?apiKey=%s&ingredients=%s&number=%d", apiKey, ingredients, numberOfRecipes)
}

func getApiNutritionLink(recipeId int) string {
	return fmt.Sprintf("https://api.spoonacular.com/recipes/%d/nutritionWidget.json?apiKey=%s", recipeId, apiKey)
}

func showFromAPI(ingredients string, numberOfRecipes int) {
	url := getApiRecipesLink(ingredients, numberOfRecipes)
	responseByte := getData(url)
	data := Recipes{}
	err := json.Unmarshal(responseByte, &data)
	errMessage(err)
	var recipes []*Recipe
	for _, meal := range data {
		nutrition := getRecipeNutrition(meal.Id)
		recipes = append(recipes, saveRecipe(meal, nutrition, ingredients, numberOfRecipes))
	}
	printOutRecipes(recipes)
}
func getRecipeNutrition(recipeId int) Nutrition {
	url := getApiNutritionLink(recipeId)
	responseByte := getData(url)
	data := Nutrition{}
	err := json.Unmarshal(responseByte, &data)
	errMessage(err)
	return data
}

func getData(url string) []byte {
	response, err := http.Get(url)
	errMessage(err)

	responseByte, err := ioutil.ReadAll(response.Body)
	errMessage(err)
	return responseByte
}
func errMessage(err error) {
	if err != nil {
		panic(err)
	}
}
