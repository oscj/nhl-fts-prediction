#!/usr/bin/python3
import os
import json
import pandas as pd
from tqdm import tqdm

def main():
    gd_path = "/Users/ojaimes/Documents/school/CMPUT466/nhl-over-under/gamedata/"
    gd_files = os.listdir(gd_path)
    
    game_stats = []
    
    for file in tqdm(gd_files):
        with open(gd_path + file, 'r') as f:
            game_data = json.load(f)
            team_stats = game_data['liveData']['boxscore']['teams']
            home_stats = team_stats['home']['teamStats']['teamSkaterStats']
            away_stats = team_stats['away']['teamStats']['teamSkaterStats']
            combined_team_stats = {**home_stats, **away_stats}
            game_stats.append(combined_team_stats)
    
    game_stats_df = pd.DataFrame(game_stats)
    game_stats_df.to_csv('csv/game_stats.csv')
    print("Saved game stats to csv")
       
if __name__ == '__main__':
    main()
