// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent"
	gen "github.com/facebookincubator/ent/entc/integration/hooks/ent"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/hook"

	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return hook.CardFunc(func(ctx context.Context, m *gen.CardMutation) (ent.Value, error) {
				fmt.Printf("Schema Hook\tOp: %s\tType: %s\tConcreteType: %T\n", m.Op(), m.Type(), m)
				return next.Mutate(ctx, m)
			})
		},
	}
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Time("boring").
			Default(time.Now),
		field.String("number").
			Immutable().
			Default("111").
			NotEmpty(),
		field.String("name").
			Optional().
			Comment("Exact name written on card"),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Card.Type),
		edge.To("best_friend", Card.Type).
			Unique(),
	}
}
