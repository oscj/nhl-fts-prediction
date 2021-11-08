# nhl-over-under

Implementation of several machine learning models to predict the total number of goals in a NHL game.

### Data Generation

This project is utilizing the NHL API to pull data for exploration and training purposes. Making API calls is time expensive, so the language of choice for this part of the project was golang. Using go's concurrency features we can quickly pull ten thousand detailed game records in < 1 min.

To generate gamedata json files for each individual game between the 2014-2015 season to the 2020-2021 season, run the following:

```bash
go run gen gen_nhl_game_stats.go
```
