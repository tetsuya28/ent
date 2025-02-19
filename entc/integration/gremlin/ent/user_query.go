// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	inters        []Interceptor
	predicates    []predicate.User
	withCard      *CardQuery
	withPets      *PetQuery
	withFiles     *FileQuery
	withGroups    *GroupQuery
	withFriends   *UserQuery
	withFollowers *UserQuery
	withFollowing *UserQuery
	withTeam      *PetQuery
	withSpouse    *UserQuery
	withChildren  *UserQuery
	withParent    *UserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...OrderFunc) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryCard chains the current query on the "card" edge.
func (uq *UserQuery) QueryCard() *CardQuery {
	query := (&CardClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.CardLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryPets chains the current query on the "pets" edge.
func (uq *UserQuery) QueryPets() *PetQuery {
	query := (&PetClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.PetsLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (uq *UserQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.FilesLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (uq *UserQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.GroupsLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryFriends chains the current query on the "friends" edge.
func (uq *UserQuery) QueryFriends() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.Both(user.FriendsLabel)
		return fromU, nil
	}
	return query
}

// QueryFollowers chains the current query on the "followers" edge.
func (uq *UserQuery) QueryFollowers() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.FollowingLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryFollowing chains the current query on the "following" edge.
func (uq *UserQuery) QueryFollowing() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.FollowingLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryTeam chains the current query on the "team" edge.
func (uq *UserQuery) QueryTeam() *PetQuery {
	query := (&PetClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.TeamLabel).InV()
		return fromU, nil
	}
	return query
}

// QuerySpouse chains the current query on the "spouse" edge.
func (uq *UserQuery) QuerySpouse() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.Both(user.SpouseLabel)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (uq *UserQuery) QueryChildren() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.ParentLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (uq *UserQuery) QueryParent() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := uq.gremlinQuery(ctx)
		fromU = gremlin.OutE(user.ParentLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(newQueryContext(ctx, TypeUser, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(1).IDs(newQueryContext(ctx, TypeUser, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) string {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(newQueryContext(ctx, TypeUser, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(2).IDs(newQueryContext(ctx, TypeUser, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) string {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = newQueryContext(ctx, TypeUser, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = newQueryContext(ctx, TypeUser, "IDs")
	if err := uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []string {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeUser, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeUser, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:        uq.config,
		limit:         uq.limit,
		offset:        uq.offset,
		order:         append([]OrderFunc{}, uq.order...),
		inters:        append([]Interceptor{}, uq.inters...),
		predicates:    append([]predicate.User{}, uq.predicates...),
		withCard:      uq.withCard.Clone(),
		withPets:      uq.withPets.Clone(),
		withFiles:     uq.withFiles.Clone(),
		withGroups:    uq.withGroups.Clone(),
		withFriends:   uq.withFriends.Clone(),
		withFollowers: uq.withFollowers.Clone(),
		withFollowing: uq.withFollowing.Clone(),
		withTeam:      uq.withTeam.Clone(),
		withSpouse:    uq.withSpouse.Clone(),
		withChildren:  uq.withChildren.Clone(),
		withParent:    uq.withParent.Clone(),
		// clone intermediate query.
		gremlin: uq.gremlin.Clone(),
		path:    uq.path,
		unique:  uq.unique,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCard(opts ...func(*CardQuery)) *UserQuery {
	query := (&CardClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCard = query
	return uq
}

// WithPets tells the query-builder to eager-load the nodes that are connected to
// the "pets" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPets(opts ...func(*PetQuery)) *UserQuery {
	query := (&PetClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPets = query
	return uq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFiles(opts ...func(*FileQuery)) *UserQuery {
	query := (&FileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFiles = query
	return uq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithGroups(opts ...func(*GroupQuery)) *UserQuery {
	query := (&GroupClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withGroups = query
	return uq
}

// WithFriends tells the query-builder to eager-load the nodes that are connected to
// the "friends" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFriends(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFriends = query
	return uq
}

// WithFollowers tells the query-builder to eager-load the nodes that are connected to
// the "followers" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFollowers(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFollowers = query
	return uq
}

// WithFollowing tells the query-builder to eager-load the nodes that are connected to
// the "following" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFollowing(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFollowing = query
	return uq
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithTeam(opts ...func(*PetQuery)) *UserQuery {
	query := (&PetClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withTeam = query
	return uq
}

// WithSpouse tells the query-builder to eager-load the nodes that are connected to
// the "spouse" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithSpouse(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withSpouse = query
	return uq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithChildren(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withChildren = query
	return uq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithParent(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withParent = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OptionalInt int `json:"optional_int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldOptionalInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OptionalInt int `json:"optional_int,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldOptionalInt).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.fields = append(uq.fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.gremlin = prev
	}
	return nil
}

func (uq *UserQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	res := &gremlin.Response{}
	traversal := uq.gremlinQuery(ctx)
	if len(uq.fields) > 0 {
		fields := make([]any, len(uq.fields))
		for i, f := range uq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := uq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var us Users
	if err := us.FromResponse(res); err != nil {
		return nil, err
	}
	us.config(uq.config)
	return us, nil
}

func (uq *UserQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := uq.gremlinQuery(ctx).Count().Query()
	if err := uq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (uq *UserQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(user.Label)
	if uq.gremlin != nil {
		v = uq.gremlin.Clone()
	}
	for _, p := range uq.predicates {
		p(v)
	}
	if len(uq.order) > 0 {
		v.Order()
		for _, p := range uq.order {
			p(v)
		}
	}
	switch limit, offset := uq.limit, uq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := uq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUser, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) gremlinScan(ctx context.Context, root *UserQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range ugb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *ugb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*ugb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := ugb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*ugb.flds)+len(ugb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUser, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) gremlinScan(ctx context.Context, root *UserQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if len(us.fields) == 1 {
		if us.fields[0] != user.FieldID {
			traversal = traversal.Values(us.fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(us.fields))
		for i, f := range us.fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := us.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
