package station

type ThemeMode string

const (
	ThemeLight ThemeMode = "light"
	ThemeDark  ThemeMode = "dark"
)

type Image struct {
	AssetID string `json:"asset_id"`
	Alt     string `json:"alt"`
}

type Plant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LatinName   string `json:"latin_name"`
	Observation string `json:"observation"`
	Image       Image  `json:"image"`
}

type VolunteerStory struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Image   Image  `json:"image"`
}

type Activity struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Schedule  string `json:"schedule"`
	Location  string `json:"location"`
	Capacity  int    `json:"capacity"`
	SignupURL string `json:"signup_url"`
}

type Article struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Summary  string   `json:"summary"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Image    Image    `json:"image"`
}

type Theme struct {
	Mode           ThemeMode `json:"mode"`
	Background     string    `json:"background"`
	Text           string    `json:"text"`
	Border         string    `json:"border"`
	ImageTreatment string    `json:"image_treatment"`
}

type AppliedFilters struct {
	Query string   `json:"query"`
	Tags  []string `json:"tags"`
}

type HomePage struct {
	StationName      string           `json:"station_name"`
	MonthLabel       string           `json:"month_label"`
	Plants           []Plant          `json:"plants"`
	VolunteerStories []VolunteerStory `json:"volunteer_stories"`
	Activities       []Activity       `json:"activities"`
	Categories       []string         `json:"categories"`
	Articles         []Article        `json:"articles"`
	AvailableTags    []string         `json:"available_tags"`
	Filters          AppliedFilters   `json:"filters"`
	Theme            Theme            `json:"theme"`
}

type HomeRequest struct {
	Query string
	Tags  []string
	Mode  ThemeMode
}

type RegisterRequest struct {
	ActivityID string `json:"activity_id"`
	Name       string `json:"name"`
	Contact    string `json:"contact"`
}

type Registration struct {
	ID         string `json:"id"`
	ActivityID string `json:"activity_id"`
	Name       string `json:"name"`
	Contact    string `json:"contact"`
	Status     string `json:"status"`
}

type RegistrationResult struct {
	Registration Registration `json:"registration"`
	Message      string       `json:"message"`
}
