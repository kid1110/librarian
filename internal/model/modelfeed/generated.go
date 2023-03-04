// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package modelfeed

import v1 "github.com/tuihub/protos/pkg/librarian/v1"

type ConverterImpl struct{}

func (c *ConverterImpl) FromPBFeed(source *v1.Feed) *Feed {
	var pModelfeedFeed *Feed
	if source != nil {
		var modelfeedFeed Feed
		modelfeedFeed.ID = FromPBInternalID((*source).Id)
		modelfeedFeed.Title = (*source).Title
		modelfeedFeed.Description = (*source).Description
		modelfeedFeed.Link = (*source).Link
		var pModelfeedPersonList []*Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pV1FeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedFeed.Authors = pModelfeedPersonList
		modelfeedFeed.Language = (*source).Language
		modelfeedFeed.Image = c.FromPBFeedImage((*source).Image)
		var pModelfeedItemList []*Item
		if (*source).Items != nil {
			pModelfeedItemList = make([]*Item, len((*source).Items))
			for j := 0; j < len((*source).Items); j++ {
				pModelfeedItemList[j] = c.FromPBFeedItem((*source).Items[j])
			}
		}
		modelfeedFeed.Items = pModelfeedItemList
		pModelfeedFeed = &modelfeedFeed
	}
	return pModelfeedFeed
}
func (c *ConverterImpl) FromPBFeedEnclosure(source *v1.FeedEnclosure) *Enclosure {
	var pModelfeedEnclosure *Enclosure
	if source != nil {
		var modelfeedEnclosure Enclosure
		modelfeedEnclosure.URL = (*source).Url
		modelfeedEnclosure.Length = (*source).Length
		modelfeedEnclosure.Type = (*source).Type
		pModelfeedEnclosure = &modelfeedEnclosure
	}
	return pModelfeedEnclosure
}
func (c *ConverterImpl) FromPBFeedImage(source *v1.FeedImage) *Image {
	var pModelfeedImage *Image
	if source != nil {
		var modelfeedImage Image
		modelfeedImage.URL = (*source).Url
		modelfeedImage.Title = (*source).Title
		pModelfeedImage = &modelfeedImage
	}
	return pModelfeedImage
}
func (c *ConverterImpl) FromPBFeedItem(source *v1.FeedItem) *Item {
	var pModelfeedItem *Item
	if source != nil {
		var modelfeedItem Item
		modelfeedItem.ID = FromPBInternalID((*source).Id)
		modelfeedItem.Title = (*source).Title
		modelfeedItem.Description = (*source).Description
		modelfeedItem.Content = (*source).Content
		modelfeedItem.Link = (*source).Link
		modelfeedItem.Updated = (*source).Updated
		modelfeedItem.UpdatedParsed = FromPBTime((*source).UpdatedParsed)
		modelfeedItem.Published = (*source).Published
		modelfeedItem.PublishedParsed = FromPBTime((*source).PublishedParsed)
		var pModelfeedPersonList []*Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pV1FeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedItem.Authors = pModelfeedPersonList
		modelfeedItem.GUID = (*source).Guid
		modelfeedItem.Image = c.FromPBFeedImage((*source).Image)
		var pModelfeedEnclosureList []*Enclosure
		if (*source).Enclosures != nil {
			pModelfeedEnclosureList = make([]*Enclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pModelfeedEnclosureList[j] = c.FromPBFeedEnclosure((*source).Enclosures[j])
			}
		}
		modelfeedItem.Enclosures = pModelfeedEnclosureList
		modelfeedItem.PublishPlatform = (*source).PublishPlatform
		pModelfeedItem = &modelfeedItem
	}
	return pModelfeedItem
}
func (c *ConverterImpl) ToPBFeed(source *Feed) *v1.Feed {
	var pV1Feed *v1.Feed
	if source != nil {
		var v1Feed v1.Feed
		v1Feed.Id = c.modelfeedFeedToPV1InternalID((*source))
		v1Feed.Title = (*source).Title
		v1Feed.Link = (*source).Link
		v1Feed.Description = (*source).Description
		var pV1FeedItemList []*v1.FeedItem
		if (*source).Items != nil {
			pV1FeedItemList = make([]*v1.FeedItem, len((*source).Items))
			for i := 0; i < len((*source).Items); i++ {
				pV1FeedItemList[i] = c.ToPBFeedItem((*source).Items[i])
			}
		}
		v1Feed.Items = pV1FeedItemList
		v1Feed.Language = (*source).Language
		v1Feed.Image = c.ToPBFeedImage((*source).Image)
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for j := 0; j < len((*source).Authors); j++ {
				pV1FeedPersonList[j] = c.pModelfeedPersonToPV1FeedPerson((*source).Authors[j])
			}
		}
		v1Feed.Authors = pV1FeedPersonList
		pV1Feed = &v1Feed
	}
	return pV1Feed
}
func (c *ConverterImpl) ToPBFeedEnclosure(source *Enclosure) *v1.FeedEnclosure {
	var pV1FeedEnclosure *v1.FeedEnclosure
	if source != nil {
		var v1FeedEnclosure v1.FeedEnclosure
		v1FeedEnclosure.Url = (*source).URL
		v1FeedEnclosure.Length = (*source).Length
		v1FeedEnclosure.Type = (*source).Type
		pV1FeedEnclosure = &v1FeedEnclosure
	}
	return pV1FeedEnclosure
}
func (c *ConverterImpl) ToPBFeedImage(source *Image) *v1.FeedImage {
	var pV1FeedImage *v1.FeedImage
	if source != nil {
		var v1FeedImage v1.FeedImage
		v1FeedImage.Url = (*source).URL
		v1FeedImage.Title = (*source).Title
		pV1FeedImage = &v1FeedImage
	}
	return pV1FeedImage
}
func (c *ConverterImpl) ToPBFeedInternalID(source Feed) v1.InternalID {
	var v1InternalID v1.InternalID
	v1InternalID.Id = int64(source.ID)
	return v1InternalID
}
func (c *ConverterImpl) ToPBFeedItem(source *Item) *v1.FeedItem {
	var pV1FeedItem *v1.FeedItem
	if source != nil {
		var v1FeedItem v1.FeedItem
		v1FeedItem.Title = (*source).Title
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pV1FeedPersonList[i] = c.pModelfeedPersonToPV1FeedPerson((*source).Authors[i])
			}
		}
		v1FeedItem.Authors = pV1FeedPersonList
		v1FeedItem.Description = (*source).Description
		v1FeedItem.Content = (*source).Content
		v1FeedItem.Guid = (*source).GUID
		v1FeedItem.Link = (*source).Link
		v1FeedItem.Image = c.ToPBFeedImage((*source).Image)
		v1FeedItem.Published = (*source).Published
		v1FeedItem.PublishedParsed = ToPBTime((*source).PublishedParsed)
		v1FeedItem.Updated = (*source).Updated
		v1FeedItem.UpdatedParsed = ToPBTime((*source).UpdatedParsed)
		var pV1FeedEnclosureList []*v1.FeedEnclosure
		if (*source).Enclosures != nil {
			pV1FeedEnclosureList = make([]*v1.FeedEnclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pV1FeedEnclosureList[j] = c.ToPBFeedEnclosure((*source).Enclosures[j])
			}
		}
		v1FeedItem.Enclosures = pV1FeedEnclosureList
		v1FeedItem.PublishPlatform = (*source).PublishPlatform
		pV1FeedItem = &v1FeedItem
	}
	return pV1FeedItem
}
func (c *ConverterImpl) modelfeedFeedToPV1InternalID(source Feed) *v1.InternalID {
	v1InternalID := c.ToPBFeedInternalID(source)
	return &v1InternalID
}
func (c *ConverterImpl) pModelfeedPersonToPV1FeedPerson(source *Person) *v1.FeedPerson {
	var pV1FeedPerson *v1.FeedPerson
	if source != nil {
		var v1FeedPerson v1.FeedPerson
		v1FeedPerson.Name = (*source).Name
		v1FeedPerson.Email = (*source).Email
		pV1FeedPerson = &v1FeedPerson
	}
	return pV1FeedPerson
}
func (c *ConverterImpl) pV1FeedPersonToPModelfeedPerson(source *v1.FeedPerson) *Person {
	var pModelfeedPerson *Person
	if source != nil {
		var modelfeedPerson Person
		modelfeedPerson.Name = (*source).Name
		modelfeedPerson.Email = (*source).Email
		pModelfeedPerson = &modelfeedPerson
	}
	return pModelfeedPerson
}
