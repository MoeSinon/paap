// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\SoftDeletes;

// class Withdraw extends Model
// {
//     use SoftDeletes;
//     protected $table = 'withdraws';

//     public $guarded = [];

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

// Withdraw holds the schema definition for the Withdraw entity.
type Withdraw struct {
	ent.Schema
}

// Fields of the Withdraw.
func (Withdraw) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
		field.Float("amount"),
	}
}

// Edges of the Withdraw.
func (Withdraw) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", Shop.Type).
			Ref("withdraws").
			Field("shop_id").
			Unique().
			Required(),
	}
}
