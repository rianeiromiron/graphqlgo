package graphqlserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	appdb "graphqlexample/internal/db"

	"github.com/graph-gophers/graphql-go"
)

const schemaText = `
	type Task {
		id: ID!
		title: String!
		description: String
		done: Boolean!
	}

	type Query {
		tasks: [Task!]!
		task(id: ID!): Task
	}

	type Mutation {
		createTask(title: String!, description: String, done: Boolean): Task!
		updateTask(id: ID!, title: String, description: String, done: Boolean): Task!
		deleteTask(id: ID!): Boolean!
	}
`

type Task struct {
	id          int64
	title       string
	description string
	done        bool
}

func (t *Task) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(t.id, 10))
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() *string {
	if t.description == "" {
		return nil
	}
	value := t.description
	return &value
}

func (t *Task) Done() bool {
	return t.done
}

type Resolver struct {
	DB *sql.DB
}

func NewHTTPHandler(db *sql.DB) http.Handler {
	schema := graphql.MustParseSchema(schemaText, &Resolver{DB: db})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(request.Query) == "" {
			http.Error(w, "Query is required", http.StatusBadRequest)
			return
		}

		response := schema.Exec(r.Context(), request.Query, "", request.Variables)
		if response == nil {
			http.Error(w, "GraphQL execution failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
}

func (r *Resolver) Tasks() ([]*Task, error) {
	items, err := appdb.ListTasks(r.DB)
	if err != nil {
		return nil, err
	}

	result := make([]*Task, 0, len(items))
	for _, item := range items {
		result = append(result, toGraphQLTask(item))
	}
	return result, nil
}

func (r *Resolver) Task(args struct{ ID string }) (*Task, error) {
	id, err := parseID(args.ID)
	if err != nil {
		return nil, err
	}

	item, err := appdb.GetTaskByID(r.DB, id)
	if err != nil {
		return nil, err
	}

	return toGraphQLTask(*item), nil
}

func (r *Resolver) CreateTask(args struct {
	Title       string
	Description *string
	Done        *bool
}) (*Task, error) {
	description := ""
	if args.Description != nil {
		description = *args.Description
	}

	done := false
	if args.Done != nil {
		done = *args.Done
	}

	item, err := appdb.CreateTask(r.DB, args.Title, description, done)
	if err != nil {
		return nil, err
	}

	return toGraphQLTask(*item), nil
}

func (r *Resolver) UpdateTask(args struct {
	ID          string
	Title       *string
	Description *string
	Done        *bool
}) (*Task, error) {
	id, err := parseID(args.ID)
	if err != nil {
		return nil, err
	}

	current, err := appdb.GetTaskByID(r.DB, id)
	if err != nil {
		return nil, err
	}

	title := current.Title
	if args.Title != nil {
		title = *args.Title
	}

	description := current.Description
	if args.Description != nil {
		description = *args.Description
	}

	done := current.Done
	if args.Done != nil {
		done = *args.Done
	}

	item, err := appdb.UpdateTask(r.DB, id, title, description, done)
	if err != nil {
		return nil, err
	}

	return toGraphQLTask(*item), nil
}

func (r *Resolver) DeleteTask(args struct{ ID string }) (bool, error) {
	id, err := parseID(args.ID)
	if err != nil {
		return false, err
	}

	deleted, err := appdb.DeleteTask(r.DB, id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func toGraphQLTask(item appdb.Task) *Task {
	return &Task{
		id:          item.ID,
		title:       item.Title,
		description: item.Description,
		done:        item.Done,
	}
}

func parseID(value string) (int64, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return 0, fmt.Errorf("id is required")
	}

	id, err := strconv.ParseInt(trimmed, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id %q: %w", value, err)
	}
	return id, nil
}
