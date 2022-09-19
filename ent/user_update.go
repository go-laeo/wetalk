// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-laeo/wetalk/ent/coin"
	"github.com/go-laeo/wetalk/ent/post"
	"github.com/go-laeo/wetalk/ent/predicate"
	"github.com/go-laeo/wetalk/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetAccount sets the "account" field.
func (uu *UserUpdate) SetAccount(s string) *UserUpdate {
	uu.mutation.SetAccount(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetSalt sets the "salt" field.
func (uu *UserUpdate) SetSalt(s string) *UserUpdate {
	uu.mutation.SetSalt(s)
	return uu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSalt(s *string) *UserUpdate {
	if s != nil {
		uu.SetSalt(*s)
	}
	return uu
}

// ClearSalt clears the value of the "salt" field.
func (uu *UserUpdate) ClearSalt() *UserUpdate {
	uu.mutation.ClearSalt()
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *UserUpdate) SetAvatarURL(s string) *UserUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarURL(*s)
	}
	return uu
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uu *UserUpdate) ClearAvatarURL() *UserUpdate {
	uu.mutation.ClearAvatarURL()
	return uu
}

// SetIntro sets the "intro" field.
func (uu *UserUpdate) SetIntro(s string) *UserUpdate {
	uu.mutation.SetIntro(s)
	return uu
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIntro(s *string) *UserUpdate {
	if s != nil {
		uu.SetIntro(*s)
	}
	return uu
}

// ClearIntro clears the value of the "intro" field.
func (uu *UserUpdate) ClearIntro() *UserUpdate {
	uu.mutation.ClearIntro()
	return uu
}

// SetCoin sets the "coin" field.
func (uu *UserUpdate) SetCoin(i int64) *UserUpdate {
	uu.mutation.ResetCoin()
	uu.mutation.SetCoin(i)
	return uu
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCoin(i *int64) *UserUpdate {
	if i != nil {
		uu.SetCoin(*i)
	}
	return uu
}

// AddCoin adds i to the "coin" field.
func (uu *UserUpdate) AddCoin(i int64) *UserUpdate {
	uu.mutation.AddCoin(i)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uu *UserUpdate) AddPostIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPostIDs(ids...)
	return uu
}

// AddPosts adds the "posts" edges to the Post entity.
func (uu *UserUpdate) AddPosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPostIDs(ids...)
}

// AddFavoritePostIDs adds the "favorite_posts" edge to the Post entity by IDs.
func (uu *UserUpdate) AddFavoritePostIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoritePostIDs(ids...)
	return uu
}

// AddFavoritePosts adds the "favorite_posts" edges to the Post entity.
func (uu *UserUpdate) AddFavoritePosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddFavoritePostIDs(ids...)
}

// AddCoinIDs adds the "coins" edge to the Coin entity by IDs.
func (uu *UserUpdate) AddCoinIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCoinIDs(ids...)
	return uu
}

// AddCoins adds the "coins" edges to the Coin entity.
func (uu *UserUpdate) AddCoins(c ...*Coin) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCoinIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uu *UserUpdate) ClearPosts() *UserUpdate {
	uu.mutation.ClearPosts()
	return uu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uu *UserUpdate) RemovePostIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePostIDs(ids...)
	return uu
}

// RemovePosts removes "posts" edges to Post entities.
func (uu *UserUpdate) RemovePosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePostIDs(ids...)
}

// ClearFavoritePosts clears all "favorite_posts" edges to the Post entity.
func (uu *UserUpdate) ClearFavoritePosts() *UserUpdate {
	uu.mutation.ClearFavoritePosts()
	return uu
}

// RemoveFavoritePostIDs removes the "favorite_posts" edge to Post entities by IDs.
func (uu *UserUpdate) RemoveFavoritePostIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoritePostIDs(ids...)
	return uu
}

// RemoveFavoritePosts removes "favorite_posts" edges to Post entities.
func (uu *UserUpdate) RemoveFavoritePosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveFavoritePostIDs(ids...)
}

// ClearCoins clears all "coins" edges to the Coin entity.
func (uu *UserUpdate) ClearCoins() *UserUpdate {
	uu.mutation.ClearCoins()
	return uu
}

// RemoveCoinIDs removes the "coins" edge to Coin entities by IDs.
func (uu *UserUpdate) RemoveCoinIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCoinIDs(ids...)
	return uu
}

// RemoveCoins removes "coins" edges to Coin entities.
func (uu *UserUpdate) RemoveCoins(c ...*Coin) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCoinIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAccount,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalt,
		})
	}
	if uu.mutation.SaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSalt,
		})
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uu.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uu.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldIntro,
		})
	}
	if uu.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldIntro,
		})
	}
	if value, ok := uu.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCoin,
		})
	}
	if value, ok := uu.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCoin,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FavoritePostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFavoritePostsIDs(); len(nodes) > 0 && !uu.mutation.FavoritePostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FavoritePostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CoinsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCoinsIDs(); len(nodes) > 0 && !uu.mutation.CoinsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CoinsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetAccount sets the "account" field.
