// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/user"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/vote"
)

// Vote is the model entity for the Vote schema.
type Vote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// The user's selection
	Selection string `json:"selection,omitempty"`
	// The selection's position
	Rank int `json:"rank,omitempty"`
	// The vote's ID
	VoteID string `json:"vote_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VoteQuery when eager-loading is set.
	Edges        VoteEdges `json:"edges"`
	user_votes   *string
	selectValues sql.SelectValues
}

// VoteEdges holds the relations/edges for other nodes in the graph.
type VoteEdges struct {
	// User who voted
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VoteEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vote.FieldID, vote.FieldRank:
			values[i] = new(sql.NullInt64)
		case vote.FieldSelection, vote.FieldVoteID:
			values[i] = new(sql.NullString)
		case vote.ForeignKeys[0]: // user_votes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vote fields.
func (v *Vote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vote.FieldSelection:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field selection", values[i])
			} else if value.Valid {
				v.Selection = value.String
			}
		case vote.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				v.Rank = int(value.Int64)
			}
		case vote.FieldVoteID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vote_id", values[i])
			} else if value.Valid {
				v.VoteID = value.String
			}
		case vote.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_votes", values[i])
			} else if value.Valid {
				v.user_votes = new(string)
				*v.user_votes = value.String
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vote.
// This includes values selected through modifiers, order, etc.
func (v *Vote) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Vote entity.
func (v *Vote) QueryUser() *UserQuery {
	return NewVoteClient(v.config).QueryUser(v)
}

// Update returns a builder for updating this Vote.
// Note that you need to call Vote.Unwrap() before calling this method if this Vote
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vote) Update() *VoteUpdateOne {
	return NewVoteClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vote) Unwrap() *Vote {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vote is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vote) String() string {
	var builder strings.Builder
	builder.WriteString("Vote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("selection=")
	builder.WriteString(v.Selection)
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", v.Rank))
	builder.WriteString(", ")
	builder.WriteString("vote_id=")
	builder.WriteString(v.VoteID)
	builder.WriteByte(')')
	return builder.String()
}

// Votes is a parsable slice of Vote.
type Votes []*Vote
