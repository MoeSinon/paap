// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;

// class Provider extends Model
// {
// 	protected $table = 'providers';

// 	protected $fillable = ['provider', 'provider_user_id', 'user_id'];

//		protected $hidden = [
//			'created_at',
//			'updated_at',
//		];
//	}
package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/MoeSinon/paap/db"
	"github.com/google/uuid"
)

// Provider holds information about a third-party authentication provider for a user.
type Provider struct {
	ent.Schema
}

// Fields defines the fields of the Provider schema.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("provider").NotEmpty(),
		field.String("provider_user_id").NotEmpty(),
		field.UUID("user_id", uuid.UUID{}).NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges defines the relationships of the Provider schema.
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("providers").
			Unique().
			Required(),
	}
}

// Annotations defines any additional annotations for the Provider schema.
func (Provider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent.Table("providers"),
	}
}

// ProviderMixin is a mixin that adds the timestamps to the Provider schema.
type ProviderMixin struct {
	mixin.Schema
}

// Fields adds the timestamps fields to the Provider schema.
func (ProviderMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Providers is the Ent schema for the Provider entity.
type Providers struct {
	ent.Schema
}

// Mixin adds the ProviderMixin to the Providers schema.
func (Providers) Mixin() []ent.Mixin {
	return []ent.Mixin{
		&ProviderMixin{},
	}
}

// ProviderQuery is a query for the Provider entity.
type ProviderQuery struct {
	*ent.Query
}

// NewProviderQuery creates a new ProviderQuery.
func NewProviderQuery() *ProviderQuery {
	return &ProviderQuery{Query: db.Client().Provider.Query()}
}

// Create creates a new Provider in the database.
func (q *ProviderQuery) Create(ctx context.Context, provider *Provider) (*Provider, error) {
	result, err := q.Query.
		Create().
		SetProvider(provider.Provider).
		SetProviderUserID(provider.ProviderUserID).
		SetUserID(provider.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return result.(*Provider), nil
}

// Get retrieves a Provider from the database.
func (q *ProviderQuery) Get(ctx context.Context, id uuid.UUID) (*Provider, error) {
	result, err := q.Query.
		Where(provider.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete deletes a Provider from the database.
func (q *ProviderQuery) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.Query.
		Where(provider.ID(id)).
		Delete().
		Exec(ctx)
	return err
}
