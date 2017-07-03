package main

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func main() {
	// Scrape a recipe from Vegan Rocks

	// Connect to neo4j
	driver := bolt.NewDriver()
	// @TODO: Put this in a config
	conn, err := driver.OpenNeo("bolt://api.omnomhub.dev:7687")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