func (uuo *UserUpdateOne) SetAccount(s string) *UserUpdateOne {
	uuo.mutation.SetAccount(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetSalt sets the "salt" field.
func (uuo *UserUpdateOne) SetSalt(s string) *UserUpdateOne {
	uuo.mutation.SetSalt(s)
	return uuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSalt(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSalt(*s)
	}
	return uuo
}

// ClearSalt clears the value of the "salt" field.
func (uuo *UserUpdateOne) ClearSalt() *UserUpdateOne {
	uuo.mutation.ClearSalt()
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *UserUpdateOne) SetAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarURL(*s)
	}
	return uuo
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uuo *UserUpdateOne) ClearAvatarURL() *UserUpdateOne {
	uuo.mutation.ClearAvatarURL()
	return uuo
}

// SetIntro sets the "intro" field.
func (uuo *UserUpdateOne) SetIntro(s string) *UserUpdateOne {
	uuo.mutation.SetIntro(s)
	return uuo
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIntro(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIntro(*s)
	}
	return uuo
}

// ClearIntro clears the value of the "intro" field.
func (uuo *UserUpdateOne) ClearIntro() *UserUpdateOne {
	uuo.mutation.ClearIntro()
	return uuo
}

// SetCoin sets the "coin" field.
func (uuo *UserUpdateOne) SetCoin(i int64) *UserUpdateOne {
	uuo.mutation.ResetCoin()
	uuo.mutation.SetCoin(i)
	return uuo
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCoin(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetCoin(*i)
	}
	return uuo
}

// AddCoin adds i to the "coin" field.
func (uuo *UserUpdateOne) AddCoin(i int64) *UserUpdateOne {
	uuo.mutation.AddCoin(i)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uuo *UserUpdateOne) AddPostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPostIDs(ids...)
	return uuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (uuo *UserUpdateOne) AddPosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPostIDs(ids...)
}

// AddFavoritePostIDs adds the "favorite_posts" edge to the Post entity by IDs.
func (uuo *UserUpdateOne) AddFavoritePostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoritePostIDs(ids...)
	return uuo
}

// AddFavoritePosts adds the "favorite_posts" edges to the Post entity.
func (uuo *UserUpdateOne) AddFavoritePosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddFavoritePostIDs(ids...)
}

// AddCoinIDs adds the "coins" edge to the Coin entity by IDs.
func (uuo *UserUpdateOne) AddCoinIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCoinIDs(ids...)
	return uuo
}

// AddCoins adds the "coins" edges to the Coin entity.
func (uuo *UserUpdateOne) AddCoins(c ...*Coin) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCoinIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uuo *UserUpdateOne) ClearPosts() *UserUpdateOne {
	uuo.mutation.ClearPosts()
	return uuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uuo *UserUpdateOne) RemovePostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePostIDs(ids...)
	return uuo
}

// RemovePosts removes "posts" edges to Post entities.
func (uuo *UserUpdateOne) RemovePosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePostIDs(ids...)
}

// ClearFavoritePosts clears all "favorite_posts" edges to the Post entity.
func (uuo *UserUpdateOne) ClearFavoritePosts() *UserUpdateOne {
	uuo.mutation.ClearFavoritePosts()
	return uuo
}

// RemoveFavoritePostIDs removes the "favorite_posts" edge to Post entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoritePostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoritePostIDs(ids...)
	return uuo
}

// RemoveFavoritePosts removes "favorite_posts" edges to Post entities.
func (uuo *UserUpdateOne) RemoveFavoritePosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveFavoritePostIDs(ids...)
}

// ClearCoins clears all "coins" edges to the Coin entity.
func (uuo *UserUpdateOne) ClearCoins() *UserUpdateOne {
	uuo.mutation.ClearCoins()
	return uuo
}

// RemoveCoinIDs removes the "coins" edge to Coin entities by IDs.
func (uuo *UserUpdateOne) RemoveCoinIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCoinIDs(ids...)
	return uuo
}

// RemoveCoins removes "coins" edges to Coin entities.
func (uuo *UserUpdateOne) RemoveCoins(c ...*Coin) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCoinIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAccount,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalt,
		})
	}
	if uuo.mutation.SaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSalt,
		})
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uuo.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uuo.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldIntro,
		})
	}
	if uuo.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldIntro,
		})
	}
	if value, ok := uuo.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCoin,
		})
	}
	if value, ok := uuo.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCoin,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FavoritePostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFavoritePostsIDs(); len(nodes) > 0 && !uuo.mutation.FavoritePostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FavoritePostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoritePostsTable,
			Columns: user.FavoritePostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CoinsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCoinsIDs(); len(nodes) > 0 && !uuo.mutation.CoinsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CoinsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CoinsTable,
			Columns: []string{user.CoinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
