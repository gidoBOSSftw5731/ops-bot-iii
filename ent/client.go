// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/signin"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/user"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/vote"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/voteresult"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Signin is the client for interacting with the Signin builders.
	Signin *SigninClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Vote is the client for interacting with the Vote builders.
	Vote *VoteClient
	// VoteResult is the client for interacting with the VoteResult builders.
	VoteResult *VoteResultClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Signin = NewSigninClient(c.config)
	c.User = NewUserClient(c.config)
	c.Vote = NewVoteClient(c.config)
	c.VoteResult = NewVoteResultClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Signin:     NewSigninClient(cfg),
		User:       NewUserClient(cfg),
		Vote:       NewVoteClient(cfg),
		VoteResult: NewVoteResultClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:        ctx,
		config:     cfg,
		Signin:     NewSigninClient(cfg),
		User:       NewUserClient(cfg),
		Vote:       NewVoteClient(cfg),
		VoteResult: NewVoteResultClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Signin.
//		Query().
//		Count(ctx)
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
	c.Signin.Use(hooks...)
	c.User.Use(hooks...)
	c.Vote.Use(hooks...)
	c.VoteResult.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Signin.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.Vote.Intercept(interceptors...)
	c.VoteResult.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *SigninMutation:
		return c.Signin.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *VoteMutation:
		return c.Vote.mutate(ctx, m)
	case *VoteResultMutation:
		return c.VoteResult.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// SigninClient is a client for the Signin schema.
type SigninClient struct {
	config
}

