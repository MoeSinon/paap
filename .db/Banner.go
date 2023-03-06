// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;

// class Banner extends Model
// {
//     protected $table = 'banners';

//     public $guarded = [];

//     protected $casts = [
//         'image'   => 'json',
//     ];



//     /**
//      * @return BelongsTo
//      */
//     public function type(): BelongsTo
//     {
//         return $this->belongsTo(Type::class, 'type_id');
//     }
// }
package ent

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Banner holds the schema definition for the Banner entity.
type Banner struct {
	ent.Schema
}

// Fields of the Banner.
func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("image", []string{}),
	}
}

// Edges of the Banner.
func (Banner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("type", Type.Type).
			Ref("banners").
			Unique(),
	}
}
