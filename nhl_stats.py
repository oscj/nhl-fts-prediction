#!/usr/local/bin/python3

from typing import Dict, List
import requests
import pprint

BASE_URL = "https://statsapi.web.nhl.com/api/v1"


def get_teams() -> List:
    """
    Returns a list of all team objects
    """
    res = requests.get(BASE_URL + "/teams")

    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["teams"]


def get_team_roster(team_id: int) -> List:
    """
    Returns a list containing the roster of a given team id.
    
    Params:
        team_id: The team to get the roster of
    """
    # ids only range from 1 to 55
    if team_id not in range(1, 56):
        return []

    res = requests.get(BASE_URL + "/teams/{}/roster".format(team_id))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["roster"]


def get_player_info(player_id: int) -> Dict:
    """
    Returns dictionary of player info.
    
    Params:
        player_id: The id of the player to look up
    """
    res = requests.get(BASE_URL + "/people/{}".format(player_id))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["people"][0]


def get_player_stats_for_season(player_id: int, season: int) -> List:
    """
    Returns list containing data of player info for a given season.
    
    Params:
        player_id: The id of the player to look for
        season: The year of the start of the season (limited to 2000 - 2021)
    """
    if season not in range(2000, 2022):
        return []

    res = requests.get(
        BASE_URL + "/people/{}/stats?stats=statsSingleSeason&season={}{}".format(player_id, season, season+1))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["stats"]


def get_game_schedule_between_dates_for_team_id(team_id: str, start_date: str, end_date: str) -> List:
    """
    Returns list of schedule information data for given parameters.
    
    Params:
        team_id: Team whose schedule we are searching for
        start_date: Start date
        end_date: End date
    """
    res = requests.get(
        BASE_URL + "/schedule?teamId={}&startDate={}&endDate={}".format(team_id, start_date, end_date))
    if not res.status_code == 200:
        return {}

    schedule =  res.json()["dates"]
    return schedule


def get_game_info(game_id: str) -> Dict:
    """
    Returns dictionary of game information data for given game ID.
    
    Params:
        game_id: The game to query
    """
    res = requests.get(BASE_URL + "/game/{}/feed/live".format(game_id))
    if not res.status_code == 200:
        return {}

    game_data = res.json()
    return game_data


if __name__ == "__main__":
    '''
    teams = get_teams()
    for team in teams:
        print(team)
        print()
        print(get_team_roster(team["id"]))
        print("\n\n\n")
    '''
    
    print("\n\nPlayer info:")
    pprint.pprint(get_player_info(8476459))
    
    print("\n\n")
    pprint.pprint(get_player_stats_for_season(8476459, 2020))

    print("\n\Schedule info:")
    pprint.pprint(get_game_schedule_between_dates_for_team_id(
        30, "2018-01-02", "2018-01-02"))

    print("\n\nGame info:")
    pprint.pprint(get_game_info(2017020608))
