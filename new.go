package goddit

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListingPost struct {
	AuthorFlairBackgroundColor string  `json:"author_flair_background_color"`
	ApprovedAtUtc              float64 `json:"approved_at_utc"`
	Subreddit                  string  `json:"subreddit"`
	Selftext                   string  `json:"selftext"`
	AuthorFullname             string  `json:"author_fullname"`
	Saved                      bool    `json:"saved"`
	ModReasonTitle             any     `json:"mod_reason_title"`
	Gilded                     int     `json:"gilded"`
	Clicked                    bool    `json:"clicked"`
	Title                      string  `json:"title"`
	LinkFlairRichtext          []struct {
		E string `json:"e"`
		T string `json:"t"`
	} `json:"link_flair_richtext"`
	SubredditNamePrefixed string  `json:"subreddit_name_prefixed"`
	Hidden                bool    `json:"hidden"`
	Pwls                  int     `json:"pwls"`
	LinkFlairCSSClass     string  `json:"link_flair_css_class"`
	Downs                 int     `json:"downs"`
	ThumbnailHeight       int     `json:"thumbnail_height"`
	TopAwardedType        any     `json:"top_awarded_type"`
	HideScore             bool    `json:"hide_score"`
	Name                  string  `json:"name"`
	Quarantine            bool    `json:"quarantine"`
	LinkFlairTextColor    string  `json:"link_flair_text_color"`
	UpvoteRatio           float64 `json:"upvote_ratio"`
	IgnoreReports         bool    `json:"ignore_reports"`
	Ups                   int     `json:"ups"`
	Domain                string  `json:"domain"`
	MediaEmbed            struct {
	} `json:"media_embed"`
	ThumbnailWidth        int    `json:"thumbnail_width"`
	AuthorFlairTemplateID string `json:"author_flair_template_id"`
	IsOriginalContent     bool   `json:"is_original_content"`
	UserReports           []any  `json:"user_reports"`
	SecureMedia           any    `json:"secure_media"`
	IsRedditMediaDomain   bool   `json:"is_reddit_media_domain"`
	IsMeta                bool   `json:"is_meta"`
	Category              any    `json:"category"`
	SecureMediaEmbed      struct {
	} `json:"secure_media_embed"`
	LinkFlairText       string `json:"link_flair_text"`
	CanModPost          bool   `json:"can_mod_post"`
	Score               int    `json:"score"`
	ApprovedBy          any    `json:"approved_by"`
	IsCreatedFromAdsUI  bool   `json:"is_created_from_ads_ui"`
	AuthorPremium       bool   `json:"author_premium"`
	Thumbnail           string `json:"thumbnail"`
	Edited              bool   `json:"edited"`
	AuthorFlairCSSClass any    `json:"author_flair_css_class"`
	AuthorFlairRichtext []struct {
		E string `json:"e"`
		T string `json:"t,omitempty"`
		A string `json:"a,omitempty"`
		U string `json:"u,omitempty"`
	} `json:"author_flair_richtext"`
	Gildings struct {
	} `json:"gildings"`
	PostHint            string  `json:"post_hint"`
	ContentCategories   any     `json:"content_categories"`
	IsSelf              bool    `json:"is_self"`
	SubredditType       string  `json:"subreddit_type"`
	Created             float64 `json:"created"`
	LinkFlairType       string  `json:"link_flair_type"`
	Wls                 int     `json:"wls"`
	RemovedByCategory   any     `json:"removed_by_category"`
	BannedBy            any     `json:"banned_by"`
	AuthorFlairType     string  `json:"author_flair_type"`
	TotalAwardsReceived int     `json:"total_awards_received"`
	AllowLiveComments   bool    `json:"allow_live_comments"`
	SelftextHTML        any     `json:"selftext_html"`
	Likes               any     `json:"likes"`
	SuggestedSort       any     `json:"suggested_sort"`
	BannedAtUtc         float64 `json:"banned_at_utc"`
	URLOverriddenByDest string  `json:"url_overridden_by_dest"`
	ViewCount           any     `json:"view_count"`
	Archived            bool    `json:"archived"`
	NoFollow            bool    `json:"no_follow"`
	Spam                bool    `json:"spam"`
	IsCrosspostable     bool    `json:"is_crosspostable"`
	Pinned              bool    `json:"pinned"`
	Over18              bool    `json:"over_18"`
	Preview             struct {
		Images []struct {
			Source struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"source"`
			Resolutions []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"resolutions"`
			Variants struct {
			} `json:"variants"`
			ID string `json:"id"`
		} `json:"images"`
		Enabled bool `json:"enabled"`
	} `json:"preview"`
	AllAwardings             []any   `json:"all_awardings"`
	Awarders                 []any   `json:"awarders"`
	MediaOnly                bool    `json:"media_only"`
	LinkFlairTemplateID      string  `json:"link_flair_template_id"`
	CanGild                  bool    `json:"can_gild"`
	Removed                  bool    `json:"removed"`
	Spoiler                  bool    `json:"spoiler"`
	Locked                   bool    `json:"locked"`
	AuthorFlairText          string  `json:"author_flair_text"`
	TreatmentTags            []any   `json:"treatment_tags"`
	Visited                  bool    `json:"visited"`
	RemovedBy                any     `json:"removed_by"`
	ModNote                  any     `json:"mod_note"`
	Distinguished            any     `json:"distinguished"`
	SubredditID              string  `json:"subreddit_id"`
	AuthorIsBlocked          bool    `json:"author_is_blocked"`
	ModReasonBy              any     `json:"mod_reason_by"`
	NumReports               int     `json:"num_reports"`
	RemovalReason            any     `json:"removal_reason"`
	LinkFlairBackgroundColor string  `json:"link_flair_background_color"`
	ID                       string  `json:"id"`
	IsRobotIndexable         bool    `json:"is_robot_indexable"`
	ReportReasons            []any   `json:"report_reasons"`
	Author                   string  `json:"author"`
	DiscussionType           any     `json:"discussion_type"`
	NumComments              int     `json:"num_comments"`
	SendReplies              bool    `json:"send_replies"`
	ContestMode              bool    `json:"contest_mode"`
	ModReports               []any   `json:"mod_reports"`
	AuthorPatreonFlair       bool    `json:"author_patreon_flair"`
	Approved                 bool    `json:"approved"`
	AuthorFlairTextColor     string  `json:"author_flair_text_color"`
	Permalink                string  `json:"permalink"`
	Stickied                 bool    `json:"stickied"`
	URL                      string  `json:"url"`
	SubredditSubscribers     int     `json:"subreddit_subscribers"`
	CreatedUtc               float64 `json:"created_utc"`
	NumCrossposts            int     `json:"num_crossposts"`
	Media                    any     `json:"media"`
	IsVideo                  bool    `json:"is_video"`
}

