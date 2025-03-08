package goddit

import (
	"fmt"
    "encoding/json"
    "log"
)

type ListingComment struct {
	LinkURL             string  `json:"link_url"`
	SubredditID         string  `json:"subreddit_id"`
	ApprovedAtUtc       any     `json:"approved_at_utc"`
	AuthorIsBlocked     bool    `json:"author_is_blocked"`
	CommentType         any     `json:"comment_type"`
	LinkTitle           string  `json:"link_title"`
	ModReasonBy         any     `json:"mod_reason_by"`
	BannedBy            any     `json:"banned_by"`
	Ups                 int     `json:"ups"`
	NumReports          int     `json:"num_reports"`
	AuthorFlairType     string  `json:"author_flair_type"`
	TotalAwardsReceived int     `json:"total_awards_received"`
	Subreddit           string  `json:"subreddit"`
	LinkAuthor          string  `json:"link_author"`
	Distinguished       any     `json:"distinguished"`
	Likes               any     `json:"likes"`
	Replies             string  `json:"replies"`
	UserReports         []any   `json:"user_reports"`
	Saved               bool    `json:"saved"`
	ID                  string  `json:"id"`
	BannedAtUtc         any     `json:"banned_at_utc"`
	ModReasonTitle      any     `json:"mod_reason_title"`
	Gilded              int     `json:"gilded"`
	Archived            bool    `json:"archived"`
	CollapsedReasonCode any     `json:"collapsed_reason_code"`
	NoFollow            bool    `json:"no_follow"`
	Spam                bool    `json:"spam"`
	NumComments         int     `json:"num_comments"`
	CanModPost          bool    `json:"can_mod_post"`
	CreatedUtc          float64 `json:"created_utc"`
	SendReplies         bool    `json:"send_replies"`
	ParentID            string  `json:"parent_id"`
	Score               int     `json:"score"`
	AuthorFullname      string  `json:"author_fullname"`
	Over18              bool    `json:"over_18"`
	ReportReasons       []any   `json:"report_reasons"`
	ApprovedBy          any     `json:"approved_by"`
	Controversiality    int     `json:"controversiality"`
	Body                string  `json:"body"`
	Edited              bool    `json:"edited"`
	Gildings            struct {
	} `json:"gildings"`
	AuthorFlairCSSClass          any     `json:"author_flair_css_class"`
	AuthorFlairBackgroundColor   any     `json:"author_flair_background_color"`
	IsSubmitter                  bool    `json:"is_submitter"`
	Downs                        int     `json:"downs"`
	AuthorFlairRichtext          []any   `json:"author_flair_richtext"`
	AuthorPatreonFlair           bool    `json:"author_patreon_flair"`
	BodyHTML                     string  `json:"body_html"`
	RemovalReason                any     `json:"removal_reason"`
	CollapsedReason              any     `json:"collapsed_reason"`
	LinkID                       string  `json:"link_id"`
	AssociatedAward              any     `json:"associated_award"`
	Stickied                     bool    `json:"stickied"`
	AuthorPremium                bool    `json:"author_premium"`
	CanGild                      bool    `json:"can_gild"`
	TopAwardedType               any     `json:"top_awarded_type"`
	UnrepliableReason            any     `json:"unrepliable_reason"`
	Approved                     bool    `json:"approved"`
	AuthorFlairTextColor         any     `json:"author_flair_text_color"`
	ScoreHidden                  bool    `json:"score_hidden"`
	Permalink                    string  `json:"permalink"`
	SubredditType                string  `json:"subreddit_type"`
	LinkPermalink                string  `json:"link_permalink"`
	Name                         string  `json:"name"`
	Created                      float64 `json:"created"`
	SubredditNamePrefixed        string  `json:"subreddit_name_prefixed"`
	AuthorFlairText              any     `json:"author_flair_text"`
	Removed                      bool    `json:"removed"`
	Author                       string  `json:"author"`
	Collapsed                    bool    `json:"collapsed"`
	Awarders                     []any   `json:"awarders"`
	AllAwardings                 []any   `json:"all_awardings"`
	Locked                       bool    `json:"locked"`
	IgnoreReports                bool    `json:"ignore_reports"`
	CollapsedBecauseCrowdControl any     `json:"collapsed_because_crowd_control"`
	ModReports                   []any   `json:"mod_reports"`
	Quarantine                   bool    `json:"quarantine"`
	ModNote                      any     `json:"mod_note"`
	TreatmentTags                []any   `json:"treatment_tags"`
	AuthorFlairTemplateID        any     `json:"author_flair_template_id"`
}

type CommentListing struct {
	Kind string `json:"kind"`
	Data struct {
		After     string `json:"after"`
		Dist      int    `json:"dist"`
		Modhash   any    `json:"modhash"`
		GeoFilter string `json:"geo_filter"`
		Children  []struct {
			Kind string           `json:"kind"`
			Data ListingComment `json:"data"`
		} `json:"children"`
		Before any `json:"before"`
	} `json:"data"`
}

func (r *Reddit) Comments(subreddit, sort string, limit int, show bool) []ListingComment {
	parameters := fmt.Sprintf("sort=%s&limit=%d", sort, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
    return r.comments(subreddit, parameters)
}

func (r *Reddit) CommentsAfter(subreddit, after, sort string, limit int, show bool) []ListingComment {
	parameters := fmt.Sprintf("sort=%s&after=%s&limit=%d", sort, after, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
    return r.comments(subreddit, parameters)
}

func (r *Reddit) CommentsBefore(subreddit, before, sort string, limit int, show bool) []ListingComment {
	parameters := fmt.Sprintf("sort=%s&before=%s&limit=%d", sort, before, limit)
	if show {
		parameters = fmt.Sprintf("%s&show=all", parameters)
	}
    return r.comments(subreddit, parameters)
}

func (r *Reddit) comments(subreddit string, parameters string) []ListingComment{
	endpoint := fmt.Sprintf("r/%s/comments?%s", subreddit, parameters)
    listingString := r.request("GET", endpoint, "")
    var listing CommentListing
    err := json.Unmarshal([]byte(listingString), &listing)
    if err != nil {
        log.Fatal(err)
    }
    var result []ListingComment
    for _, s := range listing.Data.Children {
        result = append(result, s.Data)
    }
    return result
}
