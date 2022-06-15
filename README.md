# 🌱 greenleaf - simple, type safe and easy to use query builder for MongoDB

![build-img](https://github.com/slavabobik/greenleaf/actions/workflows/build.yml/badge.svg)
[![godoc](https://godoc.org/github.com/slavabobik/greenleaf?status.png)](https://godoc.org/github.com/slavabobik/greenleaf)
[![codecov](https://codecov.io/gh/slavabobik/greenleaf/branch/master/graph/badge.svg?token=XQ85I8ANL5)](https://codecov.io/gh/slavabobik/greenleaf)
    

## Installation
To install use:

```bash
 go get github.com/slavabobik/greenleaf
```   


## Quick examples

```golang

package main

import (
	"context"

	"github.com/slavabobik/greenleaf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("testing").Collection("test")
	doc := greenleaf.M{"name": "Jhon", "tags": []string{"fast", "furious"}, "score": 128, "coins": 10000, "active": true}
	collection.InsertOne(ctx, doc)

	filter := greenleaf.Filter(
		Eq("name", "Jhon"),
		In("tags", []string{"fast", "furious"}),
		Gt("score", 100),
		Lte("score", 200),
		Exists("active", true),
	)
}

```
