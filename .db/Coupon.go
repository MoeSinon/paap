// <?php

// namespace Marvel\Database\Models;

// use Carbon\Carbon;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\SoftDeletes;

// class Coupon extends Model
// {
//     use SoftDeletes;

//     protected $table = 'coupons';

//     public $guarded = [];

//     protected $appends = ['is_valid'];

//     protected $casts = [
//         'image'   => 'json',
//     ];

//     protected static function boot()
//     {
//         parent::boot();
//         // Order by updated_at desc
//         static::addGlobalScope('order', function (Builder $builder) {
//             $builder->orderBy('updated_at', 'desc');
//         });
//     }

//     /**
//      * @return HasMany
//      */
//     public function orders(): HasMany
//     {
//         return $this->hasMany(Order::class, 'coupon_id');
//     }

//     /**
//      * @return bool
//      */
//     public function getIsValidAttribute()
//     {
//         return Carbon::now()->between($this->active_from, $this->expire_at);
//     }
// }

package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"entgo.io/ent/schema/time"
	"github.com/google/uuid"
)

// Coupon holds the schema definition for the Coupon entity.
type Coupon struct {
	ent.Schema
}

// Mixin Coupon with the TimeMixin, to add automatic timestamps
// for created_at, updated_at and deleted_at fields.
func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("code"),
		field.Float("discount_amount"),
		field.Time("active_from").
			Default(time.Now),
		field.Time("expire_at"),
		field.JSON("image", []byte{}),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
	}
}

// ValidAt returns true if the current time is between the coupon's
// active_from and expire_at fields.
func (c *Coupon) ValidAt(ctx context.Context) bool {
	now := time.Now()
	return now.After(c.ActiveFrom) && now.Before(c.ExpireAt)
}
