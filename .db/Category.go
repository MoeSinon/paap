// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\Relations\HasManyThrough;
// use Illuminate\Database\Eloquent\Relations\HasOne;

// class Category extends Model
// {
//     use Sluggable;

//     protected $table = 'categories';

//     public $guarded = [];

//     protected $casts = [
//         'image' => 'json',
//         'banner_image' => 'json',
//     ];

//     protected $appends = array('parent_id');

//     /**
//      * Get the user's full name.
//      *
//      * @return string
//      */
//     public function getParentIdAttribute()
//     {
//         return $this->parent;
//     }


//     // protected static function boot()
//     // {
//     //     parent::boot();
//     //     // Order by updated_at desc
//     //     static::addGlobalScope('order', function (Builder $builder) {
//     //         $builder->orderBy('updated_at', 'desc');
//     //     });
//     // }

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

//     /**
//      * @return BelongsToMany
//      */
//     public function products(): BelongsToMany
//     {
//         return $this->belongsToMany(Product::class, 'category_product');
//     }

//     /**
//      * @return HasMany
//      */
//     public function children()
//     {
//         return $this->hasMany('Marvel\Database\Models\Category', 'parent', 'id')->with('children')->withCount('products');
//     }

//     /**
//      * @return HasMany
//      */
//     public function subCategories()
//     {
//         return $this->hasMany('Marvel\Database\Models\Category', 'parent', 'id')->with('subCategories', 'parent')->withCount('products');
//     }

//     /**
//      * @return HasOne
//      */
//     public function parent()
//     {
//         return $this->hasOne('Marvel\Database\Models\Category', 'id', 'parent')->with('parent');
//     }
// }
package category

import (
    "context"
    "github.com/google/uuid"
    "time"
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Category holds the schema definition for the Category entity.
type Category struct {
    ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New),
        field.String("name").NotEmpty(),
        field.Text("description").NotEmpty(),
        field.JSON("image", []byte{}),
        field.JSON("banner_image", []byte{}),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("type", Type.Type),
        edge.From("parent", Category.Type).Unique(),
        edge.To("children", Category.Type),
        edge.From("products", Product.Type).RelationTable("category_product"),
    }
}

// GetParent returns the parent category of this category.
func (c *Category) GetParent(ctx context.Context) (*Category, error) {
    parent, err := c.QueryParent().Only(ctx)
    if err != nil {
        return nil, err
    }
    return parent, nil
}

// GetChildren returns the children categories of this category.
func (c *Category) GetChildren(ctx context.Context) ([]*Category, error) {
    children, err := c.QueryChildren().All(ctx)
    if err != nil {
        return nil, err
    }
    return children, nil
}

// GetSubCategories returns the sub-categories of this category, including the parent and children.
func (c *Category) GetSubCategories(ctx context.Context) ([]*Category, error) {
    subCategories, err := c.QuerySubCategories().WithParent().WithChildren().All(ctx)
    if err != nil {
        return nil, err
    }
    return subCategories, nil
}
