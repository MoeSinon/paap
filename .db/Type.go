// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\hasMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Builder;

// class Type extends Model
// {

//     use Sluggable;

//     protected $table = 'types';

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
//                 'source' => 'name',
//             ]
//         ];
//     }

//     protected $casts = [
//         'promotional_sliders'   => 'json',
//         'images' => 'json',
//         'settings'   => 'json',
//     ];

//     /**
//      * @return HasMany
//      */
//     public function products(): HasMany
//     {
//         return $this->hasMany(Product::class, 'type_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function categories(): HasMany
//     {
//         return $this->hasMany(Category::class, 'type_id');
//     }

//	    /**
//	     * @return HasMany
//	     */
//	    public function banners(): HasMany
//	    {
//	        return $this->hasMany(Banner::class, 'type_id');
//	    }
//	}
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
)

// Type holds the schema definition for the Type entity.
type Type struct {
	ent.Schema
}

// Fields of the Type.
func (Type) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug").Unique(),
		field.JSON("promotional_sliders", []string{}),
		field.JSON("images", []string{}),
		field.JSON("settings", map[string]interface{}{}),
	}
}

// Mixin of the Type.
func (Type) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Type.
func (Type) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("categories", Category.Type),
		edge.To("banners", Banner.Type),
	}
}

// Indexes of the Type.
func (Type) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}
