package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb+srv://gocha:z7xOvo0Q316DJ1bs@cluster0.wqs9js4.mongodb.net/?retryWrites=true&w=majority"
var recipes *mongo.Collection
var ctx = context.TODO()

type Recipe struct {
	ID                int          `json:"_id"`
	Title             string       `json:"title"`
	Calories          string       `json:"calories"`
	Carbs             string       `json:"carbs"`
	Protein           string       `json:"protein"`
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
	// ingredients = client.Database("recipeFinder").Collection("ingredients")
}
func saveRecipe(meals []Recipes, nutrition Nutrition) {
	for _, meal := range meals {
		recipe := &Recipe{
			ID:                meal.Id,
			Title:             meal.Title,
			Calories:          nutrition.Calories,
			Carbs:             nutrition.Carbs,
			Protein:           nutrition.Protein,
			MissedIngredients: meal.MissedIngredients,
			UsedIngredients:   meal.UsedIngredients,
		}
		_, err := recipes.InsertOne(context.TODO(), recipe)
		errMessage(err)
	}

}

func getRecipe(id int) Recipe {
	filter := bson.D{{"_id", id}}
	cursor, err := recipes.Find(context.TODO(), filter)

	errMessage(err)
	var result []*Recipe
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	return *result[0]
}
