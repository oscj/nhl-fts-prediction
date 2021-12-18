#!/usr/bin/python3
"""
parse_nhl_odds.py
-----------------
Script to parse all NHL betting data CSVs in the /degeneracy/bet_csvs/ directory and save to a singular csv file.
"""
import os
import pandas as pd
from tqdm import tqdm

def main():
    gd_path = "/Users/ojaimes/Documents/school/CMPUT466/nhl-over-under/degeneracy/bet_csvs/"
    gd_files = os.listdir(gd_path)
    
    betting_data_dfs = []
    
    for file in tqdm(gd_files):
        betting_data = pd.read_csv(gd_path + file)
        betting_data_dfs.append(betting_data)
            
    # Combine all betting data into one dataframe
    betting_data_df = pd.concat(betting_data_dfs)
    betting_data_df.to_csv('/Users/ojaimes/Documents/school/CMPUT466/nhl-over-under/degeneracy/bet_csvs/all_betting_data.csv')
    print("Saved betting data to csv")
       
if __name__ == '__main__':
    main()
