package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shurcooL/graphql"
)

type Job struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
}

type Person struct {
	ID           int          `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Address      string       `json:"address,omitempty"`
	Job          Job          `json:"job,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type Organization struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Employees []Person `json:"employees,omitempty"`
}

func main() {
	client := graphql.NewClient("http://localhost:8080/graphql", nil)

	var m struct {
		AddOrganization struct {
			Input Organization
		} `graphql:"addOrganization(input: $input)"`
	}

	input := Organization{
		Name: "Test ORG",
	}

	variables := map[string]interface{}{
		"input": input,
	}

	err := client.Mutate(context.Background(), &m, variables)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("m: %#v\n", m)
}
