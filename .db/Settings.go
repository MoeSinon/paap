// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;

// class Settings extends Model
// {
//     protected $table = 'settings';

//     public $guarded = [];

//	    protected $casts = [
//	        'options'   => 'json',
//	    ];
//	}
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("options", map[string]interface{}{}),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}

// Mixin of the Settings.
func (Settings) Mixin() []ent.Mixin {
	return nil
}
