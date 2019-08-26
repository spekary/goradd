package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"
	"github.com/goradd/goradd/web/examples/model/node"

	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"

	//"./node"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// loginBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type loginBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	personID        string
	personIDIsNull  bool
	personIDIsValid bool
	personIDIsDirty bool
	oPerson         *Person

	username        string
	usernameIsValid bool
	usernameIsDirty bool

	password        string
	passwordIsNull  bool
	passwordIsValid bool
	passwordIsDirty bool

	isEnabled        bool
	isEnabledIsValid bool
	isEnabledIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool
}

const (
	LoginIDDefault        = ""
	LoginPersonIDDefault  = ""
	LoginUsernameDefault  = ""
	LoginPasswordDefault  = ""
	LoginIsEnabledDefault = false
)

const (
	LoginID        = `ID`
	LoginPersonID  = `PersonID`
	LoginPerson    = `Person`
	LoginUsername  = `Username`
	LoginPassword  = `Password`
	LoginIsEnabled = `IsEnabled`
)

// Initialize or re-initialize a Login database object to default values.
func (o *loginBase) Initialize() {

	o.id = ""
	o.idIsValid = false
	o.idIsDirty = false

	o.personID = ""
	o.personIDIsNull = true
	o.personIDIsValid = true
	o.personIDIsDirty = true

	o.username = ""
	o.usernameIsValid = false
	o.usernameIsDirty = false

	o.password = ""
	o.passwordIsNull = true
	o.passwordIsValid = true
	o.passwordIsDirty = true

	o.isEnabled = false
	o.isEnabledIsValid = false
	o.isEnabledIsDirty = false

	o._restored = false
}

func (o *loginBase) PrimaryKey() string {
	return o.id
}

