package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"recipeFinder/cmd"
)

var apiKey string = "211df952b8374a47b1a20aa19571620b"

func generateMeals() {
	ingrediens, numberOfRecipes := cmd.Execute()
	showRecipes(ingrediens, numberOfRecipes)
}

type Recipes []struct {
	Id                int    `json: "id`
	Title             string `json: "title"`
	MissedIngredients []struct {
		Id     int     `json: "id"`
		Amount float32 `json: "amount"`
		Name   string  `json: "name"`
	} `json: "missedIngredients"`
	UsedIngredients []struct {
		Id     int     `json: "id"`
		Amount float32 `json: "amount"`
		Name   string  `json: "name"`
	} `json: "usedIngredients"`
}

type Nutrition struct {
	Calories string `json: "calories"`
	Carbs    string `json: "carbs"`
	Protein  string `json: "protein"`
}

func showRecipes(ingrediens string, number int) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/findByIngredients?apiKey=%s&ingredients=%s&number=%d", apiKey, ingrediens, number)
	responseByte := getData(url)
	data := Recipes{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		panic(err)
	}
	for _, meal := range data {
		fmt.Printf("TITLE: %s\n", meal.Title)
		nutrition := getRecipeNutrition(meal.Id)
		fmt.Printf("NUTRITION:\nCalories: %s\t Carbs: %s\t Protein: %s\n", nutrition.Calories, nutrition.Carbs, nutrition.Protein)
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
func getRecipeNutrition(recipeId int) Nutrition {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/nutritionWidget.json?apiKey=%s", recipeId, apiKey)
	responseByte := getData(url)
	data := Nutrition{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("unable to get data")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("unable to read the response")
	}
	return responseByte
}
