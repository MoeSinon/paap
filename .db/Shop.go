// <?php

// namespace Marvel\Database\Models;

// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\Relations\HasOne;

// class Shop extends Model
// {
//     use Sluggable;

//     protected $table = 'shops';

//     public $guarded = [];

//     protected $casts = [
//         'logo' => 'json',
//         'cover_image' => 'json',
//         'address' => 'json',
//         'settings' => 'json',
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
//     public function balance(): HasOne
//     {
//         return $this->hasOne(Balance::class, 'shop_id');
//     }
//     /**
//      * @return BelongsTo
//      */
//     public function owner(): BelongsTo
//     {
//         return $this->belongsTo(User::class, 'owner_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function orders(): HasMany
//     {
//         return $this->hasMany(Order::class, 'shop_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function attributes(): HasMany
//     {
//         return $this->hasMany(Attribute::class, 'shop_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function products(): HasMany
//     {
//         return $this->hasMany(Product::class, 'shop_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function withdraws(): HasMany
//     {
//         return $this->hasMany(Withdraw::class, 'shop_id');
//     }

//     /**
//      * @return BelongsToMany
//      */
//     public function staffs(): HasMany
//     {
//         return $this->hasMany(User::class, 'shop_id');
//     }

//	    /**
//	     * @return BelongsToMany
//	     */
//	    public function categories(): BelongsToMany
//	    {
//	        return $this->belongsToMany(Category::class, 'category_shop');
//	    }
//	}
package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Shop holds the schema definition for the Shop entity.
type Shop struct {
	ent.Schema
}

// Fields of the Shop.
func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.JSON("logo", []byte("{}")),
		field.JSON("cover_image", []byte("{}")),
		field.JSON("address", []byte("{}")),
		field.JSON("settings", []byte("{}")),
	}
}

// Edges of the Shop.
func (Shop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("balance", Balance.Type),
		edge.From("owner", User.Type).Ref("owned_shops").Unique(),
		edge.To("orders", Order.Type),
		edge.To("attributes", Attribute.Type),
		edge.To("products", Product.Type),
		edge.To("withdraws", Withdraw.Type),
		edge.From("staffs", User.Type).Ref("shops").Unique(),
		edge.From("categories", Category.Type).Ref("shops").Unique(),
	}
}
