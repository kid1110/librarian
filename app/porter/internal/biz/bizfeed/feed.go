package bizfeed

import (
	"context"

	"github.com/tuihub/librarian/model/modelfeed"

	"github.com/muzhou233/go-favicon"
)

type FeedUseCase struct {
	rss     RSSRepo
	favicon *favicon.Finder
}

type RSSRepo interface {
	Parse(string) (*modelfeed.Feed, error)
	Get(string) (string, error)
}

func NewFeed(rss RSSRepo) *FeedUseCase {
	return &FeedUseCase{
		rss,
		favicon.New(favicon.IgnoreManifest),
	}
}

type FeedDestination int

const (
	FeedDestinationUnspecified FeedDestination = iota
	FeedDestinationTelegram
)

func (f *FeedUseCase) GetFeed(ctx context.Context, url string) (*modelfeed.Feed, error) {
	data, err := f.rss.Get(url)
	if err != nil {
		return nil, err
	}
	feed, err := f.rss.Parse(data)
	if err != nil {
		return nil, err
	}
	if len(feed.Link) > 0 {
		if icons, err1 := f.favicon.Find(feed.Link); err1 == nil && len(icons) > 0 {
			for _, icon := range icons {
				if icon.Height > 0 && icon.Width > 0 {
					feed.Image = &modelfeed.Image{
						URL:   icons[0].URL,
						Title: "",
					}
					break
				}
			}
		}
	}
	return feed, nil
}