// ID returns the loaded value of ID.
func (o *loginBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *loginBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

func (o *loginBase) PersonID() string {
	if o._restored && !o.personIDIsValid {
		panic("personID was not selected in the last query and so is not valid")
	}
	return o.personID
}

// PersonIDIsValid returns true if the value was loaded from the database or has been set.
func (o *loginBase) PersonIDIsValid() bool {
	return o.personIDIsValid
}

// PersonIDIsNull returns true if the related database value is null.
func (o *loginBase) PersonIDIsNull() bool {
	return o.personIDIsNull
}

// Person returns the current value of the loaded Person, and nil if its not loaded.
func (o *loginBase) Person() *Person {
	return o.oPerson
}

// LoadPerson returns the related Person. If it is not already loaded,
// it will attempt to load it first.
func (o *loginBase) LoadPerson(ctx context.Context) *Person {
	if !o.personIDIsValid {
		return nil
	}

	if o.oPerson == nil {
		// Load and cache
		o.oPerson = LoadPerson(ctx, o.PersonID())
	}
	return o.oPerson
}

func (o *loginBase) SetPersonID(i interface{}) {
	o.personIDIsValid = true
	if i == nil {
		if !o.personIDIsNull {
			o.personIDIsNull = true
			o.personIDIsDirty = true
			o.personID = ""
			o.oPerson = nil
		}
	} else {
		v := i.(string)
		if o.personIDIsNull ||
			!o._restored ||
			o.personID != v {

			o.personIDIsNull = false
			o.personID = v
			o.personIDIsDirty = true
			o.oPerson = nil
		}
	}
}

func (o *loginBase) SetPerson(v *Person) {
	o.personIDIsValid = true
	if v == nil {
		if !o.personIDIsNull || !o._restored {
			o.personIDIsNull = true
			o.personIDIsDirty = true
			o.personID = ""
			o.oPerson = nil
		}
	} else {
		o.oPerson = v
		if o.personIDIsNull || !o._restored || o.personID != v.PrimaryKey() {
			o.personIDIsNull = false
			o.personID = v.PrimaryKey()
			o.personIDIsDirty = true
		}
	}
}

func (o *loginBase) Username() string {
	if o._restored && !o.usernameIsValid {
		panic("username was not selected in the last query and so is not valid")
	}
	return o.username
}

// UsernameIsValid returns true if the value was loaded from the database or has been set.
func (o *loginBase) UsernameIsValid() bool {
	return o.usernameIsValid
}

// SetUsername sets the value of Username in the object, to be saved later using the Save() function.
func (o *loginBase) SetUsername(v string) {
	o.usernameIsValid = true
	if o.username != v || !o._restored {
		o.username = v
		o.usernameIsDirty = true
	}

}

func (o *loginBase) Password() string {
	if o._restored && !o.passwordIsValid {
		panic("password was not selected in the last query and so is not valid")
	}
	return o.password
}

// PasswordIsValid returns true if the value was loaded from the database or has been set.
func (o *loginBase) PasswordIsValid() bool {
	return o.passwordIsValid
}

// PasswordIsNull returns true if the related database value is null.
func (o *loginBase) PasswordIsNull() bool {
	return o.passwordIsNull
}

func (o *loginBase) SetPassword(i interface{}) {
	o.passwordIsValid = true
	if i == nil {
		if !o.passwordIsNull {
			o.passwordIsNull = true
			o.passwordIsDirty = true
			o.password = ""
		}
	} else {
		v := i.(string)
		if o.passwordIsNull ||
			!o._restored ||
			o.password != v {

			o.passwordIsNull = false
			o.password = v
			o.passwordIsDirty = true
		}
	}
}

func (o *loginBase) IsEnabled() bool {
	if o._restored && !o.isEnabledIsValid {
		panic("isEnabled was not selected in the last query and so is not valid")
	}
	return o.isEnabled
}

// IsEnabledIsValid returns true if the value was loaded from the database or has been set.
func (o *loginBase) IsEnabledIsValid() bool {
	return o.isEnabledIsValid
}

// SetIsEnabled sets the value of IsEnabled in the object, to be saved later using the Save() function.
func (o *loginBase) SetIsEnabled(v bool) {
	o.isEnabledIsValid = true
	if o.isEnabled != v || !o._restored {
		o.isEnabled = v
		o.isEnabledIsDirty = true
	}

}

// GetAlias returns the alias for the given key.
func (o *loginBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// Load returns a Login from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func LoadLogin(ctx context.Context, primaryKey string, joinOrSelectNodes ...query.NodeI) *Login {
	return queryLogins(ctx).Where(Equal(node.Login().ID(), primaryKey)).joinOrSelect(joinOrSelectNodes...).Get(ctx)
}

// LoadLoginByUsername queries for a single Login object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryLogins() to start a query builder.
func LoadLoginByUsername(ctx context.Context, username string, joinOrSelectNodes ...query.NodeI) *Login {
	return queryLogins(ctx).
		Where(Equal(node.Login().Username(), username)).
		joinOrSelect(joinOrSelectNodes...).
		Get(ctx)
}

// LoadLoginByPersonID queries for a single Login object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryLogins() to start a query builder.
func LoadLoginByPersonID(ctx context.Context, person_id string, joinOrSelectNodes ...query.NodeI) *Login {
	return queryLogins(ctx).
		Where(Equal(node.Login().PersonID(), person_id)).
		joinOrSelect(joinOrSelectNodes...).
		Get(ctx)
}

// The LoginsBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type LoginsBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newLoginBuilder() *LoginsBuilder {
	b := &LoginsBuilder{
		base: db.GetDatabase("goradd").
			NewBuilder(),
	}
	return b.Join(node.Login())
}

// Load terminates the query builder, performs the query, and returns a slice of Login objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *LoginsBuilder) Load(ctx context.Context) (loginSlice []*Login) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Login)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		loginSlice = append(loginSlice, o)
	}
	return loginSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *LoginsBuilder) LoadI(ctx context.Context) (loginSlice []interface{}) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Login)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		loginSlice = append(loginSlice, o)
	}
	return loginSlice
}

