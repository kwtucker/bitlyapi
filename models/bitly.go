package models

const (
	BitlyAPIV4 = "https://api-ssl.bitly.com/v4/"
)

// user
type User struct {
	Created          string  `json:"created"`
	Modified         string  `json:"modified"`
	Login            string  `json:"login"`
	IsActive         bool    `json:"is_active"`
	Is2FaEnabled     bool    `json:"is_2fa_enabled"`
	Name             string  `json:"name"`
	Emails           []Email `json:"emails"`
	IsSsoUser        bool    `json:"is_sso_user"`
	DefaultGroupGUID string  `json:"default_group_guid"`
}

type Email struct {
	Email      string `json:"email"`
	IsPrimary  bool   `json:"is_primary"`
	IsVerified bool   `json:"is_verified"`
}

/// group
type Group struct {
	Pagination Pagination `json:"pagination"`
	Links      []Link     `json:"links"`
}

type Link struct {
	References     map[string]string `json:"references"`
	Archived       bool              `json:"archived"`
	Tags           []string          `json:"tags"`
	CreatedAt      string            `json:"created_at"`
	Title          string            `json:"title"`
	Deeplinks      []Deeplink        `json:"deeplinks"`
	CreatedBy      string            `json:"created_by"`
	LongURL        string            `json:"long_url"`
	ClientID       string            `json:"client_id"`
	CustomBitlinks []string          `json:"custom_bitlinks"`
	Link           string            `json:"link"`
	ID             string            `json:"id"`
}

type Deeplink struct {
	Bitlink     string `json:"bitlink"`
	InstallURL  string `json:"install_url"`
	Created     string `json:"created"`
	AppURIPath  string `json:"app_uri_path"`
	Modified    string `json:"modified"`
	InstallType string `json:"install_type"`
	AppGUID     string `json:"app_guid"`
	GUID        string `json:"guid"`
	OS          string `json:"os"`
}

type References struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type Pagination struct {
	Total int64  `json:"total"`
	Size  int64  `json:"size"`
	Prev  string `json:"prev"`
	Page  int64  `json:"page"`
	Next  string `json:"next"`
}

// country
type ClickMetrics struct {
	Units         int64    `json:"units"`
	Facet         string   `json:"facet"`
	UnitReference string   `json:"unit_reference"`
	Unit          string   `json:"unit"`
	Metrics       []Metric `json:"metrics"`
}

type Metric struct {
	Clicks int64  `json:"clicks"`
	Value  string `json:"value"`
}

type ErrorResponse struct {
	Message     string  `json:"message"`
	Errors      []Error `json:"errors"`
	Resource    string  `json:"resource"`
	Description string  `json:"description"`
}

type Error struct {
	Field     string `json:"field"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}
