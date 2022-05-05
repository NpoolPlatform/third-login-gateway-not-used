// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/notification/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/notification/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notification/pkg/db/ent/mailbox"
	"github.com/NpoolPlatform/notification/pkg/db/ent/notification"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readuser"
	"github.com/NpoolPlatform/notification/pkg/db/ent/template"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Announcement is the client for interacting with the Announcement builders.
	Announcement *AnnouncementClient
	// MailBox is the client for interacting with the MailBox builders.
	MailBox *MailBoxClient
	// Notification is the client for interacting with the Notification builders.
	Notification *NotificationClient
	// ReadUser is the client for interacting with the ReadUser builders.
	ReadUser *ReadUserClient
	// Template is the client for interacting with the Template builders.
	Template *TemplateClient
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
	c.Announcement = NewAnnouncementClient(c.config)
	c.MailBox = NewMailBoxClient(c.config)
	c.Notification = NewNotificationClient(c.config)
	c.ReadUser = NewReadUserClient(c.config)
	c.Template = NewTemplateClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		Announcement: NewAnnouncementClient(cfg),
		MailBox:      NewMailBoxClient(cfg),
		Notification: NewNotificationClient(cfg),
		ReadUser:     NewReadUserClient(cfg),
		Template:     NewTemplateClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		Announcement: NewAnnouncementClient(cfg),
		MailBox:      NewMailBoxClient(cfg),
		Notification: NewNotificationClient(cfg),
		ReadUser:     NewReadUserClient(cfg),
		Template:     NewTemplateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Announcement.
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
	c.Announcement.Use(hooks...)
	c.MailBox.Use(hooks...)
	c.Notification.Use(hooks...)
	c.ReadUser.Use(hooks...)
	c.Template.Use(hooks...)
}

// AnnouncementClient is a client for the Announcement schema.
type AnnouncementClient struct {
	config
}

// NewAnnouncementClient returns a client for the Announcement from the given config.
func NewAnnouncementClient(c config) *AnnouncementClient {
	return &AnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `announcement.Hooks(f(g(h())))`.
func (c *AnnouncementClient) Use(hooks ...Hook) {
	c.hooks.Announcement = append(c.hooks.Announcement, hooks...)
}

// Create returns a create builder for Announcement.
func (c *AnnouncementClient) Create() *AnnouncementCreate {
	mutation := newAnnouncementMutation(c.config, OpCreate)
	return &AnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Announcement entities.
func (c *AnnouncementClient) CreateBulk(builders ...*AnnouncementCreate) *AnnouncementCreateBulk {
	return &AnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Announcement.
func (c *AnnouncementClient) Update() *AnnouncementUpdate {
	mutation := newAnnouncementMutation(c.config, OpUpdate)
	return &AnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnnouncementClient) UpdateOne(a *Announcement) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncement(a))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnnouncementClient) UpdateOneID(id uuid.UUID) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncementID(id))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Announcement.
