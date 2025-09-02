// --- Proposed Data Structures ---

// Holds info about the currently logged-in user
type AuthenticatedUser struct {
    IsLoggedIn  bool
    Username    string
    ProfilePic  string // URL to their picture
}

// Represents a single user for profile pages or match feeds
type UserProfile struct {
    Username    string
    Age         int
    Handicap    float64
    Bio         string
    Photos      []string // A list of image URLs
}

// Represents a single golf course
type GolfCourse struct {
    Name        string
    Location    string
    Par         int
    Rating      float64
    ImageURL    string
}

// Represents a scheduled tee time
type TeeTime struct {
    CourseName  string
    Date        string // Or time.Time for more control
    Time        string
    Players     []UserProfile // Profiles of who is playing
}


// The "Evolved" GolfPageData struct
type GolfPageData struct {
    // --- 1. Basic Page Info (What you have now) ---
    PageTitle   string
    
    // --- 2. Who is logged in? ---
    CurrentUser AuthenticatedUser

    // --- 3. What is the specific content for THIS page? ---
    // We use different fields for different pages. Only one will be filled at a time.
    Users     []UserProfile // For the matching feed or contact list
    Profile   UserProfile   // For a single user's profile page
    Courses   []GolfCourse  // For the course selection page
    Schedule  []TeeTime     // For the calendar/schedule page
}
