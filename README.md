# recipeFinder

CLI app in Go for listing meals that can be prepered with minimal number of missing ingredients.
This application is for job task and for starting learn Go.

API for getting meals: [spoonacular food api](https://spoonacular.com/food-api)

packages used in program:

Cobra for making CLI usage: [https://github.com/spf13/cobra](https://github.com/spf13/cobra)\
MongoDB for local database : [mongoDB installation](https://pkg.go.dev/go.mongodb.org/mongo-driver#section-readme)

CLI app accepts two params:\
`--ingredients`, which will be followed by a comma-separated list of ingredients. \
(eg. `-ingredients=tomatoes,eggs,pasta`) \
`--numberOfRecipes`, which will let the user specify the maximum number of recipes they'd like to get \
(eg. `--numberOfRecipes=5`)

## Install
1. clone repo
2. make sure you have installed go
3. in terminal, run `go get go.mongodb.org/mongo-driver/mongo`
4. next import command: `go get -u github.com/spf13/cobra@latest`
5. `go build`
6. Example call after building: `./recipeFinder --ingredients=tomatoes,eggs,pasta --numberOfRecipes=5`

If you want to install cli on your mongodb - you can edit variable `uri` in file `db_conn.db`
