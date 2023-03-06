// <?php

// namespace Marvel\Database\Models;

// use Illuminate\Database\Eloquent\Model;
// use Illuminate\Database\Eloquent\Relations\BelongsTo;
// use Illuminate\Database\Eloquent\Relations\BelongsToMany;
// use Cviebrock\EloquentSluggable\Sluggable;
// use Illuminate\Database\Eloquent\Relations\HasMany;
// use Illuminate\Database\Eloquent\Relations\HasOne;
// use Spatie\MediaLibrary\HasMedia;
// use Spatie\MediaLibrary\InteractsWithMedia;
// use Spatie\MediaLibrary\MediaCollections\Models\Media;

// class Attachment extends Model implements HasMedia
// {
//     use InteractsWithMedia;

//     protected $table = 'attachments';

//     public $guarded = [];

//     public function registerMediaConversions(Media $media = null): void
//     {
//         $this->addMediaConversion('thumbnail')
//             ->width(368)
//             ->height(232)
//             ->nonQueued();
//     }
// }

package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/facebook/ent/schema/edge"
	"github.com/spf13/cast"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("collection_name"),
		field.Int("collection_id"),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
	}
}

// Mixin of the Attachment.
func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		// Add the "Media" edge, which connects Attachments to Media.
		// Here, "Media" is the name of the "Media" entity in the schema.
		// The "Attachment" entity is connected to "Media" via a many-to-one relationship.
		// The "Media" entity is connected to "Attachment" via a one-to-many relationship.
		// The foreign key is stored in the "media_id" column of the "attachments" table.
		edge.To("Media", Media.Type).
			Unique().
			Field("media_id").
			Required(),
	}
}

// Convert a Laravel Media object to an ent Media object.
func (a *Attachment) mediaToEntMedia(ctx context.Context, m models.Media) (*media.Media, error) {
	mObj := media.New(a, m.ID, m.Name, m.FileName, m.MediaType, m.Disk, m.Size, m.Manipulations, m.CustomProperties, m.OrderColumn)

	if m.ModelType != "" {
		mObj.SetModelType(m.ModelType)
	}

	if m.ModelID != "" {
		mObj.SetModelID(cast.ToInt(m.ModelID))
	}

	// set the URL generator function for the media entity
	mObj.SetUrlGenerator(func(_ context.Context, mediaObj *media.Media) (string, error) {
		return mediaObj.GetUrl(), nil
	})

	// store the media
	_, err := a.Tx(ctx).Save(mObj)
	if err != nil {
		return nil, err
	}

	return mObj, nil
}

// Add a new media object for this attachment.
func (a *Attachment) AddMedia(ctx context.Context, m models.Media) (*media.Media, error) {
	mediaObj, err := a.mediaToEntMedia(ctx, m)
	if err != nil {
		return nil, err
	}

	// add the media to the attachment
	return mediaObj, a.Media(func(q *ent.Query) {
		q.Where(media.IDEQ(mediaObj.ID))
	}).Save(ctx)
}

// Get all the media objects for this attachment.
func (a *Attachment) GetMedia(ctx context.Context) ([]*media.Media, error) {
	query := a.Query().WithMedia()

	att, err := query.First(ctx)
	if err != nil {
		return nil, err
	}

	if att == nil {
		return nil, nil
	}

	return att.Edges.Media, nil
}
