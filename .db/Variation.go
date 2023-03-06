// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;

// class Variation extends Model
// {
//     protected $table = 'variation_options';

//     public $guarded = [];

//     protected $casts = [
//         'options'   => 'json',
//     ];

//	    /**
//	     * @return BelongsTo
//	     */
//	    public function product(): BelongsTo
//	    {
//	        return $this->belongsTo(Product::class, 'product_id');
//	    }
//	}
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Variation holds the schema definition for the Variation entity.
type Variation struct {
	ent.Schema
}

// Fields of the Variation.
func (Variation) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("options", map[string]interface{}{}),
	}
}

// Edges of the Variation.
func (Variation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("variations").
			Field("product_id").
			Unique(),
	}
}
