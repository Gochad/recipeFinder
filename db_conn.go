package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb+srv://gocha:z7xOvo0Q316DJ1bs@cluster0.wqs9js4.mongodb.net/?retryWrites=true&w=majority"
var recipes *mongo.Collection
var ctx = context.TODO()

type Recipe struct {
	ID                int          `json: "_id"`
	Title             string       `json: "title"`
	Calories          string       `json: "calories"`
	Carbs             string       `json: "carbs"`
	Protein           string       `json: "protein"`
	Ingredients       string       `json: "ingredients"`
	NumberOfRecipes   int          `json: "numberOfRecipes"`
	MissedIngredients []Ingredient `json: "missedIngredients"`
	UsedIngredients   []Ingredient `json: "usedIngredients"`
}

func init() {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	recipes = client.Database("recipeFinder").Collection("recipes")
}
func saveRecipe(meal Meal, nutrition Nutrition, ingredients string, numberOfRecipes int) {
	recipe := &Recipe{
		ID:                meal.Id,
		Title:             meal.Title,
		Calories:          nutrition.Calories,
		Carbs:             nutrition.Carbs,
		Protein:           nutrition.Protein,
		Ingredients:       ingredients,
		NumberOfRecipes:   numberOfRecipes,
		MissedIngredients: meal.MissedIngredients,
		UsedIngredients:   meal.UsedIngredients,
	}
	fmt.Println(recipe)
	_, err := recipes.InsertOne(context.TODO(), recipe)
	errMessage(err)
}

func getRecipes(ingredients string, numberOfRecipes int) []*Recipe {
	filter := bson.D{{"ingredients", ingredients}, {"numberofrecipes", numberOfRecipes}}
	cursor, err := recipes.Find(context.TODO(), filter)
	errMessage(err)
	var result []*Recipe
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	return result
}
