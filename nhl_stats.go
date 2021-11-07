package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/oscjaimes/nhl-over-under/models"
)

func get_game_schedule(season string) ([]byte, error) {
	url := "https://statsapi.web.nhl.com/api/v1/schedule?season=" + season
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, nil
}

// func get_game_stats(game_id string) {
// 	url := "https://statsapi.web.nhl.com/api/v1/game/" + game_id + "/feed/live"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println(resp.Status)
// }

// 1. get all seasons, fast b/c only a few api calls -> collect in an array
// 2. format data -> array of urls you will call
// 3. create worker pool: fixed num == num of cores,
//   create channel
//   for i in numcores
//     run goroutine, pass in channel
//       iterate over channel, this will loop until channel closes -> gives urls
//       do api call, save to disk using a unique filename based on the url
//         make func that does call, something like:
//			for { res := http.Get(...); if res.statusCode / 100 == 2 { unmarshal and either return string/array or write straight to disk} else sleep}
// 4. in main, iterate over array, pass in urls into channel
// 5. sleep for n milliseconds, so you dont overload server. This is your "bottleneck"

func main() {
	seasons := []string{"20142015", "20152016", "20162017", "20172018", "20182019", "20192020", "20202021"}

	for _, season := range seasons {
		schedule, _ := get_game_schedule(season)
		ss, _ := models.UnmarshalSeasonSchedule(schedule)
		fmt.Println(ss.TotalGames)
	}

}
