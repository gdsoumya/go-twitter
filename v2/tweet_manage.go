package twitter

import "fmt"

// CreateTweetRequest is the details of a tweet to create
type CreateTweetRequest struct {
	DirectMessageDeepLink string           `json:"direct_message_deep_link,omitempty"`
	ForSuperFollowersOnly bool             `json:"for_super_followers_only,omitempty"`
	QuoteTweetID          string           `json:"quote_tweet_id,omitempty"`
	Text                  string           `json:"text,omitempty"`
	ReplySettings         string           `json:"reply_settings"`
	Geo                   CreateTweetGeo   `json:"geo"`
	Media                 CreateTweetMedia `json:"media"`
	Poll                  CreateTweetPoll  `json:"poll"`
	Reply                 CreateTweetReply `json:"reply"`
}

func (t CreateTweetRequest) validate() error {
	if err := t.Media.validate(); err != nil {
		return fmt.Errorf("create tweet error: %w", err)
	}
	if err := t.Poll.validate(); err != nil {
		return fmt.Errorf("create tweet error: %w", err)
	}
	if err := t.Reply.validate(); err != nil {
		return fmt.Errorf("create tweet error: %w", err)
	}
	if len(t.Media.IDs) == 0 && len(t.Text) == 0 {
		return fmt.Errorf("create tweet text is required if no media ids %w", ErrParameter)
	}
	return nil
}

// CreateTweetGeo allows for the tweet to coontain geo
type CreateTweetGeo struct {
	PlaceID string `json:"place_id"`
}

// CreateTweetMedia allows for updated media to attached.
// If the tagged user ids are present, then ids must be present.
type CreateTweetMedia struct {
	IDs           []string `json:"media_ids"`
	TaggedUserIDs []string `json:"tagged_user_ids"`
}

func (m CreateTweetMedia) validate() error {
	if len(m.TaggedUserIDs) > 0 && len(m.IDs) == 0 {
		return fmt.Errorf("media ids are required if tagged user ids are present %w", ErrParameter)
	}
	return nil
}

// CreateTweetPoll allows for a poll to be posted as the tweet
type CreateTweetPoll struct {
	DurationMinutes int      `json:"duration_minutes"`
	Options         []string `json:"options"`
}

func (p CreateTweetPoll) validate() error {
	if len(p.Options) > 0 && p.DurationMinutes <= 0 {
		return fmt.Errorf("poll duration minutes are required with options %w", ErrParameter)
	}
	return nil
}

// CreateTweetReply sets the reply setting for the tweet
type CreateTweetReply struct {
	ExcludeReplyUserIDs []string `json:"exclude_reply_user_ids"`
	InReplyToTweetID    string   `json:"in_reply_to_tweet_id"`
}

func (r CreateTweetReply) validate() error {
	if len(r.ExcludeReplyUserIDs) > 0 && len(r.InReplyToTweetID) == 0 {
		return fmt.Errorf("in reply to tweet is needs to be present if excluded reply user ids are present %w", ErrParameter)
	}
	return nil
}

// CreateTweetData is the data returned when creating a tweet
type CreateTweetData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// CreateTweetResponse is the response returned by the create tweet
type CreateTweetResponse struct {
	Tweet *CreateTweetData `json:"data"`
}

// DeleteTweetData is the indication of the deletion of tweet
type DeleteTweetData struct {
	Deleted bool `json:"deleted"`
}

// DeleteTweetResponse is the response returned by the delete tweet
type DeleteTweetResponse struct {
	Tweet *DeleteTweetData `json:"data"`
}
