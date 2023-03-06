// <?php

// namespace Marvel\Database\Models;

// use App\Enums\RoleType;
// use Illuminate\Database\Eloquent\Relations\HasOne;
// use Spatie\Permission\Traits\HasRoles;
// use Illuminate\Database\Eloquent\Builder;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Foundation\Auth\User as Authenticatable;
// use Illuminate\Notifications\Notifiable;
// use Laravel\Sanctum\HasApiTokens;


// class User extends Authenticatable
// {
//     use Notifiable;
//     use HasRoles;
//     use HasApiTokens;


//     protected $guard_name = 'api';

//     /**
//      * The attributes that are mass assignable.
//      *
//      * @var array
//      */
//     protected $fillable = [
//         'name', 'email', 'password', 'is_active', 'shop_id'
//     ];

//     /**
//      * The attributes that should be hidden for arrays.
//      *
//      * @var array
//      */
//     protected $hidden = [
//         'password', 'remember_token',
//     ];

//     protected static function boot()
//     {
//         parent::boot();
//         // Order by updated_at desc
//         static::addGlobalScope('order', function (Builder $builder) {
//             $builder->orderBy('updated_at', 'desc');
//         });
//     }

//     /**
//      * The attributes that should be cast to native types.
//      *
//      * @var array
//      */
//     protected $casts = [
//         'email_verified_at' => 'datetime',
//     ];


//     /**
//      * @return HasMany
//      */
//     public function address(): HasMany
//     {
//         return $this->hasMany(Address::class, 'customer_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function orders(): HasMany
//     {
//         return $this->hasMany(Order::class, 'customer_id')->with(['products.variation_options', 'status']);
//     }

//     /**
//      * @return HasOne
//      */
//     public function profile(): HasOne
//     {
//         return $this->hasOne(Profile::class, 'customer_id');
//     }

//     /**
//      * @return HasOne
//      */
//     public function shops(): HasMany
//     {
//         return $this->hasMany(Shop::class, 'owner_id');
//     }

//     /**
//      * @return BelongsTo
//      */
//     public function managed_shop(): BelongsTo
//     {
//         return $this->belongsTo(Shop::class, 'shop_id');
//     }

//     /**
//      * @return HasMany
//      */
//     public function providers(): HasMany
//     {
//         return $this->hasMany(Provider::class, 'user_id', 'id');
//     }
// }
package schema

import (
    "time"
    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}),
        field.String("name"),
        field.String("email"),
        field.String("password"),
        field.Bool("is_active").
            Default(true),
        field.Time("created_at").
            Default(time.Now),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("shop", Shop.Type).
            Ref("owner").
            Unique().
            Field("shop_id"),
        edge.To("addresses", Address.Type).
            Field("customer"),
        edge.To("orders", Order.Type).
            Field("customer"),
        edge.To("profile", Profile.Type).
            Field("customer").
            Unique(),
        edge.To("providers", Provider.Type).
            Field("user"),
    }
}
