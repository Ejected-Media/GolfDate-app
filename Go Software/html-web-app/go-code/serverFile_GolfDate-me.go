# .. www.GolfDate.me  ° // - Website App ~

package main



import (


)


func app_welcome_center_page() {


}

func app_infornation_booth_page() {


}

func app_login_center_page() {


}


func app_about_app_page() {


}

func app_dashboard_interface_page() {



}

func app_player_profile_page() {


}

func app_contact_list_page() {


}

func app_team_chat_page() {


}

func app_date_planner_page() {


}

func app_location_map_page() {


}

func app_schedule_calendar_page() {


}

func app_course_map_page() {


}

func app_score_card_page() {


}

func app_match_history_page() {


}


//  .  html url routes 
//  .  as well as terminal cli logs

func main() {

  appName := "www.GolfDate.me  ° // - Website App ~"


  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/user", indexHandler)
  http.HandleFunc("/account", indexHandler)
  http.HandleFunc("/profile", indexHandler)
  
  http.HandleFunc("/portfolio", indexHandler)
  http.HandleFunc("/resume", indexHandler)
  
  http.HandleFunc("/settings", indexHandler)
  
http.HandleFunc("/settings", indexHandler)
- What - // Match Type °

http.HandleFunc("/settings", indexHandler)
- When - // Tee Time ° 
http.HandleFunc("/settings", indexHandler)
- Where - // Course Location °
http.HandleFunc("/settings", indexHandler)
- Who - // Player Buddy Contact List °
http.HandleFunc("/settings", indexHandler)
- Why - // Reason for Match °
http.HandleFunc("/settings", indexHandler)
- How - // Price and Payment Info °

--- 
_ ` ... ` ~ 

### _ non-user
http.HandleFunc("/settings", indexHandler)
+ Welcome Screen
http.HandleFunc("/settings", indexHandler)
+ Information Booth 
http.HandleFunc("/settings", indexHandler)
+ Login Center
http.HandleFunc("/settings", indexHandler)
+ About App


### _ setting match

http.HandleFunc("/settings", indexHandler)
+ Dashboard Interface 
http.HandleFunc("/settings", indexHandler)
+ Player Profile 
http.HandleFunc("/settings", indexHandler)
+ Contact List
http.HandleFunc("/settings", indexHandler)
+ Team Chat
http.HandleFunc("/settings", indexHandler)
+ Date Planner


### _ tee time

http.HandleFunc("/settings", indexHandler)
+ Location Map
http.HandleFunc("/settings", indexHandler)
+ Schedule Calendar 
http.HandleFunc("/settings", indexHandler)
+ Course Map
http.HandleFunc("/settings", indexHandler)
+ Score Card
http.HandleFunc("/settings", indexHandler)
+ Match History






// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
  
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }

} // ~ func main()
