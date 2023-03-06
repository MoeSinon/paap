// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\hasMany;

// class OrderStatus extends Model
// {

//     protected $table = 'order_status';

//     public $guarded = [];

//     protected static function boot()
//     {
//         parent::boot();
//         static::addGlobalScope('order', function (Builder $builder) {
//             $builder->orderBy('serial', 'asc');
//         });
//     }

//	    /**
//	     * @return hasMany
//	     */
//	    public function orders(): hasMany
//	    {
//	        return $this->hasMany(Order::class, 'status');
//	    }
//	}
package schema

import (
	"context"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// OrderStatus holds the schema definition for the OrderStatus entity.
type OrderStatus struct {
	ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("serial"),
	}
}

// Edges of the OrderStatus.
func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
	}
}

// QueryOrders returns the query for retrieving all orders with the given status.
func (o *OrderStatus) QueryOrders() ent.Query {
	return o.Query().QueryOrders()
}

// QueryOrders returns the query for retrieving all orders with the given status.
func (q *OrderStatusQuery) QueryOrders() ent.Query {
	return q.Query().QueryOrders()
}

// QueryOrders returns the query for retrieving all orders with the given status.
func (p *OrderStatusPredicate) QueryOrders() ent.Query {
	return p.Query().QueryOrders()
}

// QueryOrders returns the query for retrieving all orders with the given status.
func (p OrderStatusPredicate) QueryOrders() ent.Query {
	return p.Query().QueryOrders()
}

// Orders returns the orders associated with the order status.
func (o *OrderStatus) Orders(ctx context.Context) ([]*Order, error) {
	return o.QueryOrders().All(ctx)
}
