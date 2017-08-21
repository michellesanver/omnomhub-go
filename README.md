[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

This is a baby gopher project. You are more than welcome to come with feedback in the way the code is written. 

# OmnomHub
The goal of omnomhub is to be a recipe site for people who like to alter recipes, if ever so slightly. Like Github, but for cooking. There are a lot of copies of recipes where people slightly alter them to match their dietary needs or tastebuds, at omnomhub you will always be able to see the original recipe, and if you want - Make your own version of it. Just like in github you can also collaborate on recipes, discuss and suggest changes to make them better.

## Installation
We are using drifter for neo4j. To run it ensure you have initiated your git submodules. Then run `vagrant up`. You can then access the neo4j browser from api.omnomhub.dev:7474. 

To run the go API simply do this in the core of the project: 
`go run main.go`
