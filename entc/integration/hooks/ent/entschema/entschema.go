// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"time"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent/schema"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent/card"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.

func init() {
	cardHooks := schema.Card{}.Hooks()
	for i, h := range cardHooks {
		card.Hooks[i] = h
	}
	cardFields := schema.Card{}.Fields()
	// cardDescBoring is the schema descriptor for boring field.
	cardDescBoring := cardFields[0].Descriptor()
	// card.DefaultBoring holds the default value on creation for the boring field.
	card.DefaultBoring = cardDescBoring.Default.(func() time.Time)
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[1].Descriptor()
	// card.DefaultNumber holds the default value on creation for the number field.
	card.DefaultNumber = cardDescNumber.Default.(string)
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
}

// entc build info
const (
	Version = "(devel)"
)