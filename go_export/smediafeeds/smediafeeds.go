package smediafeeds


// SocialMedia interface for dynamic types
type SocialMedia interface {
	Feed() []string
	Fame() int
}

// Facebook struct type to implement interface
type Facebook struct {
	URL     string 
	Name    string
	Friends int 
}

// Twitter struct type to implement interface
type Twitter struct {
	URL       string
	Username  string
	Followers int
}

// Linkedin struct type to implement interface
type Linkedin struct {
	URL         string
	Name        string
	Connections int
}

// Feed returns top Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook Feeds", 
		"Tech will play a major role in the ultimate uprising",
		"Stay Safe"}
}

// Fame shows how popular a user is. Higher number of friends higher fame
func (f *Facebook) Fame() int {
	return f.Friends
}

// Feed returns top Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter Feeds",
		"My first tweet",
		"BOPDaddyChallenge"}
}

// Fame shows how popular a user is. Higher number of followers higher fame
func (t *Twitter) Fame() int {
	return t.Followers
}

// Feed returns top Linkedin posts
func (l *Linkedin) Feed() []string {
	return []string{
		"Linkedin Feeds",
		"My career journey",
		"5 year anniversary at HNG"}
}

// Fame shows how popular a user is. Higher number of connections higher fame
func (l *Linkedin) Fame() int {
	return l.Connections
}


