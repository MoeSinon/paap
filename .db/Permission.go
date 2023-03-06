// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Spatie\Permission\Models\Permission as SpatiePermission;

// class Permission extends SpatiePermission
//
//	{
//	    /**
//	     * A permission belongs to some users of the model associated with its guard.
//	     */
//	    public function users(): BelongsToMany
//	    {
//	        return $this->morphedByMany(
//	            getModelForGuard($this->attributes['guard_name']),
//	            'model',
//	            config('permission.table_names.model_has_permissions'),
//	            'permission_id',
//	            config('permission.column_names.model_morph_key')
//	        )->with('profile', 'address', 'shop');
//	    }
//	}
package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/beego/beego/v2/core/logs"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Immutable(),
		field.String("guard_name").Default("").Immutable(),
		field.Time("created_at").Immutable().Default(now),
		field.Time("updated_at").Default(now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("model_permissions", ModelPermission.Type).StorageKey(edge.Column("permission_id")),
	}
}

// Users returns the users associated with the permission.
func (Permission) Users() ent.Association {
	return permission.Users()
}

// WithProfile adds the profile information to the permission query.
func (Permission) WithProfile() ent.QueryModifier {
	return func(q *ent.Selector) {
		q.With(ProfileTable)
	}
}

// WithAddress adds the address information to the permission query.
func (Permission) WithAddress() ent.QueryModifier {
	return func(q *ent.Selector) {
		q.With(AddressTable)
	}
}

// WithShop adds the shop information to the permission query.
func (Permission) WithShop() ent.QueryModifier {
	return func(q *ent.Selector) {
		q.With(ShopTable)
	}
}

// New creates a new Permission entity.
func (Permission) New() ent.Entity {
	return &Permission{}
}

// UpdateTime implements the ent.Updater interface.
func (Permission) UpdateTime(_ context.Context, _ *ent.UpdateConfig) (ent.Value, error) {
	return now(), nil
}

// LogFields returns the fields to be logged for the Permission.
func (Permission) LogFields() []string {
	return []string{"id", "name"}
}

// OnCreate hooks into the creation of a Permission.
func (Permission) OnCreate(ctx context.Context, entObj *ent.Entity) {
	logs.Info("Permission created with ID: %d", entObj.ID)
}

// OnUpdate hooks into the update of a Permission.
func (Permission) OnUpdate(ctx context.Context, entObj *ent.Entity) {
	logs.Info("Permission updated with ID: %d", entObj.ID)
}

// OnDelete hooks into the deletion of a Permission.
func (Permission) OnDelete(ctx context.Context, entObj *ent.Entity) {
	logs.Info("Permission deleted with ID: %d", entObj.ID)
}
