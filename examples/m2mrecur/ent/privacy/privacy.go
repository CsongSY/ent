// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent"
)

var (
	// Allow may be returned by read/write rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by read/write rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by read/write rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

type (
	// ReadPolicy combines multiple read rules into a single policy.
	ReadPolicy []ReadRule

	// ReadRule defines the interface deciding whether a read is allowed.
	ReadRule interface {
		EvalRead(context.Context, ent.Value) error
	}
)

// EvalRead evaluates a load against a read policy.
func (rules ReadPolicy) EvalRead(ctx context.Context, v ent.Value) error {
	for _, rule := range rules {
		if err := rule.EvalRead(ctx, v); !errors.Is(err, ErrSkip) {
			return err
		}
	}
	return nil
}

// ReadRuleFunc type is an adapter to allow the use of
// ordinary functions as read rules.
type ReadRuleFunc func(context.Context, ent.Value) error

// Eval calls f(ctx, v).
func (f ReadRuleFunc) EvalRead(ctx context.Context, v ent.Value) error {
	return f(ctx, v)
}

type (
	// WritePolicy combines multiple write rules into a single policy.
	WritePolicy []WriteRule

	// WriteRule defines the interface deciding whether a write is allowed.
	WriteRule interface {
		EvalWrite(context.Context, ent.Mutation) error
	}
)

// EvalWrite evaluates a mutation against a write policy.
func (rules WritePolicy) EvalWrite(ctx context.Context, m ent.Mutation) error {
	for _, rule := range rules {
		switch err := rule.EvalWrite(ctx, m); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// WriteRuleFunc type is an adapter to allow the use of
// ordinary functions as write rules.
type WriteRuleFunc func(context.Context, ent.Mutation) error

// Eval calls f(ctx, m).
func (f WriteRuleFunc) EvalWrite(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups ReadPolicy and WritePolicy.
type Policy struct {
	ReadPolicy
	WritePolicy
}

// ReadRules appends a set read of rules to policy.
func (p *Policy) ReadRules(rules ...ReadRule) *Policy {
	p.ReadPolicy = append(p.ReadPolicy, rules...)
	return p
}

// WriteRules appends a set read of rules to policy.
func (p *Policy) WriteRules(rules ...WriteRule) *Policy {
	p.WritePolicy = append(p.WritePolicy, rules...)
	return p
}

// AlwaysAllowReadWrite returns a privacy policy allowing both reads and writes.
func AlwaysAllowReadWrite() ent.Policy {
	return alwaysAllowReadWrite{}
}

type alwaysAllowReadWrite struct{}

func (alwaysAllowReadWrite) EvalRead(context.Context, ent.Value) error     { return nil }
func (alwaysAllowReadWrite) EvalWrite(context.Context, ent.Mutation) error { return nil }
