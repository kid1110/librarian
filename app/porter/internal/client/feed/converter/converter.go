package converter

import (
	"time"

	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/mmcdole/gofeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --packagePath github.com/tuihub/librarian/app/porter/internal/client/feed/converter --packageName converter --output ./generated.go ./

// goverter:converter
type Converter interface {
	// goverter:matchIgnoreCase
	// goverter:ignore InternalID
	ToPBFeed(t *gofeed.Feed) *modelfeed.Feed
	// goverter:matchIgnoreCase
	// goverter:ignore InternalID
	// goverter:map UpdatedParsed | TimeToTime
	// goverter:map PublishedParsed | TimeToTime
	// goverter:ignore PublishPlatform
	ToPBFeedItem(t *gofeed.Item) *modelfeed.Item
}

func NewConverter() Converter {
	return &ConverterImpl{}
}

func TimeToTime(t *time.Time) *time.Time {
	return t
}