type PostListing struct {
	Kind string `json:"kind"`
	Data struct {
		After     string `json:"after"`
		Dist      int    `json:"dist"`
		Modhash   any    `json:"modhash"`
		GeoFilter string `json:"geo_filter"`
		Children  []struct {
			Kind string      `json:"kind"`
			Data ListingPost `json:"data"`
		} `json:"children"`
		Before any `json:"before"`
	} `json:"data"`
}

func (r *Reddit) NewAfter(subreddit, after string, limit int, show bool) []ListingPost {
	parameters := fmt.Sprintf("after=%s&limit=%d", after, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) NewBefore(subreddit, before string, limit int, show bool) []ListingPost {
	parameters := fmt.Sprintf("before=%s&limit=%d", before, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) New(subreddit string, limit int, show bool) []ListingPost {
	parameters := fmt.Sprintf("limit=%d", limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
	return r.new(subreddit, parameters)
}

func (r *Reddit) new(subreddit, parameters string) []ListingPost {
	endpoint := fmt.Sprintf("r/%s/new?%s", subreddit, parameters)
	listingString := r.request("GET", endpoint, "")
	var listing PostListing
	err := json.Unmarshal([]byte(listingString), &listing)
	if err != nil {
		log.Fatal(err)
	}
	var result []ListingPost
	for _, s := range listing.Data.Children {
		result = append(result, s.Data)
	}
	return result
}
