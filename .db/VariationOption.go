// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;

// class VariationOption extends Model
// {
//     protected $table = 'variation_options';

//	    public $guarded = [];
//	}
package schema

import "entgo.io/ent"

// VariationOption holds the schema definition for the VariationOption entity.
type VariationOption struct {
	ent.Schema
}

// Fields of the VariationOption.
func (VariationOption) Fields() []ent.Field {
	return nil // You would define the fields here.
}

// Edges of the VariationOption.
func (VariationOption) Edges() []ent.Edge {
	return nil // You would define the edges here.
}
