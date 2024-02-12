package model

type ID struct {
	Oid string `json:"$oid"`
}

type Acquisition struct {
	PriceAmount       int    `json:"price_amount"`
	PriceCurrencyCode string `json:"price_currency_code"`
	TermCode          string `json:"term_code"`
	SourceURL         string `json:"source_url"`
	SourceDescription string `json:"source_description"`
	AcquiredYear      int    `json:"acquired_year"`
	AcquiredMonth     int    `json:"acquired_month"`
	AcquiredDay       int    `json:"acquired_day"`
	AcquiringCompany  struct {
		Name      string `json:"name"`
		Permalink string `json:"permalink"`
	} `json:"acquiring_company"`
}

type Acquisitions []struct {
	PriceAmount       int    `json:"price_amount"`
	PriceCurrencyCode string `json:"price_currency_code"`
	TermCode          string `json:"term_code"`
	SourceURL         string `json:"source_url"`
	SourceDescription string `json:"source_description"`
	AcquiredYear      int    `json:"acquired_year"`
	AcquiredMonth     int    `json:"acquired_month"`
	AcquiredDay       int    `json:"acquired_day"`
	Company           struct {
		Name      string `json:"name"`
		Permalink string `json:"permalink"`
	} `json:"company"`
}

type Competitions []struct {
	Competitor struct {
		Name      string `json:"name"`
		Permalink string `json:"permalink"`
	} `json:"competitor"`
}

type CreatedAt struct {
	Date int64 `json:"$date"`
}

type ExternalLink struct {
	ExternalURL string `json:"external_url"`
	Title       string `json:"title"`
}

type VideoEmbed struct {
	EmbedCode   string `json:"embed_code"`
	Description string `json:"description"`
}

type Office struct {
	Description string  `json:"description"`
	Address1    string  `json:"address1"`
	Address2    string  `json:"address2"`
	ZipCode     string  `json:"zip_code"`
	City        string  `json:"city"`
	StateCode   string  `json:"state_code"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Product struct {
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
}

type Company struct {
	ID                ID             `json:"_id"`
	Acquisition       Acquisition    `json:"acquisition"`
	Acquisitions      Acquisitions   `json:"acquisitions"`
	AliasList         string         `json:"alias_list"`
	BlogFeedURL       string         `json:"blog_feed_url"`
	BlogURL           string         `json:"blog_url"`
	CategoryCode      string         `json:"category_code"`
	Competitions      Competitions   `json:"competitions"`
	CreatedAt         CreatedAt      `json:"created_at"`
	CrunchbaseURL     string         `json:"crunchbase_url"`
	DeadpooledYear    int            `json:"deadpooled_year"`
	Description       string         `json:"description"`
	EmailAddress      string         `json:"email_address"`
	ExternalLinks     []ExternalLink `json:"external_links"`
	FoundedDay        int            `json:"founded_day"`
	FoundedMonth      int            `json:"founded_month"`
	FoundedYear       int            `json:"founded_year"`
	HomepageURL       string         `json:"homepage_url"`
	Name              string         `json:"name"`
	NumberOfEmployees int            `json:"number_of_employees"`
	Offices           []Office       `json:"offices"`
	Overview          string         `json:"overview"`
	Permalink         string         `json:"permalink"`
	PhoneNumber       string         `json:"phone_number"`
	Products          []Product      `json:"products"`
	TagList           string         `json:"tag_list"`
	TotalMoneyRaised  string         `json:"total_money_raised"`
	TwitterUsername   string         `json:"twitter_username"`
	UpdatedAt         string         `json:"updated_at"`
	VideoEmbeds       []VideoEmbed   `json:"video_embeds"`
}
