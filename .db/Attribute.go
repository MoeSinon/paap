// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\HasMany;

// class Attribute extends Model
// {
//     use Sluggable;

//     protected $table = 'attributes';

//     public $guarded = [];

//     /**
//      * Return the sluggable configuration array for this model.
//      *
//      * @return array
//      */
//     public function sluggable(): array
//     {
//         return [
//             'slug' => [
//                 'source' => 'name'
//             ]
//         ];
//     }

//     /**
//      * @return HasMany
//      */
//     public function values(): HasMany
//     {
//         return $this->hasMany(AttributeValue::class, 'attribute_id');
//     }

//     /**
//      * @return BelongsToMany
//      */
//     public function shop(): BelongsTo
//     {
//         return $this->belongsTo(Shop::class, 'shop_id');
//     }
// }

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Mixin of the Attribute.
func (Attribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug"),
		field.Int("shop_id"),
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		ent.Edge{
			Name: "values",
			Type: ent.OneToMany,
			// references the AttributeValue entity
			Target: ent.WeakEntityType("AttributeValue"),
			// the foreign key for AttributeValue entity
			ForeignKey: "attribute_id",
		},
		ent.Edge{
			Name: "shop",
			Type: ent.ManyToOne,
			// references the Shop entity
			Target: Shop.Type,
			// the foreign key for Shop entity
			ForeignKey: "shop_id",
		},
	}
}