func (c *AnnouncementClient) Delete() *AnnouncementDelete {
	mutation := newAnnouncementMutation(c.config, OpDelete)
	return &AnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AnnouncementClient) DeleteOne(a *Announcement) *AnnouncementDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AnnouncementClient) DeleteOneID(id uuid.UUID) *AnnouncementDeleteOne {
	builder := c.Delete().Where(announcement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnnouncementDeleteOne{builder}
}

// Query returns a query builder for Announcement.
func (c *AnnouncementClient) Query() *AnnouncementQuery {
	return &AnnouncementQuery{
		config: c.config,
	}
}

// Get returns a Announcement entity by its id.
func (c *AnnouncementClient) Get(ctx context.Context, id uuid.UUID) (*Announcement, error) {
	return c.Query().Where(announcement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnnouncementClient) GetX(ctx context.Context, id uuid.UUID) *Announcement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AnnouncementClient) Hooks() []Hook {
	return c.hooks.Announcement
}

// MailBoxClient is a client for the MailBox schema.
type MailBoxClient struct {
	config
}

// NewMailBoxClient returns a client for the MailBox from the given config.
func NewMailBoxClient(c config) *MailBoxClient {
	return &MailBoxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailbox.Hooks(f(g(h())))`.
func (c *MailBoxClient) Use(hooks ...Hook) {
	c.hooks.MailBox = append(c.hooks.MailBox, hooks...)
}

// Create returns a create builder for MailBox.
func (c *MailBoxClient) Create() *MailBoxCreate {
	mutation := newMailBoxMutation(c.config, OpCreate)
	return &MailBoxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MailBox entities.
func (c *MailBoxClient) CreateBulk(builders ...*MailBoxCreate) *MailBoxCreateBulk {
	return &MailBoxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MailBox.
func (c *MailBoxClient) Update() *MailBoxUpdate {
	mutation := newMailBoxMutation(c.config, OpUpdate)
	return &MailBoxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailBoxClient) UpdateOne(mb *MailBox) *MailBoxUpdateOne {
	mutation := newMailBoxMutation(c.config, OpUpdateOne, withMailBox(mb))
	return &MailBoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailBoxClient) UpdateOneID(id uuid.UUID) *MailBoxUpdateOne {
	mutation := newMailBoxMutation(c.config, OpUpdateOne, withMailBoxID(id))
	return &MailBoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MailBox.
func (c *MailBoxClient) Delete() *MailBoxDelete {
	mutation := newMailBoxMutation(c.config, OpDelete)
	return &MailBoxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailBoxClient) DeleteOne(mb *MailBox) *MailBoxDeleteOne {
	return c.DeleteOneID(mb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailBoxClient) DeleteOneID(id uuid.UUID) *MailBoxDeleteOne {
	builder := c.Delete().Where(mailbox.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailBoxDeleteOne{builder}
}

// Query returns a query builder for MailBox.
func (c *MailBoxClient) Query() *MailBoxQuery {
	return &MailBoxQuery{
		config: c.config,
	}
}

// Get returns a MailBox entity by its id.
func (c *MailBoxClient) Get(ctx context.Context, id uuid.UUID) (*MailBox, error) {
	return c.Query().Where(mailbox.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailBoxClient) GetX(ctx context.Context, id uuid.UUID) *MailBox {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MailBoxClient) Hooks() []Hook {
	return c.hooks.MailBox
}

// NotificationClient is a client for the Notification schema.
type NotificationClient struct {
	config
}

// NewNotificationClient returns a client for the Notification from the given config.
func NewNotificationClient(c config) *NotificationClient {
	return &NotificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notification.Hooks(f(g(h())))`.
func (c *NotificationClient) Use(hooks ...Hook) {
	c.hooks.Notification = append(c.hooks.Notification, hooks...)
}

// Create returns a create builder for Notification.
func (c *NotificationClient) Create() *NotificationCreate {
	mutation := newNotificationMutation(c.config, OpCreate)
	return &NotificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notification entities.
func (c *NotificationClient) CreateBulk(builders ...*NotificationCreate) *NotificationCreateBulk {
	return &NotificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notification.
func (c *NotificationClient) Update() *NotificationUpdate {
	mutation := newNotificationMutation(c.config, OpUpdate)
	return &NotificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotificationClient) UpdateOne(n *Notification) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotification(n))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotificationClient) UpdateOneID(id uuid.UUID) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotificationID(id))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notification.
func (c *NotificationClient) Delete() *NotificationDelete {
	mutation := newNotificationMutation(c.config, OpDelete)
	return &NotificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NotificationClient) DeleteOne(n *Notification) *NotificationDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NotificationClient) DeleteOneID(id uuid.UUID) *NotificationDeleteOne {
	builder := c.Delete().Where(notification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotificationDeleteOne{builder}
}

// Query returns a query builder for Notification.
func (c *NotificationClient) Query() *NotificationQuery {
	return &NotificationQuery{
		config: c.config,
	}
}

// Get returns a Notification entity by its id.
func (c *NotificationClient) Get(ctx context.Context, id uuid.UUID) (*Notification, error) {
	return c.Query().Where(notification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotificationClient) GetX(ctx context.Context, id uuid.UUID) *Notification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotificationClient) Hooks() []Hook {
	return c.hooks.Notification
}

// ReadUserClient is a client for the ReadUser schema.
type ReadUserClient struct {
	config
}

// NewReadUserClient returns a client for the ReadUser from the given config.
func NewReadUserClient(c config) *ReadUserClient {
	return &ReadUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `readuser.Hooks(f(g(h())))`.
func (c *ReadUserClient) Use(hooks ...Hook) {
	c.hooks.ReadUser = append(c.hooks.ReadUser, hooks...)
}

// Create returns a create builder for ReadUser.
func (c *ReadUserClient) Create() *ReadUserCreate {
	mutation := newReadUserMutation(c.config, OpCreate)
	return &ReadUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReadUser entities.
func (c *ReadUserClient) CreateBulk(builders ...*ReadUserCreate) *ReadUserCreateBulk {
	return &ReadUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReadUser.
func (c *ReadUserClient) Update() *ReadUserUpdate {
	mutation := newReadUserMutation(c.config, OpUpdate)
	return &ReadUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReadUserClient) UpdateOne(ru *ReadUser) *ReadUserUpdateOne {
	mutation := newReadUserMutation(c.config, OpUpdateOne, withReadUser(ru))
	return &ReadUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReadUserClient) UpdateOneID(id uuid.UUID) *ReadUserUpdateOne {
	mutation := newReadUserMutation(c.config, OpUpdateOne, withReadUserID(id))
	return &ReadUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReadUser.
func (c *ReadUserClient) Delete() *ReadUserDelete {
	mutation := newReadUserMutation(c.config, OpDelete)
	return &ReadUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReadUserClient) DeleteOne(ru *ReadUser) *ReadUserDeleteOne {
	return c.DeleteOneID(ru.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReadUserClient) DeleteOneID(id uuid.UUID) *ReadUserDeleteOne {
	builder := c.Delete().Where(readuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReadUserDeleteOne{builder}
}

// Query returns a query builder for ReadUser.
func (c *ReadUserClient) Query() *ReadUserQuery {
	return &ReadUserQuery{
		config: c.config,
	}
}

// Get returns a ReadUser entity by its id.
func (c *ReadUserClient) Get(ctx context.Context, id uuid.UUID) (*ReadUser, error) {
	return c.Query().Where(readuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReadUserClient) GetX(ctx context.Context, id uuid.UUID) *ReadUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReadUserClient) Hooks() []Hook {
	return c.hooks.ReadUser
}

// TemplateClient is a client for the Template schema.
type TemplateClient struct {
	config
}

// NewTemplateClient returns a client for the Template from the given config.
func NewTemplateClient(c config) *TemplateClient {
	return &TemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `template.Hooks(f(g(h())))`.
func (c *TemplateClient) Use(hooks ...Hook) {
	c.hooks.Template = append(c.hooks.Template, hooks...)
}

// Create returns a create builder for Template.
func (c *TemplateClient) Create() *TemplateCreate {
	mutation := newTemplateMutation(c.config, OpCreate)
	return &TemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Template entities.
func (c *TemplateClient) CreateBulk(builders ...*TemplateCreate) *TemplateCreateBulk {
	return &TemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Template.
func (c *TemplateClient) Update() *TemplateUpdate {
	mutation := newTemplateMutation(c.config, OpUpdate)
	return &TemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TemplateClient) UpdateOne(t *Template) *TemplateUpdateOne {
	mutation := newTemplateMutation(c.config, OpUpdateOne, withTemplate(t))
	return &TemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TemplateClient) UpdateOneID(id uuid.UUID) *TemplateUpdateOne {
	mutation := newTemplateMutation(c.config, OpUpdateOne, withTemplateID(id))
	return &TemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Template.
func (c *TemplateClient) Delete() *TemplateDelete {
	mutation := newTemplateMutation(c.config, OpDelete)
	return &TemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TemplateClient) DeleteOne(t *Template) *TemplateDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TemplateClient) DeleteOneID(id uuid.UUID) *TemplateDeleteOne {
	builder := c.Delete().Where(template.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TemplateDeleteOne{builder}
}

// Query returns a query builder for Template.
func (c *TemplateClient) Query() *TemplateQuery {
	return &TemplateQuery{
		config: c.config,
	}
}

// Get returns a Template entity by its id.
func (c *TemplateClient) Get(ctx context.Context, id uuid.UUID) (*Template, error) {
	return c.Query().Where(template.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TemplateClient) GetX(ctx context.Context, id uuid.UUID) *Template {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TemplateClient) Hooks() []Hook {
	return c.hooks.Template
}
