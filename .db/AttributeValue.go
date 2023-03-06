// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;

// class AttributeValue extends Model
// {
//     protected $table = 'attribute_values';

//     public $guarded = [];

//     /**
//      * @return BelongsTo
//      */
//     public function attribute(): BelongsTo
//     {
//         return $this->belongsTo(Attribute::class, 'attribute_id');
//     }

//     /**
//      * @return BelongsToMany
//      */
//     public function products(): BelongsToMany
//     {
//         return $this->belongsToMany(Product::class, 'attribute_product');
//     }
// }

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OrderStatus struct {
	ent.Schema
}

func (OrderStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("serial"),
	}
}

func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		ent.Edge{
			Name:    "orders",
			Type:    ent.RelationOneToMany,
			Target:  (*Order)(nil),
			Reverse: true,
		},
	}
}
