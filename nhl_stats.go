package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func get_game_stats_and_save_to_disk(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	gameInfo, err := models.UnmarshalGameInfo(body)
	file, err := json.MarshalIndent(gameInfo, "", " ")
	if err != nil {
		return err
	}
	gamePk := strconv.Itoa(int(gameInfo.GamePk))
	fileStr := "gamedata/" + gamePk + ".json"
	fmt.Println(fileStr)
	err = ioutil.WriteFile(fileStr, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	seasons := []string{"20142015", "20152016", "20162017", "20172018", "20182019", "20192020", "20202021"}
	var schedules []models.SeasonSchedule

	// For each season, get the schedule
	for _, season := range seasons {
		schedule, _ := get_game_schedule(season)
		ss, _ := models.UnmarshalSeasonSchedule(schedule)
		schedules = append(schedules, ss)
	}

	// For each game in each season schedule, get the api endpoint for the game stats
	var endpoints []string
	for i, schedule := range schedules {
		fmt.Println(i, seasons[i], schedule.TotalGames)
		for _, dates := range schedule.Dates {
			for _, game := range dates.Games {
				gamepk := strconv.Itoa(int(game.GamePk))
				endpoints = append(endpoints, "https://statsapi.web.nhl.com/api/v1/game/"+gamepk+"/feed/live")
			}
		}
	}

	fmt.Println()

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

	numJobs := len(endpoints)
	jobs := make(chan string, numJobs)
	results := make(chan error, numJobs)
	numworkers := 50

	// create pool of workers
	for w := 0; w < numworkers; w++ {
		go worker(w, jobs, results)
	}

	// give work to workers (send endpoints to channel)
	for j := 0; j < numJobs; j++ {
		jobs <- endpoints[j]
	}

	for a := 0; a < numJobs; a++ { // wait for all jobs to finish
		fmt.Println(<-results)
	}
}

func worker(id int, jobs chan string, results chan error) {
	for j := range jobs {
		fmt.Printf("worker: %d, got job: %s\n", id, j)
		err := get_game_stats_and_save_to_disk(j)
		results <- err
	}
}
