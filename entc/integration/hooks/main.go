package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent"
	_ "github.com/facebookincubator/ent/entc/integration/hooks/ent/entschema"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/hook"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			fmt.Println("start")
			defer fmt.Println("end")
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			fmt.Printf("Before Hook\tOp: %s\tType: %s\tConcreteType: %T\n", m.Op(), m.Type(), m)
			defer fmt.Println("Done!")
			if ns, ok := m.(interface{ SetName(string) }); ok {
				ns.SetName("hook name")
			}
			return next.Mutate(ctx, m)
		})
	})

	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			fmt.Println("Concrete hook\t", m.Op())
			return next.Mutate(ctx, m)
		})
	})
	client.Use(hook.On(LogWithConfig(os.Stdout), ent.OpUpdate, ent.OpCreate)) // change Op1|Op2|...

	u := client.Card.Create().SetNumber("A").SaveX(ctx)
	u.Update().SetName("Boring2").SaveX(ctx)
	client.Card.Update().SetName("foo").SaveX(ctx)
	client.Card.DeleteOneID(u.ID).ExecX(ctx)
	client.Card.Delete().ExecX(ctx)
}

func LogWithConfig(w io.Writer) ent.Hook {
	if w == nil {
		w = os.Stdout
	}
	return func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			fmt.Fprintln(w, "Logging Hook:\t", m.Op())
			return next.Mutate(ctx, m)
		})
	}
}
