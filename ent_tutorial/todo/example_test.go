package todo

import (
	"context"
	"fmt"
	"log"
	"todo/ent"
	enttodo "todo/ent/todo"

	"entgo.io/ent/dialect"
	// "github.com/mattn/go-sqlite3"
)

func Example_todo() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// ...
	task1, err := client.Todo.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Println(task1)
	// ...
	task2, err := client.Todo.Create().SetText("Add Tracking Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Printf("%d: %q\n", task2.ID, task2.Text)

	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("fataled connecting todo2 to its parent: %v", err)
	}

	items, err = client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}

	items, err = client.Todo.Query().Where(enttodo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
    for _, t := range items {
        fmt.Printf("%d: %q\n", t.ID, t.Text)
    }

	items, err = client.Todo.Query().
		Where(
			enttodo.Not(
				enttodo.HasParent(),
			),
			enttodo.HasChildren(),
		).
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
    // 子TODOを通じて親TODOを取得し、
    // クエリが正確に1つのTODOを返すことを期待します。
	parent, err := client.Todo.Query().
		Where(enttodo.HasParent()).
		QueryParent().
		Only(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Printf("%d: %q\n", parent.ID, parent.Text)
}
