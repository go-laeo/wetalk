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
	"github.com/go-laeo/wetalk/ent/user"
)

// CoinCreate is the builder for creating a Coin entity.
type CoinCreate struct {
	config
	mutation *CoinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDeal sets the "deal" field.
func (cc *CoinCreate) SetDeal(s string) *CoinCreate {
	cc.mutation.SetDeal(s)
	return cc
}

// SetAmount sets the "amount" field.
func (cc *CoinCreate) SetAmount(i int64) *CoinCreate {
	cc.mutation.SetAmount(i)
	return cc
}

// SetBalance sets the "balance" field.
func (cc *CoinCreate) SetBalance(i int64) *CoinCreate {
	cc.mutation.SetBalance(i)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CoinCreate) SetCreatedAt(t time.Time) *CoinCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CoinCreate) SetNillableCreatedAt(t *time.Time) *CoinCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CoinCreate) SetUpdatedAt(t time.Time) *CoinCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CoinCreate) SetNillableUpdatedAt(t *time.Time) *CoinCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CoinCreate) SetOwnerID(id int) *CoinCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CoinCreate) SetOwner(u *User) *CoinCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the CoinMutation object of the builder.
func (cc *CoinCreate) Mutation() *CoinMutation {
	return cc.mutation
}

// Save creates the Coin in the database.
func (cc *CoinCreate) Save(ctx context.Context) (*Coin, error) {
	var (
		err  error
		node *Coin
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Coin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CoinCreate) SaveX(ctx context.Context) *Coin {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CoinCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CoinCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CoinCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := coin.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := coin.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CoinCreate) check() error {
	if _, ok := cc.mutation.Deal(); !ok {
		return &ValidationError{Name: "deal", err: errors.New(`ent: missing required field "Coin.deal"`)}
	}
	if _, ok := cc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Coin.amount"`)}
	}
	if _, ok := cc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "Coin.balance"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Coin.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Coin.updated_at"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Coin.owner"`)}
	}
	return nil
}

func (cc *CoinCreate) sqlSave(ctx context.Context) (*Coin, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CoinCreate) createSpec() (*Coin, *sqlgraph.CreateSpec) {
	var (
		_node = &Coin{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coin.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.Deal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldDeal,
		})
		_node.Deal = value
	}
	if value, ok := cc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coin.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := cc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coin.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coin.OwnerTable,
			Columns: []string{coin.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_coins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Coin.Create().
//		SetDeal(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinUpsert) {
//			SetDeal(v+v).
//		}).
//		Exec(ctx)
func (cc *CoinCreate) OnConflict(opts ...sql.ConflictOption) *CoinUpsertOne {
	cc.conflict = opts
	return &CoinUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CoinCreate) OnConflictColumns(columns ...string) *CoinUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CoinUpsertOne{
		create: cc,
	}
}

type (
	// CoinUpsertOne is the builder for "upsert"-ing
	//  one Coin node.
	CoinUpsertOne struct {
		create *CoinCreate
	}

	// CoinUpsert is the "OnConflict" setter.
	CoinUpsert struct {
		*sql.UpdateSet
	}
)

// SetDeal sets the "deal" field.
func (u *CoinUpsert) SetDeal(v string) *CoinUpsert {
	u.Set(coin.FieldDeal, v)
	return u
}

// UpdateDeal sets the "deal" field to the value that was provided on create.
func (u *CoinUpsert) UpdateDeal() *CoinUpsert {
	u.SetExcluded(coin.FieldDeal)
	return u
}

// SetAmount sets the "amount" field.
func (u *CoinUpsert) SetAmount(v int64) *CoinUpsert {
	u.Set(coin.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CoinUpsert) UpdateAmount() *CoinUpsert {
	u.SetExcluded(coin.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *CoinUpsert) AddAmount(v int64) *CoinUpsert {
	u.Add(coin.FieldAmount, v)
	return u
}

// SetBalance sets the "balance" field.
func (u *CoinUpsert) SetBalance(v int64) *CoinUpsert {
	u.Set(coin.FieldBalance, v)
	return u
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CoinUpsert) UpdateBalance() *CoinUpsert {
	u.SetExcluded(coin.FieldBalance)
	return u
}

// AddBalance adds v to the "balance" field.
func (u *CoinUpsert) AddBalance(v int64) *CoinUpsert {
	u.Add(coin.FieldBalance, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsert) SetCreatedAt(v time.Time) *CoinUpsert {
	u.Set(coin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsert) UpdateCreatedAt() *CoinUpsert {
	u.SetExcluded(coin.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsert) SetUpdatedAt(v time.Time) *CoinUpsert {
	u.Set(coin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsert) UpdateUpdatedAt() *CoinUpsert {
	u.SetExcluded(coin.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CoinUpsertOne) UpdateNewValues() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Coin.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CoinUpsertOne) Ignore() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinUpsertOne) DoNothing() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinCreate.OnConflict
// documentation for more info.
func (u *CoinUpsertOne) Update(set func(*CoinUpsert)) *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeal sets the "deal" field.
func (u *CoinUpsertOne) SetDeal(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetDeal(v)
	})
}

// UpdateDeal sets the "deal" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateDeal() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateDeal()
	})
}

// SetAmount sets the "amount" field.
func (u *CoinUpsertOne) SetAmount(v int64) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *CoinUpsertOne) AddAmount(v int64) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateAmount() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateAmount()
	})
}

// SetBalance sets the "balance" field.
func (u *CoinUpsertOne) SetBalance(v int64) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CoinUpsertOne) AddBalance(v int64) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateBalance() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateBalance()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsertOne) SetCreatedAt(v time.Time) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateCreatedAt() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsertOne) SetUpdatedAt(v time.Time) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateUpdatedAt() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CoinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CoinUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CoinUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CoinCreateBulk is the builder for creating many Coin entities in bulk.
type CoinCreateBulk struct {
	config
	builders []*CoinCreate
	conflict []sql.ConflictOption
}

// Save creates the Coin entities in the database.
func (ccb *CoinCreateBulk) Save(ctx context.Context) ([]*Coin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Coin, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CoinCreateBulk) SaveX(ctx context.Context) []*Coin {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CoinCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CoinCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Coin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinUpsert) {
//			SetDeal(v+v).
//		}).
//		Exec(ctx)
func (ccb *CoinCreateBulk) OnConflict(opts ...sql.ConflictOption) *CoinUpsertBulk {
	ccb.conflict = opts
	return &CoinUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CoinCreateBulk) OnConflictColumns(columns ...string) *CoinUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CoinUpsertBulk{
		create: ccb,
	}
}

// CoinUpsertBulk is the builder for "upsert"-ing
// a bulk of Coin nodes.
type CoinUpsertBulk struct {
	create *CoinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CoinUpsertBulk) UpdateNewValues() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CoinUpsertBulk) Ignore() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinUpsertBulk) DoNothing() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinCreateBulk.OnConflict
// documentation for more info.
func (u *CoinUpsertBulk) Update(set func(*CoinUpsert)) *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeal sets the "deal" field.
func (u *CoinUpsertBulk) SetDeal(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetDeal(v)
	})
}

// UpdateDeal sets the "deal" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateDeal() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateDeal()
	})
}

// SetAmount sets the "amount" field.
func (u *CoinUpsertBulk) SetAmount(v int64) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *CoinUpsertBulk) AddAmount(v int64) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateAmount() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateAmount()
	})
}

// SetBalance sets the "balance" field.
func (u *CoinUpsertBulk) SetBalance(v int64) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CoinUpsertBulk) AddBalance(v int64) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateBalance() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateBalance()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsertBulk) SetCreatedAt(v time.Time) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateCreatedAt() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsertBulk) SetUpdatedAt(v time.Time) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateUpdatedAt() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CoinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CoinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}