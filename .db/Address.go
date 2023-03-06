package schema

import (
    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/field"
    "github.com/facebook/ent/schema/edge"
)

type Address struct {
    ent.Schema
}

func (Address) Fields() []ent.Field {
    return []ent.Field{
        field.JSON("address", []byte{}),
    }
}

func (Address) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("customer", User.Type).
            Ref("addresses").
            Field("customer_id").
            Unique(),
    }
}

func (Address) Table() string {
    return "address"
}
