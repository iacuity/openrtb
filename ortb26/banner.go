package ortb26

import "encoding/json"

// Banner object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	Formats      []Format            `json:"format,omitempty"`   // Array of format objects representing the banner sizes permitted.
	Width        int                 `json:"w,omitempty"`        // Width
	Height       int                 `json:"h,omitempty"`        // Height
	BlockedTypes []BannerType        `json:"btype,omitempty"`    // Blocked banner types
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"`    // Blocked creative attributes
	Position     AdPosition          `json:"pos,omitempty"`      // Ad Position
	MIMEs        []string            `json:"mimes,omitempty"`    // Whitelist of content MIME types supported
	TopFrame     int                 `json:"topframe,omitempty"` // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	ExpDirs      []ExpDir            `json:"expdir,omitempty"`   // Specify properties for an expandable ad
	APIs         []APIFramework      `json:"api,omitempty"`      // List of supported API frameworks
	ID           string              `json:"id,omitempty"`       // A unique identifier. Recommended when Banner objects are used with a Video object
	VCM          int                 `json:"vcm,omitempty"`      // Represents the relationship with video. 0 = concurrent, 1 = end-card
	ETime        int                 `json:"etime,omitempty"`    // The exposure time in seconds per view that the creative will be displayed before refreshing to the next creative
	Ext          json.RawMessage     `json:"ext,omitempty"`
}
