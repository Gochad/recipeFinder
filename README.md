# recipeFinder

CLI app in Go for listing meals that can be prepered with minimal number of missing ingredients.
This application is for job task and for starting learn Go.

API for getting meals: [spoonacular food api](https://spoonacular.com/food-api)

packages used in program:

Cobra for making CLI usage: [https://github.com/spf13/cobra](https://github.com/spf13/cobra)\
MongoDB for local database : [mongoDB installation](https://pkg.go.dev/go.mongodb.org/mongo-driver#section-readme)

CLI app accepts two params:\
`--ingredients`, which will be followed by a comma-separated list of ingredients. (eg. `-ingredients=tomatoes,eggs,pasta`) \
`--numberOfRecipes`, which will let the user specify the maximum number of recipes they'd like to get (eg. `--numberOfRecipes=5`)

Example call after building: `./recipeFinder --ingredients=tomatoes,eggs,pasta --numberOfRecipes=5`
