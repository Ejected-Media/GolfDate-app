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
