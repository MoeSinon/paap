// <?php

// namespace Marvel\Database\Models;

// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\SoftDeletes;
// use Illuminate\Support\Facades\Log;
// use Marvel\Http\Controllers\TagController;
// use Marvel\Traits\Excludable;
// use Illuminate\Support\Facades\DB;

// class Product extends Model
// {
//     use Sluggable, SoftDeletes, Excludable;

//     public $guarded = [];

//     protected $table = 'products';

//     protected $appends = array('sold');

//     protected $casts = [
//         'image'   => 'json',
//         'gallery' => 'json',
//         'video' => 'json',
//     ];

//     // protected static function boot()
//     // {
//     //     parent::boot();
//     //     // Order by updated_at desc
//     //     static::addGlobalScope('order', function (Builder $builder) {
//     //         $builder->orderBy('updated_at', 'desc');
//     //     });
//     // }

//     /**
//      * Get the index name for the model.
//      *
//      * @return string
//      */
//     public function searchableAs()
//     {
//         return 'products_index';
//     }

//     /**
//      * Get the indexable data array for the model.
//      *
//      * @return array
//      */
//     public function toSearchableArray()
//     {
//         $array = $this->toArray();
//         return $array;
//     }

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
//      * @return BelongsTo
//      */
//     public function shop(): BelongsTo
//     {
//         return $this->belongsTo(Shop::class, 'shop_id');
//     }

//     /**
//      * @return BelongsTo
//      */
//     public function shipping(): BelongsTo
//     {
//         return $this->belongsTo(Shipping::class, 'shipping_class_id');
//     }

//     /**
//      * @return BelongsToMany
//      */
//     public function categories(): BelongsToMany
//     {
//         return $this->belongsToMany(Category::class, 'category_product');
//     }
//     /**
//      * @return BelongsToMany
//      */
//     public function tags(): BelongsToMany
//     {
//         return $this->belongsToMany(Tag::class, 'product_tag');
//     }

//     /**
//      * @return HasMany
//      */
//     public function variation_options(): HasMany
//     {
//         return $this->hasMany(Variation::class, 'product_id');
//     }

//     /**
//      * @return belongsToMany
//      */
//     public function orders(): belongsToMany
//     {
//         return $this->belongsToMany(Order::class)->withTimestamps();
//     }

//     /**
//      * @return BelongsToMany
//      */
//     public function variations(): BelongsToMany
//     {
//         return $this->belongsToMany(AttributeValue::class, 'attribute_product');
//     }

//	    /**
//	     * @return int|mixed
//	     */
//	    public function getSoldAttribute()
//	    {
//	        return DB::table('order_product')
//	            ->join('orders', 'orders.id', '=' , 'order_product.order_id')
//	            ->where('order_product.product_id', '=', $this->id)
//	            ->where('orders.parent_id', '=', null)
//	            ->sum('order_quantity');
//	    }
//	}
package schema

import (
	"context"
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("image", []byte{}),
		field.JSON("gallery", []byte{}),
		field.JSON("video", []byte{}),
		field.Int("sold"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional(),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("type", Type.Type).
			Ref("products").
			Unique().
			Required(),
		edge.From("shop", Shop.Type).
			Ref("products").
			Unique().
			Required(),
		edge.From("shipping", Shipping.Type).
			Ref("products").
			Unique().
			Required(),
		edge.From("categories", Category.Type).
			Ref("products").
			Required().
			Unique().
			// Use the join table "category_product".
			Annotations(ent.Annotation{
				ent.AssociationTable("category_product"),
			}),
		edge.From("tags", Tag.Type).
			Ref("products").
			Required().
			Unique().
			// Use the join table "product_tag".
			Annotations(ent.Annotation{
				ent.AssociationTable("product_tag"),
			}),
		edge.To("variations", Variation.Type).
			Annotations(ent.Annotation{
				ent.AssociationTable("attribute_product"),
			}),
		edge.From("orders", Order.Type).
			Ref("products").
			Required().
			// Use the join table "order_product".
			Annotations(ent.Annotation{
				ent.AssociationTable("order_product"),
			}),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		// Index the "name" field.
		index.Fields("name"),
		// Add a full-text search index on all fields.
		index.Fields("name", "image", "gallery", "video").
			Type("FULLTEXT"),
	}
}

func (Product) Hooks() []ent.Hook {
	return []ent.Hook{
		// Set the "updated_at" field to the current time on update.
		hook.OnUpdate(func(s ent.Schema, rd ent.Responder) {
			rd.(ent.Mutator).Set("updated_at", time.Now())
		}),
		// Soft-delete the entity by setting the "deleted_at" field to the current time on delete.
		hook.OnDelete(func(s ent.Schema, rd ent.Responder) {
			rd.(ent.Mutator).Set("deleted_at", time.Now())
		}),
	}
}

// Return a new query for all products that have been sold.
func (Product) Sold(ctx context.Context) *ent.ProductQuery {
	return client.Product.
		Query().
		Where(func(s *ent.Selector) {
			s.JoinOrderProduct().
				JoinOrders().
				Where(order_product.HasProductWith(product.IDEQ(s.ID()))).
				Where(orders.ParentIDIsNil()).
				GroupBy(order_product.FieldProductID).
				Select(order_product.FieldOrderQuantity.Sum())
		})
}
