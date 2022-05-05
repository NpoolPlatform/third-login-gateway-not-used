// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notification/pkg/db/ent/template"
	"github.com/google/uuid"
)

// TemplateUpdate is the builder for updating Template entities.
type TemplateUpdate struct {
	config
	hooks    []Hook
	mutation *TemplateMutation
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tu *TemplateUpdate) Where(ps ...predicate.Template) *TemplateUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetContent sets the "content" field.
func (tu *TemplateUpdate) SetContent(s string) *TemplateUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetAppID sets the "app_id" field.
func (tu *TemplateUpdate) SetAppID(u uuid.UUID) *TemplateUpdate {
	tu.mutation.SetAppID(u)
	return tu
}

// SetLangID sets the "lang_id" field.
func (tu *TemplateUpdate) SetLangID(u uuid.UUID) *TemplateUpdate {
	tu.mutation.SetLangID(u)
	return tu
}

// SetUsedFor sets the "used_for" field.
func (tu *TemplateUpdate) SetUsedFor(s string) *TemplateUpdate {
	tu.mutation.SetUsedFor(s)
	return tu
}

// SetCreateAt sets the "create_at" field.
func (tu *TemplateUpdate) SetCreateAt(u uint32) *TemplateUpdate {
	tu.mutation.ResetCreateAt()
	tu.mutation.SetCreateAt(u)
	return tu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableCreateAt(u *uint32) *TemplateUpdate {
	if u != nil {
		tu.SetCreateAt(*u)
	}
	return tu
}

// AddCreateAt adds u to the "create_at" field.
func (tu *TemplateUpdate) AddCreateAt(u int32) *TemplateUpdate {
	tu.mutation.AddCreateAt(u)
	return tu
}

// SetUpdateAt sets the "update_at" field.
func (tu *TemplateUpdate) SetUpdateAt(u uint32) *TemplateUpdate {
	tu.mutation.ResetUpdateAt()
	tu.mutation.SetUpdateAt(u)
	return tu
}

// AddUpdateAt adds u to the "update_at" field.
func (tu *TemplateUpdate) AddUpdateAt(u int32) *TemplateUpdate {
	tu.mutation.AddUpdateAt(u)
	return tu
}

// SetDeleteAt sets the "delete_at" field.
func (tu *TemplateUpdate) SetDeleteAt(u uint32) *TemplateUpdate {
	tu.mutation.ResetDeleteAt()
	tu.mutation.SetDeleteAt(u)
	return tu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableDeleteAt(u *uint32) *TemplateUpdate {
	if u != nil {
		tu.SetDeleteAt(*u)
	}
	return tu
}

// AddDeleteAt adds u to the "delete_at" field.
func (tu *TemplateUpdate) AddDeleteAt(u int32) *TemplateUpdate {
	tu.mutation.AddDeleteAt(u)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TemplateUpdate) SetTitle(s string) *TemplateUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// Mutation returns the TemplateMutation object of the builder.
func (tu *TemplateUpdate) Mutation() *TemplateMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplateUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplateUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplateUpdate) defaults() {
	if _, ok := tu.mutation.UpdateAt(); !ok {
		v := template.UpdateDefaultUpdateAt()
		tu.mutation.SetUpdateAt(v)
	}
}

func (tu *TemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: template.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldContent,
		})
	}
	if value, ok := tu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: template.FieldAppID,
		})
	}
	if value, ok := tu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: template.FieldLangID,
		})
	}
	if value, ok := tu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldUsedFor,
		})
	}
	if value, ok := tu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldCreateAt,
		})
	}
	if value, ok := tu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldCreateAt,
		})
	}
	if value, ok := tu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldUpdateAt,
		})
	}
	if value, ok := tu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldUpdateAt,
		})
	}
	if value, ok := tu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldDeleteAt,
		})
	}
	if value, ok := tu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldDeleteAt,
		})
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldTitle,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TemplateUpdateOne is the builder for updating a single Template entity.
type TemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TemplateMutation
}

