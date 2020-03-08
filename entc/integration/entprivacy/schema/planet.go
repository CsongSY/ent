package schema

import (
	"context"
	"log"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/entc/integration/entprivacy"
	"github.com/facebookincubator/ent/entc/integration/entprivacy/privacy"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Planet defines the schema of a planet.
type Planet struct {
	ent.Schema
}

func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Immutable().
			NotEmpty().
			Unique(),
		field.Uint("age").
			Optional(),
	}
}

func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("neighbors", Planet.Type),
	}
}

func (Planet) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				mutation := m.(*entprivacy.PlanetMutation)
				log.Println("Planet mutation of type", mutation.Type())
				return next.Mutate(ctx, m)
			})
		},
	}
}

func (Planet) Policy() ent.Policy {
	return privacy.AlwaysAllowReadWrite()
}