// Get is a convenience method to return only the first item found in a query. It is equivalent to adding
// Limit(1,0) to the query, and then getting the first item from the returned slice.
// Limits with joins do not currently work, so don't try it if you have a join
// TODO: Change this to Load1 to be more descriptive and avoid confusion with other Getters
func (b *LoginsBuilder) Get(ctx context.Context) *Login {
	results := b.Limit(1, 0).Load(ctx)
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *LoginsBuilder) Expand(n query.NodeI) *LoginsBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *LoginsBuilder) Join(n query.NodeI, conditions ...query.NodeI) *LoginsBuilder {
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
func (b *LoginsBuilder) Where(c query.NodeI) *LoginsBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *LoginsBuilder) OrderBy(nodes ...query.NodeI) *LoginsBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *LoginsBuilder) Limit(maxRowCount int, offset int) *LoginsBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *LoginsBuilder) Select(nodes ...query.NodeI) *LoginsBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *LoginsBuilder) Alias(name string, n query.NodeI) *LoginsBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *LoginsBuilder) Distinct() *LoginsBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *LoginsBuilder) GroupBy(nodes ...query.NodeI) *LoginsBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *LoginsBuilder) Having(node query.NodeI) *LoginsBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
func (b *LoginsBuilder) Count(ctx context.Context, distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(ctx, distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *LoginsBuilder) Delete(ctx context.Context) {
	b.base.Delete(ctx)
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *LoginsBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect us a private helper function for the Load* functions
func (b *LoginsBuilder) joinOrSelect(nodes ...query.NodeI) *LoginsBuilder {
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

func CountLoginByID(ctx context.Context, id string) uint {
	return queryLogins(ctx).Where(Equal(node.Login().ID(), id)).Count(ctx, false)
}

func CountLoginByPersonID(ctx context.Context, personID string) uint {
	return queryLogins(ctx).Where(Equal(node.Login().PersonID(), personID)).Count(ctx, false)
}

func CountLoginByUsername(ctx context.Context, username string) uint {
	return queryLogins(ctx).Where(Equal(node.Login().Username(), username)).Count(ctx, false)
}

func CountLoginByPassword(ctx context.Context, password string) uint {
	return queryLogins(ctx).Where(Equal(node.Login().Password(), password)).Count(ctx, false)
}

func CountLoginByIsEnabled(ctx context.Context, isEnabled bool) uint {
	return queryLogins(ctx).Where(Equal(node.Login().IsEnabled(), isEnabled)).Count(ctx, false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// If linkParent is true we will have child relationships use a pointer back to the parent object. If false, it will create a separate object.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
// Also, if any joins are conditional, that might affect which child objects are included, so in this situation, linkParent should be false
func (o *loginBase) load(m map[string]interface{}, linkParent bool, objThis *Login, objParent interface{}, parentKey string) {
	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false
		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["person_id"]; ok {
		if v == nil {
			o.personID = ""
			o.personIDIsNull = true
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else if o.personID, ok = v.(string); ok {
			o.personIDIsNull = false
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for person_id.")
		}
	} else {
		o.personIDIsValid = false
		o.personIDIsNull = true
		o.personID = ""
	}
	if linkParent && parentKey == "Person" {
		o.oPerson = objParent.(*Person)
		o.personIDIsValid = true
		o.personIDIsDirty = false
	} else if v, ok := m["Person"]; ok {
		if oPerson, ok2 := v.(map[string]interface{}); ok2 {
			o.oPerson = new(Person)
			o.oPerson.load(oPerson, linkParent, o.oPerson, objThis, "Logins")
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for oPerson object.")
		}
	} else {
		o.oPerson = nil
	}

	if v, ok := m["username"]; ok && v != nil {
		if o.username, ok = v.(string); ok {
			o.usernameIsValid = true
			o.usernameIsDirty = false
		} else {
			panic("Wrong type found for username.")
		}
	} else {
		o.usernameIsValid = false
		o.username = ""
	}

	if v, ok := m["password"]; ok {
		if v == nil {
			o.password = ""
			o.passwordIsNull = true
			o.passwordIsValid = true
			o.passwordIsDirty = false
		} else if o.password, ok = v.(string); ok {
			o.passwordIsNull = false
			o.passwordIsValid = true
			o.passwordIsDirty = false
		} else {
			panic("Wrong type found for password.")
		}
	} else {
		o.passwordIsValid = false
		o.passwordIsNull = true
		o.password = ""
	}
	if v, ok := m["is_enabled"]; ok && v != nil {
		if o.isEnabled, ok = v.(bool); ok {
			o.isEnabledIsValid = true
			o.isEnabledIsDirty = false
		} else {
			panic("Wrong type found for is_enabled.")
		}
	} else {
		o.isEnabledIsValid = false
		o.isEnabled = false
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *loginBase) Save(ctx context.Context) {
	if o._restored {
		o.Update(ctx)
	} else {
		o.Insert(ctx)
	}
}

// Update will update the values in the database, saving any changed values.
func (o *loginBase) Update(ctx context.Context) {
	if !o._restored {
		panic("Cannot update a record that was not originally read from the database.")
	}
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	d.Update(ctx, "login", m, "id", fmt.Sprint(o.id))
	o.resetDirtyStatus()
}

// Insert forces the object to be inserted into the database. If the object was loaded from the database originally,
// this will create a duplicate in the database.
func (o *loginBase) Insert(ctx context.Context) {
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	id := d.Insert(ctx, "login", m)
	o.id = id
	o.resetDirtyStatus()
	o._restored = true
}

func (o *loginBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	if o.personIDIsDirty {
		if o.personIDIsNull {
			fields["person_id"] = nil
		} else {
			fields["person_id"] = o.personID
		}
	}

	if o.usernameIsDirty {
		fields["username"] = o.username
	}

	if o.passwordIsDirty {
		if o.passwordIsNull {
			fields["password"] = nil
		} else {
			fields["password"] = o.password
		}
	}

	if o.isEnabledIsDirty {
		fields["is_enabled"] = o.isEnabled
	}

	return
}

// Delete deletes the associated record from the database.
func (o *loginBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "login", "id", o.id)
}

// deleteLogin deletes the associated record from the database.
func deleteLogin(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "login", "id", pk)
}

func (o *loginBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.personIDIsDirty = false
	o.usernameIsDirty = false
	o.passwordIsDirty = false
	o.isEnabledIsDirty = false
}

func (o *loginBase) IsDirty() bool {
	return o.idIsDirty ||
		o.personIDIsDirty ||
		o.usernameIsDirty ||
		o.passwordIsDirty ||
		o.isEnabledIsDirty
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *loginBase) Get(key string) interface{} {

	switch key {
	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "PersonID":
		if !o.personIDIsValid {
			return nil
		}
		return o.personID

	case "Person":
		return o.Person()

	case "Username":
		if !o.usernameIsValid {
			return nil
		}
		return o.username

	case "Password":
		if !o.passwordIsValid {
			return nil
		}
		return o.password

	case "IsEnabled":
		if !o.isEnabledIsValid {
			return nil
		}
		return o.isEnabled

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database object over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *loginBase) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err = encoder.Encode(o.id); err != nil {
		return
	}
	if err = encoder.Encode(o.idIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.idIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.personID); err != nil {
		return
	}
	if err = encoder.Encode(o.personIDIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.personIDIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.personIDIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.oPerson); err != nil {
		return
	}
	if err = encoder.Encode(o.username); err != nil {
		return
	}
	if err = encoder.Encode(o.usernameIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.usernameIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.password); err != nil {
		return
	}
	if err = encoder.Encode(o.passwordIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.passwordIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.passwordIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.isEnabled); err != nil {
		return
	}
	if err = encoder.Encode(o.isEnabledIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.isEnabledIsDirty); err != nil {
		return
	}

	if o._aliases == nil {
		if err = encoder.Encode(false); err != nil {
			return
		}
	} else {
		if err = encoder.Encode(true); err != nil {
			return
		}
		if err = encoder.Encode(o._aliases); err != nil {
			return
		}
	}

	if err = encoder.Encode(o._restored); err != nil {
		return
	}

	return
}

func (o *loginBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&o.id); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.personID); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.oPerson); err != nil {
		return
	}
	if err = dec.Decode(&o.username); err != nil {
		return
	}
	if err = dec.Decode(&o.usernameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.usernameIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.password); err != nil {
		return
	}
	if err = dec.Decode(&o.passwordIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.passwordIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.passwordIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.isEnabled); err != nil {
		return
	}
	if err = dec.Decode(&o.isEnabledIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.isEnabledIsDirty); err != nil {
		return
	}

	var hasAliases bool
	if err = dec.Decode(&hasAliases); err != nil {
		return
	}
	if hasAliases {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}

	return err
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object.
func (o *loginBase) MarshalJSON() (data []byte, err error) {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.personIDIsValid {
		if o.personIDIsNull {
			v["personID"] = nil
		} else {
			v["personID"] = o.personID
		}
	}

	if val := o.Person(); val != nil {
		v["person"] = val
	}
	if o.usernameIsValid {
		v["username"] = o.username
	}

	if o.passwordIsValid {
		if o.passwordIsNull {
			v["password"] = nil
		} else {
			v["password"] = o.password
		}
	}

	if o.isEnabledIsValid {
		v["isEnabled"] = o.isEnabled
	}

	return json.Marshal(v)
}