// NewSigninClient returns a client for the Signin from the given config.
func NewSigninClient(c config) *SigninClient {
	return &SigninClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `signin.Hooks(f(g(h())))`.
func (c *SigninClient) Use(hooks ...Hook) {
	c.hooks.Signin = append(c.hooks.Signin, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `signin.Intercept(f(g(h())))`.
func (c *SigninClient) Intercept(interceptors ...Interceptor) {
	c.inters.Signin = append(c.inters.Signin, interceptors...)
}

// Create returns a builder for creating a Signin entity.
func (c *SigninClient) Create() *SigninCreate {
	mutation := newSigninMutation(c.config, OpCreate)
	return &SigninCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Signin entities.
func (c *SigninClient) CreateBulk(builders ...*SigninCreate) *SigninCreateBulk {
	return &SigninCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Signin.
func (c *SigninClient) Update() *SigninUpdate {
	mutation := newSigninMutation(c.config, OpUpdate)
	return &SigninUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SigninClient) UpdateOne(s *Signin) *SigninUpdateOne {
	mutation := newSigninMutation(c.config, OpUpdateOne, withSignin(s))
	return &SigninUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SigninClient) UpdateOneID(id int) *SigninUpdateOne {
	mutation := newSigninMutation(c.config, OpUpdateOne, withSigninID(id))
	return &SigninUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Signin.
func (c *SigninClient) Delete() *SigninDelete {
	mutation := newSigninMutation(c.config, OpDelete)
	return &SigninDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SigninClient) DeleteOne(s *Signin) *SigninDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SigninClient) DeleteOneID(id int) *SigninDeleteOne {
	builder := c.Delete().Where(signin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SigninDeleteOne{builder}
}

// Query returns a query builder for Signin.
func (c *SigninClient) Query() *SigninQuery {
	return &SigninQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSignin},
		inters: c.Interceptors(),
	}
}

// Get returns a Signin entity by its id.
func (c *SigninClient) Get(ctx context.Context, id int) (*Signin, error) {
	return c.Query().Where(signin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SigninClient) GetX(ctx context.Context, id int) *Signin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Signin.
func (c *SigninClient) QueryUser(s *Signin) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(signin.Table, signin.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, signin.UserTable, signin.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SigninClient) Hooks() []Hook {
	return c.hooks.Signin
}

// Interceptors returns the client interceptors.
func (c *SigninClient) Interceptors() []Interceptor {
	return c.inters.Signin
}

func (c *SigninClient) mutate(ctx context.Context, m *SigninMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SigninCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SigninUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SigninUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SigninDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Signin mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySignins queries the signins edge of a User.
func (c *UserClient) QuerySignins(u *User) *SigninQuery {
	query := (&SigninClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(signin.Table, signin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SigninsTable, user.SigninsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVotes queries the votes edge of a User.
func (c *UserClient) QueryVotes(u *User) *VoteQuery {
	query := (&VoteClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(vote.Table, vote.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.VotesTable, user.VotesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// VoteClient is a client for the Vote schema.
type VoteClient struct {
	config
}

// NewVoteClient returns a client for the Vote from the given config.
func NewVoteClient(c config) *VoteClient {
	return &VoteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vote.Hooks(f(g(h())))`.
func (c *VoteClient) Use(hooks ...Hook) {
	c.hooks.Vote = append(c.hooks.Vote, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `vote.Intercept(f(g(h())))`.
func (c *VoteClient) Intercept(interceptors ...Interceptor) {
	c.inters.Vote = append(c.inters.Vote, interceptors...)
}

// Create returns a builder for creating a Vote entity.
func (c *VoteClient) Create() *VoteCreate {
	mutation := newVoteMutation(c.config, OpCreate)
	return &VoteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vote entities.
func (c *VoteClient) CreateBulk(builders ...*VoteCreate) *VoteCreateBulk {
	return &VoteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vote.
func (c *VoteClient) Update() *VoteUpdate {
	mutation := newVoteMutation(c.config, OpUpdate)
	return &VoteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoteClient) UpdateOne(v *Vote) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVote(v))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoteClient) UpdateOneID(id int) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVoteID(id))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vote.
func (c *VoteClient) Delete() *VoteDelete {
	mutation := newVoteMutation(c.config, OpDelete)
	return &VoteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VoteClient) DeleteOne(v *Vote) *VoteDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VoteClient) DeleteOneID(id int) *VoteDeleteOne {
	builder := c.Delete().Where(vote.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoteDeleteOne{builder}
}

// Query returns a query builder for Vote.
func (c *VoteClient) Query() *VoteQuery {
	return &VoteQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVote},
		inters: c.Interceptors(),
	}
}

// Get returns a Vote entity by its id.
func (c *VoteClient) Get(ctx context.Context, id int) (*Vote, error) {
	return c.Query().Where(vote.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoteClient) GetX(ctx context.Context, id int) *Vote {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Vote.
func (c *VoteClient) QueryUser(v *Vote) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vote.UserTable, vote.UserColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VoteClient) Hooks() []Hook {
	return c.hooks.Vote
}

// Interceptors returns the client interceptors.
func (c *VoteClient) Interceptors() []Interceptor {
	return c.inters.Vote
}

func (c *VoteClient) mutate(ctx context.Context, m *VoteMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VoteCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VoteUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VoteDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Vote mutation op: %q", m.Op())
	}
}

// VoteResultClient is a client for the VoteResult schema.
type VoteResultClient struct {
	config
}

// NewVoteResultClient returns a client for the VoteResult from the given config.
func NewVoteResultClient(c config) *VoteResultClient {
	return &VoteResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `voteresult.Hooks(f(g(h())))`.
func (c *VoteResultClient) Use(hooks ...Hook) {
	c.hooks.VoteResult = append(c.hooks.VoteResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `voteresult.Intercept(f(g(h())))`.
func (c *VoteResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.VoteResult = append(c.inters.VoteResult, interceptors...)
}

// Create returns a builder for creating a VoteResult entity.
func (c *VoteResultClient) Create() *VoteResultCreate {
	mutation := newVoteResultMutation(c.config, OpCreate)
	return &VoteResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of VoteResult entities.
func (c *VoteResultClient) CreateBulk(builders ...*VoteResultCreate) *VoteResultCreateBulk {
	return &VoteResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for VoteResult.
func (c *VoteResultClient) Update() *VoteResultUpdate {
	mutation := newVoteResultMutation(c.config, OpUpdate)
	return &VoteResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoteResultClient) UpdateOne(vr *VoteResult) *VoteResultUpdateOne {
	mutation := newVoteResultMutation(c.config, OpUpdateOne, withVoteResult(vr))
	return &VoteResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoteResultClient) UpdateOneID(id int) *VoteResultUpdateOne {
	mutation := newVoteResultMutation(c.config, OpUpdateOne, withVoteResultID(id))
	return &VoteResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for VoteResult.
func (c *VoteResultClient) Delete() *VoteResultDelete {
	mutation := newVoteResultMutation(c.config, OpDelete)
	return &VoteResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VoteResultClient) DeleteOne(vr *VoteResult) *VoteResultDeleteOne {
	return c.DeleteOneID(vr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VoteResultClient) DeleteOneID(id int) *VoteResultDeleteOne {
	builder := c.Delete().Where(voteresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoteResultDeleteOne{builder}
}

// Query returns a query builder for VoteResult.
func (c *VoteResultClient) Query() *VoteResultQuery {
	return &VoteResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVoteResult},
		inters: c.Interceptors(),
	}
}

// Get returns a VoteResult entity by its id.
func (c *VoteResultClient) Get(ctx context.Context, id int) (*VoteResult, error) {
	return c.Query().Where(voteresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoteResultClient) GetX(ctx context.Context, id int) *VoteResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VoteResultClient) Hooks() []Hook {
	return c.hooks.VoteResult
}

// Interceptors returns the client interceptors.
func (c *VoteResultClient) Interceptors() []Interceptor {
	return c.inters.VoteResult
}

func (c *VoteResultClient) mutate(ctx context.Context, m *VoteResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VoteResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VoteResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VoteResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VoteResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown VoteResult mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Signin, User, Vote, VoteResult []ent.Hook
	}
	inters struct {
		Signin, User, Vote, VoteResult []ent.Interceptor
	}
)
