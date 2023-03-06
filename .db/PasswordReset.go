// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;

// class PasswordReset extends Model
//
//	{
//	    /**
//	     * @var string[]
//	     */
//	    protected $fillable = [
//	        'email', 'token'
//	    ];
//	}
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/mixin"
)

// PasswordReset holds the schema definition for the PasswordReset entity.
type PasswordReset struct {
	ent.Schema
}

// Fields of the PasswordReset.
func (PasswordReset) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("token"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
	}
}

// Mixin of the PasswordReset.
func (PasswordReset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
