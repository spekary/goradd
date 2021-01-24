package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/goradd/goradd/pkg/orm/broadcast"
	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"
	"github.com/goradd/goradd/pkg/stringmap"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"

	//"./node"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// giftBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type giftBase struct {
	number        int
	numberIsValid bool
	numberIsDirty bool

	name        string
	nameIsValid bool
	nameIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool

	// The original primary key for updates
	_originalPK int
}

const (
	GiftNumberDefault = 0
	GiftNameDefault   = ""
)

const (
	Gift_Number = `Number`
	Gift_Name   = `Name`
)

// Initialize or re-initialize a Gift database object to default values.
func (o *giftBase) Initialize() {

	o.number = 0
	o.numberIsValid = false
	o.numberIsDirty = false

	o.name = ""
	o.nameIsValid = false
	o.nameIsDirty = false

	o._restored = false
}

func (o *giftBase) PrimaryKey() int {
	return o.number
}

// Number returns the loaded value of Number.
func (o *giftBase) Number() int {
	if o._restored && !o.numberIsValid {
		panic("number was not selected in the last query and has not been set, and so is not valid")
	}
	return o.number
}

// NumberIsValid returns true if the value was loaded from the database or has been set.
func (o *giftBase) NumberIsValid() bool {
	return o.numberIsValid
}

// SetNumber sets the value of Number in the object, to be saved later using the Save() function.
func (o *giftBase) SetNumber(v int) {
	o.numberIsValid = true
	if o.number != v || !o._restored {
		o.number = v
		o.numberIsDirty = true
	}

}

// Name returns the loaded value of Name.
func (o *giftBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *giftBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *giftBase) SetName(v string) {
	o.nameIsValid = true
	if o.name != v || !o._restored {
		o.name = v
		o.nameIsDirty = true
	}

}

// GetAlias returns the alias for the given key.
func (o *giftBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// Load returns a Gift from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func LoadGift(ctx context.Context, primaryKey int, joinOrSelectNodes ...query.NodeI) *Gift {
	return queryGifts(ctx).Where(Equal(node.Gift().Number(), primaryKey)).joinOrSelect(joinOrSelectNodes...).Get()
}

// LoadGiftByNumber queries for a single Gift object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryGifts() to start a query builder.
func LoadGiftByNumber(ctx context.Context, number int, joinOrSelectNodes ...query.NodeI) *Gift {
	q := queryGifts(ctx)
	q = q.Where(Equal(node.Gift().Number(), number))
	return q.
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasGiftByNumber returns true if the
// given unique index values exist in the database.
func HasGiftByNumber(ctx context.Context, number int) bool {
	q := queryGifts(ctx)
	q = q.Where(Equal(node.Gift().Number(), number))
	return q.Count(false) == 1
}

// The GiftsBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type GiftsBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newGiftBuilder(ctx context.Context) *GiftsBuilder {
	b := &GiftsBuilder{
		base: db.GetDatabase("goradd").NewBuilder(ctx),
	}
	return b.Join(node.Gift())
}

// Load terminates the query builder, performs the query, and returns a slice of Gift objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *GiftsBuilder) Load() (giftSlice []*Gift) {
	results := b.base.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Gift)
		o.load(item, o, nil, "")
		giftSlice = append(giftSlice, o)
	}
	return giftSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *GiftsBuilder) LoadI() (giftSlice []interface{}) {
	results := b.base.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Gift)
		o.load(item, o, nil, "")
		giftSlice = append(giftSlice, o)
	}
	return giftSlice
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *GiftsBuilder) Get() *Gift {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *GiftsBuilder) Expand(n query.NodeI) *GiftsBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *GiftsBuilder) Join(n query.NodeI, conditions ...query.NodeI) *GiftsBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *GiftsBuilder) Where(c query.NodeI) *GiftsBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *GiftsBuilder) OrderBy(nodes ...query.NodeI) *GiftsBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *GiftsBuilder) Limit(maxRowCount int, offset int) *GiftsBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *GiftsBuilder) Select(nodes ...query.NodeI) *GiftsBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *GiftsBuilder) Alias(name string, n query.NodeI) *GiftsBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *GiftsBuilder) Distinct() *GiftsBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *GiftsBuilder) GroupBy(nodes ...query.NodeI) *GiftsBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *GiftsBuilder) Having(node query.NodeI) *GiftsBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *GiftsBuilder) Count(distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *GiftsBuilder) Delete() {
	b.base.Delete()
	broadcast.BulkChange(b.base.Context(), "goradd", "gift")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *GiftsBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *GiftsBuilder) joinOrSelect(nodes ...query.NodeI) *GiftsBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountGiftByNumber(ctx context.Context, number int) uint {
	return queryGifts(ctx).Where(Equal(node.Gift().Number(), number)).Count(false)
}

