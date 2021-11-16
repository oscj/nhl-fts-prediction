# nhl-over-under

### Table of Contents
1. [Project Background and Context](#project-background-and-context)
2. [Data Generation](#data-generation)
3. [Model Implementations](#model-implementations)
4. [Results](#results)
5. [Future Considerations](#future-considerations)
6. [Resources](#resources)

### Project Background and Context

This repository holds implementations of several machine learning models and algorithms to predict the total number of goals of a given NHL game.

This project was created as a term project for CMPUT466 (Machine Learning) at the University of Alberta in the Fall of 2021. 

The name of this repository, _nhl_over_under_, alludes that this is an idea that can be integrated with over-under sports betting strategies.

Over/Under betting in hockey is wagering on the total combined score of a hockey game. Oddsmakers set an expected total number of goals to be scored by both sides. Bettors can wager on whether the final score will go Over or Under that total. - [covers.com](https://www.covers.com/nhl/how-to-bet-hockey)

This project aims to explore the data and variables that impact the final score of a NHL game and see how well an automated betting strategy(ies) would work with historical betting data.

### Data Generation

This project is utilizing the NHL API to pull data for exploration and training purposes. Making API calls is time expensive, so the language of choice for this part of the project was golang. Using go's concurrency features we can quickly and easily pull ten thousand detailed game records in < 1 min.

To generate gamedata json files for each individual game between the 2014-2015 season to the 2020-2021 season, run the following:

```bash
go run gen gen_nhl_game_stats.go
```

### Model Implementations
1. [Closed Form Linear Regression](#closed-form-linear-regression)
2. [Gradient Descent Linear Regression ](#gradient-descent-linear-regression)
3. [Linear Classification](#linear-classification)
4. [Neural Network](#neural-network)

#### Closed Form Linear Regression
// todo

#### Gradient Descent Linear Regression
// todo

#### Linear Classification
// todo

#### Neural Network
// todo

### Future Considerations
// todo

### Project Results
// todo

### Resources
// todo