// SetContent sets the "content" field.
func (tuo *TemplateUpdateOne) SetContent(s string) *TemplateUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetAppID sets the "app_id" field.
func (tuo *TemplateUpdateOne) SetAppID(u uuid.UUID) *TemplateUpdateOne {
	tuo.mutation.SetAppID(u)
	return tuo
}

// SetLangID sets the "lang_id" field.
func (tuo *TemplateUpdateOne) SetLangID(u uuid.UUID) *TemplateUpdateOne {
	tuo.mutation.SetLangID(u)
	return tuo
}

// SetUsedFor sets the "used_for" field.
func (tuo *TemplateUpdateOne) SetUsedFor(s string) *TemplateUpdateOne {
	tuo.mutation.SetUsedFor(s)
	return tuo
}

// SetCreateAt sets the "create_at" field.
func (tuo *TemplateUpdateOne) SetCreateAt(u uint32) *TemplateUpdateOne {
	tuo.mutation.ResetCreateAt()
	tuo.mutation.SetCreateAt(u)
	return tuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableCreateAt(u *uint32) *TemplateUpdateOne {
	if u != nil {
		tuo.SetCreateAt(*u)
	}
	return tuo
}

// AddCreateAt adds u to the "create_at" field.
func (tuo *TemplateUpdateOne) AddCreateAt(u int32) *TemplateUpdateOne {
	tuo.mutation.AddCreateAt(u)
	return tuo
}

// SetUpdateAt sets the "update_at" field.
func (tuo *TemplateUpdateOne) SetUpdateAt(u uint32) *TemplateUpdateOne {
	tuo.mutation.ResetUpdateAt()
	tuo.mutation.SetUpdateAt(u)
	return tuo
}

// AddUpdateAt adds u to the "update_at" field.
func (tuo *TemplateUpdateOne) AddUpdateAt(u int32) *TemplateUpdateOne {
	tuo.mutation.AddUpdateAt(u)
	return tuo
}

// SetDeleteAt sets the "delete_at" field.
func (tuo *TemplateUpdateOne) SetDeleteAt(u uint32) *TemplateUpdateOne {
	tuo.mutation.ResetDeleteAt()
	tuo.mutation.SetDeleteAt(u)
	return tuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableDeleteAt(u *uint32) *TemplateUpdateOne {
	if u != nil {
		tuo.SetDeleteAt(*u)
	}
	return tuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (tuo *TemplateUpdateOne) AddDeleteAt(u int32) *TemplateUpdateOne {
	tuo.mutation.AddDeleteAt(u)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TemplateUpdateOne) SetTitle(s string) *TemplateUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// Mutation returns the TemplateMutation object of the builder.
func (tuo *TemplateUpdateOne) Mutation() *TemplateMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TemplateUpdateOne) Select(field string, fields ...string) *TemplateUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Template entity.
func (tuo *TemplateUpdateOne) Save(ctx context.Context) (*Template, error) {
	var (
		err  error
		node *Template
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveX(ctx context.Context) *Template {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplateUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateAt(); !ok {
		v := template.UpdateDefaultUpdateAt()
		tuo.mutation.SetUpdateAt(v)
	}
}

func (tuo *TemplateUpdateOne) sqlSave(ctx context.Context) (_node *Template, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: template.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Template.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, template.FieldID)
		for _, f := range fields {
			if !template.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != template.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldContent,
		})
	}
	if value, ok := tuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: template.FieldAppID,
		})
	}
	if value, ok := tuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: template.FieldLangID,
		})
	}
	if value, ok := tuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldUsedFor,
		})
	}
	if value, ok := tuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldCreateAt,
		})
	}
	if value, ok := tuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldCreateAt,
		})
	}
	if value, ok := tuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldUpdateAt,
		})
	}
	if value, ok := tuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldUpdateAt,
		})
	}
	if value, ok := tuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldDeleteAt,
		})
	}
	if value, ok := tuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldDeleteAt,
		})
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldTitle,
		})
	}
	_node = &Template{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
