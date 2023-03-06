// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\Relations\HasOne;

// class Tag extends Model
// {
//     use Sluggable;

//     protected $table = 'tags';

//     public $guarded = [];

//     protected $casts = [
//         'image' => 'json',
//     ];

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
//      * @return BelongsTo
//      */
//     public function type(): BelongsTo
//     {
//         return $this->belongsTo(Type::class, 'type_id');
//     }

//	    /**
//	     * @return BelongsToMany
//	     */
//	    public function products(): BelongsToMany
//	    {
//	        return $this->belongsToMany(Product::class, 'product_tag');
//	    }
//	}
package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("image", []byte{}),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("type", Type.Type),
		edge.From("products", Product.Type).
			Ref("tags").
			Unique(),
	}
}
