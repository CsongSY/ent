// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent/migrate"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Planet is the client for interacting with the Planet builders.
	Planet *PlanetClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Planet: NewPlanetClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config: cfg,
		Planet: NewPlanetClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Planet.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config: cfg,
		Schema: migrate.NewSchema(cfg.driver),
		Planet: NewPlanetClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Planet.Use(hooks...)
}

// PlanetClient is a client for the Planet schema.
type PlanetClient struct {
	config
}

// NewPlanetClient returns a client for the Planet from the given config.
func NewPlanetClient(c config) *PlanetClient {
	return &PlanetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `planet.Hooks(f(g(h())))`.
func (c *PlanetClient) Use(hooks ...Hook) {
	c.hooks = append(c.hooks[:len(c.hooks):len(c.hooks)], hooks...)
}

// Create returns a create builder for Planet.
func (c *PlanetClient) Create() *PlanetCreate {
	hooks := c.hooks
	mutation := newPlanetMutation(c.config, OpCreate)
	return &PlanetCreate{config: c.config, hooks: hooks, mutation: mutation}
}

// Update returns an update builder for Planet.
func (c *PlanetClient) Update() *PlanetUpdate {
	hooks := c.hooks
	mutation := newPlanetMutation(c.config, OpUpdate)
	return &PlanetUpdate{config: c.config, hooks: hooks, mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanetClient) UpdateOne(pl *Planet) *PlanetUpdateOne {
	return c.UpdateOneID(pl.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanetClient) UpdateOneID(id uint64) *PlanetUpdateOne {
	hooks := c.hooks
	mutation := newPlanetMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &PlanetUpdateOne{config: c.config, hooks: hooks, mutation: mutation}
}

// Delete returns a delete builder for Planet.
func (c *PlanetClient) Delete() *PlanetDelete {
	hooks := c.hooks
	mutation := newPlanetMutation(c.config, OpDelete)
	return &PlanetDelete{config: c.config, hooks: hooks, mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlanetClient) DeleteOne(pl *Planet) *PlanetDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlanetClient) DeleteOneID(id uint64) *PlanetDeleteOne {
	builder := c.Delete().Where(planet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlanetDeleteOne{builder}
}

// Create returns a query builder for Planet.
func (c *PlanetClient) Query() *PlanetQuery {
	return &PlanetQuery{config: c.config}
}

// Get returns a Planet entity by its id.
func (c *PlanetClient) Get(ctx context.Context, id uint64) (*Planet, error) {
	return c.Query().Where(planet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanetClient) GetX(ctx context.Context, id uint64) *Planet {
	pl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pl
}

// QueryNeighbors queries the neighbors edge of a Planet.
func (c *PlanetClient) QueryNeighbors(pl *Planet) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	id := pl.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(planet.Table, planet.FieldID, id),
		sqlgraph.To(planet.Table, planet.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, planet.NeighborsTable, planet.NeighborsPrimaryKey...),
	)
	query.sql = sqlgraph.Neighbors(pl.driver.Dialect(), step)

	return query
}
