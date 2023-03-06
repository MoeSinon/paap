// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;

// class Profile extends Model
// {
//     protected $table = 'user_profiles';

//     public $guarded = [];

//     protected $casts = [
//         'socials' => 'json',
//         'avatar' => 'json',
//     ];

//	    /**
//	     * @return BelongsTo
//	     */
//	    public function customer(): BelongsTo
//	    {
//	        return $this->belongsTo(User::class, 'customer_id');
//	    }
//	}
package schema

import (
	"database/sql"
	"encoding/json"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("socials").
			NotEmpty().
			Default("{}").
			Comment("Social links in JSON format"),
		field.String("avatar").
			NotEmpty().
			Default("{}").
			Comment("Avatar information in JSON format"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", User.Type).
			Ref("profile").
			Field("customer_id").
			Unique(),
	}
}

// Mixin defines the timestamp mixin for Profile.
func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// SetSocials sets the social links for a profile.
func (p *Profile) SetSocials(links map[string]string) error {
	b, err := json.Marshal(links)
	if err != nil {
		return err
	}
	p.Socials = string(b)
	return nil
}

// GetSocials gets the social links for a profile.
func (p *Profile) GetSocials() (map[string]string, error) {
	var links map[string]string
	err := json.Unmarshal([]byte(p.Socials), &links)
	if err != nil {
		return nil, err
	}
	return links, nil
}

// SetAvatar sets the avatar information for a profile.
func (p *Profile) SetAvatar(info map[string]interface{}) error {
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}
	p.Avatar = string(b)
	return nil
}

// GetAvatar gets the avatar information for a profile.
func (p *Profile) GetAvatar() (map[string]interface{}, error) {
	var info map[string]interface{}
	err := json.Unmarshal([]byte(p.Avatar), &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetUser returns the user associated with this profile.
func (p *Profile) GetUser() (*User, error) {
	customer, err := p.QueryCustomer().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return customer, nil
}
