// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mybanks-api/ent/bank"
	"mybanks-api/ent/currencyrate"
	"mybanks-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CurrencyRateUpdate is the builder for updating CurrencyRate entities.
type CurrencyRateUpdate struct {
	config
	hooks    []Hook
	mutation *CurrencyRateMutation
}

// Where appends a list predicates to the CurrencyRateUpdate builder.
func (cru *CurrencyRateUpdate) Where(ps ...predicate.CurrencyRate) *CurrencyRateUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetCurrency sets the "currency" field.
func (cru *CurrencyRateUpdate) SetCurrency(s string) *CurrencyRateUpdate {
	cru.mutation.SetCurrency(s)
	return cru
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (cru *CurrencyRateUpdate) SetNillableCurrency(s *string) *CurrencyRateUpdate {
	if s != nil {
		cru.SetCurrency(*s)
	}
	return cru
}

// SetRate sets the "rate" field.
func (cru *CurrencyRateUpdate) SetRate(f float64) *CurrencyRateUpdate {
	cru.mutation.ResetRate()
	cru.mutation.SetRate(f)
	return cru
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (cru *CurrencyRateUpdate) SetNillableRate(f *float64) *CurrencyRateUpdate {
	if f != nil {
		cru.SetRate(*f)
	}
	return cru
}

// AddRate adds f to the "rate" field.
func (cru *CurrencyRateUpdate) AddRate(f float64) *CurrencyRateUpdate {
	cru.mutation.AddRate(f)
	return cru
}

// SetBankID sets the "bank" edge to the Bank entity by ID.
func (cru *CurrencyRateUpdate) SetBankID(id int) *CurrencyRateUpdate {
	cru.mutation.SetBankID(id)
	return cru
}

// SetBank sets the "bank" edge to the Bank entity.
func (cru *CurrencyRateUpdate) SetBank(b *Bank) *CurrencyRateUpdate {
	return cru.SetBankID(b.ID)
}

// Mutation returns the CurrencyRateMutation object of the builder.
func (cru *CurrencyRateUpdate) Mutation() *CurrencyRateMutation {
	return cru.mutation
}

// ClearBank clears the "bank" edge to the Bank entity.
func (cru *CurrencyRateUpdate) ClearBank() *CurrencyRateUpdate {
	cru.mutation.ClearBank()
	return cru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CurrencyRateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CurrencyRateUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CurrencyRateUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CurrencyRateUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CurrencyRateUpdate) check() error {
	if cru.mutation.BankCleared() && len(cru.mutation.BankIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CurrencyRate.bank"`)
	}
	return nil
}

func (cru *CurrencyRateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(currencyrate.Table, currencyrate.Columns, sqlgraph.NewFieldSpec(currencyrate.FieldID, field.TypeInt))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.Currency(); ok {
		_spec.SetField(currencyrate.FieldCurrency, field.TypeString, value)
	}
	if value, ok := cru.mutation.Rate(); ok {
		_spec.SetField(currencyrate.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.AddedRate(); ok {
		_spec.AddField(currencyrate.FieldRate, field.TypeFloat64, value)
	}
	if cru.mutation.BankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   currencyrate.BankTable,
			Columns: []string{currencyrate.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.BankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   currencyrate.BankTable,
			Columns: []string{currencyrate.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currencyrate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CurrencyRateUpdateOne is the builder for updating a single CurrencyRate entity.
type CurrencyRateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CurrencyRateMutation
}

// SetCurrency sets the "currency" field.
func (cruo *CurrencyRateUpdateOne) SetCurrency(s string) *CurrencyRateUpdateOne {
	cruo.mutation.SetCurrency(s)
	return cruo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (cruo *CurrencyRateUpdateOne) SetNillableCurrency(s *string) *CurrencyRateUpdateOne {
	if s != nil {
		cruo.SetCurrency(*s)
	}
	return cruo
}

// SetRate sets the "rate" field.
func (cruo *CurrencyRateUpdateOne) SetRate(f float64) *CurrencyRateUpdateOne {
	cruo.mutation.ResetRate()
	cruo.mutation.SetRate(f)
	return cruo
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (cruo *CurrencyRateUpdateOne) SetNillableRate(f *float64) *CurrencyRateUpdateOne {
	if f != nil {
		cruo.SetRate(*f)
	}
	return cruo
}

// AddRate adds f to the "rate" field.
func (cruo *CurrencyRateUpdateOne) AddRate(f float64) *CurrencyRateUpdateOne {
	cruo.mutation.AddRate(f)
	return cruo
}

// SetBankID sets the "bank" edge to the Bank entity by ID.
func (cruo *CurrencyRateUpdateOne) SetBankID(id int) *CurrencyRateUpdateOne {
	cruo.mutation.SetBankID(id)
	return cruo
}

// SetBank sets the "bank" edge to the Bank entity.
func (cruo *CurrencyRateUpdateOne) SetBank(b *Bank) *CurrencyRateUpdateOne {
	return cruo.SetBankID(b.ID)
}

// Mutation returns the CurrencyRateMutation object of the builder.
func (cruo *CurrencyRateUpdateOne) Mutation() *CurrencyRateMutation {
	return cruo.mutation
}

// ClearBank clears the "bank" edge to the Bank entity.
func (cruo *CurrencyRateUpdateOne) ClearBank() *CurrencyRateUpdateOne {
	cruo.mutation.ClearBank()
	return cruo
}

// Where appends a list predicates to the CurrencyRateUpdate builder.
func (cruo *CurrencyRateUpdateOne) Where(ps ...predicate.CurrencyRate) *CurrencyRateUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CurrencyRateUpdateOne) Select(field string, fields ...string) *CurrencyRateUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CurrencyRate entity.
func (cruo *CurrencyRateUpdateOne) Save(ctx context.Context) (*CurrencyRate, error) {
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CurrencyRateUpdateOne) SaveX(ctx context.Context) *CurrencyRate {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CurrencyRateUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CurrencyRateUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CurrencyRateUpdateOne) check() error {
	if cruo.mutation.BankCleared() && len(cruo.mutation.BankIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CurrencyRate.bank"`)
	}
	return nil
}

func (cruo *CurrencyRateUpdateOne) sqlSave(ctx context.Context) (_node *CurrencyRate, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(currencyrate.Table, currencyrate.Columns, sqlgraph.NewFieldSpec(currencyrate.FieldID, field.TypeInt))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CurrencyRate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, currencyrate.FieldID)
		for _, f := range fields {
			if !currencyrate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != currencyrate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.Currency(); ok {
		_spec.SetField(currencyrate.FieldCurrency, field.TypeString, value)
	}
	if value, ok := cruo.mutation.Rate(); ok {
		_spec.SetField(currencyrate.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.AddedRate(); ok {
		_spec.AddField(currencyrate.FieldRate, field.TypeFloat64, value)
	}
	if cruo.mutation.BankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   currencyrate.BankTable,
			Columns: []string{currencyrate.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.BankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   currencyrate.BankTable,
			Columns: []string{currencyrate.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CurrencyRate{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currencyrate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
