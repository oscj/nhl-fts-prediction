#!/usr/bin/python3
"""
parse_nhl_game_stats.py
-----------------------
Script to parse all NHL game data CSVs in the /gamedata directory and save to a singular csv file.
"""
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
            
            home_keys = list(home_stats.keys())
            away_keys = list(away_stats.keys())
            
            for hkey, akey in zip(home_keys, away_keys):
                new_hkey = "home_" + hkey
                new_akey = "away_" + akey
                home_stats[new_hkey] = home_stats[hkey]
                away_stats[new_akey] = away_stats[akey]
                del home_stats[hkey]
                del away_stats[akey]
            
            combined_team_stats = {**home_stats, **away_stats}
            game_stats.append(combined_team_stats)
    
    game_stats_df = pd.DataFrame(game_stats)
    game_stats_df.to_csv('csv/game_stats.csv')
    print("Saved game stats to csv")
       
if __name__ == '__main__':
    main()
