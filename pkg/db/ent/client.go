// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/thirdauth"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ThirdAuth is the client for interacting with the ThirdAuth builders.
	ThirdAuth *ThirdAuthClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ThirdAuth = NewThirdAuthClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		ThirdAuth: NewThirdAuthClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		ThirdAuth: NewThirdAuthClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ThirdAuth.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ThirdAuth.Use(hooks...)
}

// ThirdAuthClient is a client for the ThirdAuth schema.
type ThirdAuthClient struct {
	config
}

// NewThirdAuthClient returns a client for the ThirdAuth from the given config.
func NewThirdAuthClient(c config) *ThirdAuthClient {
	return &ThirdAuthClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `thirdauth.Hooks(f(g(h())))`.
func (c *ThirdAuthClient) Use(hooks ...Hook) {
	c.hooks.ThirdAuth = append(c.hooks.ThirdAuth, hooks...)
}

// Create returns a create builder for ThirdAuth.
func (c *ThirdAuthClient) Create() *ThirdAuthCreate {
	mutation := newThirdAuthMutation(c.config, OpCreate)
	return &ThirdAuthCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ThirdAuth entities.
func (c *ThirdAuthClient) CreateBulk(builders ...*ThirdAuthCreate) *ThirdAuthCreateBulk {
	return &ThirdAuthCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ThirdAuth.
func (c *ThirdAuthClient) Update() *ThirdAuthUpdate {
	mutation := newThirdAuthMutation(c.config, OpUpdate)
	return &ThirdAuthUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ThirdAuthClient) UpdateOne(ta *ThirdAuth) *ThirdAuthUpdateOne {
	mutation := newThirdAuthMutation(c.config, OpUpdateOne, withThirdAuth(ta))
	return &ThirdAuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ThirdAuthClient) UpdateOneID(id uuid.UUID) *ThirdAuthUpdateOne {
	mutation := newThirdAuthMutation(c.config, OpUpdateOne, withThirdAuthID(id))
	return &ThirdAuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ThirdAuth.
func (c *ThirdAuthClient) Delete() *ThirdAuthDelete {
	mutation := newThirdAuthMutation(c.config, OpDelete)
	return &ThirdAuthDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ThirdAuthClient) DeleteOne(ta *ThirdAuth) *ThirdAuthDeleteOne {
	return c.DeleteOneID(ta.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ThirdAuthClient) DeleteOneID(id uuid.UUID) *ThirdAuthDeleteOne {
	builder := c.Delete().Where(thirdauth.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ThirdAuthDeleteOne{builder}
}

// Query returns a query builder for ThirdAuth.
func (c *ThirdAuthClient) Query() *ThirdAuthQuery {
	return &ThirdAuthQuery{
		config: c.config,
	}
}

// Get returns a ThirdAuth entity by its id.
func (c *ThirdAuthClient) Get(ctx context.Context, id uuid.UUID) (*ThirdAuth, error) {
	return c.Query().Where(thirdauth.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ThirdAuthClient) GetX(ctx context.Context, id uuid.UUID) *ThirdAuth {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ThirdAuthClient) Hooks() []Hook {
	hooks := c.hooks.ThirdAuth
	return append(hooks[:len(hooks):len(hooks)], thirdauth.Hooks[:]...)
}