func CountGiftByName(ctx context.Context, name string) uint {
	return queryGifts(ctx).Where(Equal(node.Gift().Name(), name)).Count(false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *giftBase) load(m map[string]interface{}, objThis *Gift, objParent interface{}, parentKey string) {
	if v, ok := m["number"]; ok && v != nil {
		if o.number, ok = v.(int); ok {
			o.numberIsValid = true
			o.numberIsDirty = false
			o._originalPK = o.number
		} else {
			panic("Wrong type found for number.")
		}
	} else {
		o.numberIsValid = false
		o.number = 0
	}

	if v, ok := m["name"]; ok && v != nil {
		if o.name, ok = v.(string); ok {
			o.nameIsValid = true
			o.nameIsDirty = false
		} else {
			panic("Wrong type found for name.")
		}
	} else {
		o.nameIsValid = false
		o.name = ""
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *giftBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *giftBase) update(ctx context.Context) {
	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if !o._restored {
			panic("Cannot update a record that was not originally read from the database.")
		}

		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "gift", modifiedFields, "number", o._originalPK)
		}

	}) // transaction
	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "gift", o._originalPK, stringmap.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the item into the database. Related items will be saved.
func (o *giftBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if !o.numberIsValid {
			panic("a value for Number is required, and there is no default value. Call SetNumber() before inserting the record.")
		}

		if !o.nameIsValid {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}
		m := o.getValidFields()

		d.Insert(ctx, "gift", m)
		id := o.PrimaryKey()
		o._originalPK = id

	}) // transaction
	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "gift", o.PrimaryKey())
}

func (o *giftBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.numberIsDirty {
		fields["number"] = o.number
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	return
}

func (o *giftBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.numberIsValid {
		fields["number"] = o.number
	}
	if o.nameIsValid {
		fields["name"] = o.name
	}
	return
}

// Delete deletes the associated record from the database.
func (o *giftBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "gift", "number", o.number)
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(o.number))
}

// deleteGift deletes the associated record from the database.
func deleteGift(ctx context.Context, pk int) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "gift", "number", pk)
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(pk))
}

func (o *giftBase) resetDirtyStatus() {
	o.numberIsDirty = false
	o.nameIsDirty = false

}

func (o *giftBase) IsDirty() bool {
	return o.numberIsDirty ||
		o.nameIsDirty
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *giftBase) Get(key string) interface{} {

	switch key {
	case "Number":
		if !o.numberIsValid {
			return nil
		}
		return o.number

	case "Name":
		if !o.nameIsValid {
			return nil
		}
		return o.name

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *giftBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.number); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.numberIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.numberIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.name); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, err
	}

	if o._aliases == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o._aliases); err != nil {
			return nil, err
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *giftBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr

	if err = dec.Decode(&o.number); err != nil {
		return
	}
	if err = dec.Decode(&o.numberIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.numberIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.name); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&isPtr); err != nil {
		return
	}
	if isPtr {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}
	if err = dec.Decode(&o._originalPK); err != nil {
		return
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *giftBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *giftBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.numberIsValid {
		v["number"] = o.number
	}

	if o.nameIsValid {
		v["name"] = o.name
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the gift. The gift can be a
// newly created object, or one loaded from the database.
//
// After unmarshalling, the object is not  saved. You must call Save to insert it into the database
// or update it.
//
// Unmarshalling of sub-objects, as in objects linked via foreign keys, is not currently supported.
//
// The fields it expects are:
//   "number" - int
//   "name" - string
func (o *giftBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in gift to modify the json before sending it here.
func (o *giftBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {
		case "number":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if n, ok := v.(float64); ok {
					o.SetNumber(int(n))
				} else {
					return fmt.Errorf("json field %s must be a number", k)
				}
			}
		case "name":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}
				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetName(s)
				}
			}

		}
	}
	return
}

// Custom functions. See goradd/codegen/templates/orm/modelBase.
