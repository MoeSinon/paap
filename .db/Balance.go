// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;

// class Balance extends Model
// {
//     protected $table = 'balances';

//     public $guarded = [];

//     protected $casts = [
//         'payment_info' => 'json',
//     ];

//	    /**
//	     * @return BelongsTo
//	     */
//	    public function shop(): BelongsTo
//	    {
//	        return $this->belongsTo(Shop::class, 'shop_id');
//	    }
//	}
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Balance holds the schema definition for the Balance entity.
type Balance struct {
	ent.Schema
}

// Fields of the Balance.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("shop_id"),
		field.Float("amount"),
		field.JSON("payment_info", map[string]interface{}{}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Balance.
func (Balance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", Shop.Type).
			Ref("balances").
			Field("shop_id").
			Unique(),
	}
}
