// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Illuminate\Database\Eloquent\SoftDeletes;

// class Order extends Model
// {
//     use SoftDeletes;

//     protected $table = 'orders';

//     public $guarded = [];

//     protected $casts = [
//         'shipping_address' => 'json',
//         'billing_address'  => 'json',
//     ];

//     protected static function boot()
//     {
//         parent::boot();
//         // Order by created_at desc
//         static::addGlobalScope('order', function (Builder $builder) {
//             $builder->orderBy('created_at', 'desc');
//         });
//     }

//     protected $with = ['customer', 'status', 'products.variation_options'];

//     /**
//      * @return belongsToMany
//      */
//     public function products(): belongsToMany
//     {
//         return $this->belongsToMany(Product::class)
//             ->withPivot('order_quantity', 'unit_price', 'subtotal', 'variation_option_id')
//             ->withTimestamps();
//     }

//     /**
//      * @return belongsTo
//      */
//     public function status(): belongsTo
//     {
//         return $this->belongsTo(OrderStatus::class, 'status');
//     }

//     /**
//      * @return belongsTo
//      */
//     public function coupon(): belongsTo
//     {
//         return $this->belongsTo(Coupon::class, 'coupon_id');
//     }

//     /**
//      * @return belongsTo
//      */
//     public function customer(): belongsTo
//     {
//         return $this->belongsTo(User::class, 'customer_id');
//     }

//     /**
//      * @return BelongsTo
//      */
//     public function shop(): BelongsTo
//     {
//         return $this->belongsTo(Shop::class, 'shop_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function children()
//     {
//         return $this->hasMany('Marvel\Database\Models\Order', 'parent_id', 'id');
//     }

//     /**
//      * @return HasOne
//      */
//     public function parent_order()
//     {
//         return $this->hasOne('Marvel\Database\Models\Order', 'id', 'parent_id');
//     }
// }
package order

import (
    "time"

    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Order holds the schema definition for the Order entity.
type Order struct {
    ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
    return []ent.Field{
        field.Time("created_at").
            Default(time.Now),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
        field.Time("deleted_at").
            Optional(),
        field.JSON("shipping_address", &Address{}),
        field.JSON("billing_address", &Address{}),
    }
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("products", Product.Type).
            Annotations(ent.Annotation{
                entql.Annotation: `@withPivot("order_quantity", "unit_price", "subtotal", "variation_option_id")`,
            }),
        edge.From("status", OrderStatus.Type).
            Ref("orders").
            Field("status").
            Unique(),
        edge.From("coupon", Coupon.Type).
            Ref("orders").
            Field("coupon_id").
            Unique(),
        edge.From("customer", User.Type).
            Ref("orders").
            Field("customer_id").
            Unique(),
        edge.From("shop", Shop.Type).
            Ref("orders").
            Field("shop_id").
            Unique(),
        edge.To("children", Order.Type).
            From("parent").
            Field("parent_id"),
        edge.From("parent", Order.Type).
            Ref("children").
            Unique(),
    }
}

// Address holds the schema definition for an address.
type Address struct {
    ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
    return []ent.Field{
        field.String("street"),
        field.String("city"),
        field.String("state"),
        field.String("zip"),
        field.String("country"),
    }
}
