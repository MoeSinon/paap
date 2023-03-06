// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\HasMany;

// class Shipping extends Model
// {
//     protected $table = 'shipping_classes';
//     public $guarded = [];

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
//     public function products(): HasMany
//     {
//         return $this->hasMany(Product::class, 'shipping_class_id');
//     }
// }
package schema

import (
    "context"
    "time"

    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/mixin"
    "entgo.io/ent/dialect/sql"
)

// Shipping holds the schema definition for the Shipping entity.
type Shipping struct {
    ent.Schema
}

// Fields of the Shipping.
func (Shipping) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

// Edges of the Shipping.
func (Shipping) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("products", Product.Type).
            Annotations(FromSchema),
    }
}

// Mixin of the Shipping.
func (Shipping) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

// Annotations of the Shipping.
func (Shipping) Annotations() []schema.Annotation {
    return []schema.Annotation{
        schema.TableAnnotation{
            Name: "shipping_classes",
        },
    }
}

// Query scope to order by updated_at desc.
func OrderByUpdatedAtDesc() ent.Mutator {
    return func(ctx context.Context, q *ent.Query) (interface{}, error) {
        return q.Order(ent.Desc("updated_at")), nil
    }
}

// Product holds the schema definition for the Product entity.
type Product struct {
    ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("shipping", Shipping.Type).
            Ref("products").
            Field("shipping_class_id").
            Annotations(FromSchema),
    }
}

// Annotations of the Product.
func (Product) Annotations() []schema.Annotation {
    return []schema.Annotation{
        schema.TableAnnotation{
            Name: "products",
        },
    }
}

// FromSchema annotation to link the schema definition with the model.
func FromSchema() schema.Annotation {
    return schema.Annotation{
        Annotations: []schema.Annotation{
            schema.StructTagAnnotation{
                Name: "json",
                Args: []string:"-",
            },
        },
    }
}

// Initialize the database schema.
func InitSchema(client *ent.Client) error {
    return client.Schema.Create(
        context.Background(),
        sql.WithGlobalMutators(OrderByUpdatedAtDesc),
    )
}
