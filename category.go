package writeas

// Category represents a post tag with additional metadata, like a title and slug.
type Category struct {
	Hashtag string `json:"hashtag"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
}
