// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\Relations\HasOne;

// class Tax extends Model
// {
//     protected $table = 'tax_classes';

//     public $guarded = [];

//	    protected static function boot()
//	    {
//	        parent::boot();
//	        // Order by updated_at desc
//	        static::addGlobalScope('order', function (Builder $builder) {
//	            $builder->orderBy('updated_at', 'desc');
//	        });
//	    }
//	}
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Tax holds the schema definition for the Tax entity.
type Tax struct {
	ent.Schema
}

// Fields of the Tax.
func (Tax) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("rate"),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Tax.
func (Tax) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("store", Store.Type).
			Ref("taxes").
			Unique(),
	}
}

// Mixin of the Tax.
func (Tax) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
