from typing import List
import requests

BASE_URL = "https://statsapi.web.nhl.com/api/v1"


def get_teams() -> List:
    res = requests.get(BASE_URL + "/teams")

    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["teams"]


def get_team_roster(team_id: int) -> List:
    # ids only range from 1 to 55
    if team_id not in range(1, 56):
        return []

    res = requests.get(BASE_URL + "/teams/{}/roster".format(team_id))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["roster"]


def get_player_info(player_id: int) -> List:
    res = requests.get(BASE_URL + "/people/{}".format(player_id))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["people"]


def get_player_stats_for_season(player_id: int, season: int) -> List:
    if season not in range(2000, 2022):
        return []

    res = requests.get(
        BASE_URL + "/people/{}/stats?stats=statsSingleSeason&season={}{}".format(player_id, season, season+1))
    if not res.status_code == 200:
        return []

    rjson = res.json()
    return rjson["stats"]


if __name__ == "__main__":
    '''
    teams = get_teams()
    for team in teams:
        print(team)
        print()
        print(get_team_roster(team["id"]))
        print("\n\n\n")
    '''

    print(get_player_stats_for_season(8476459, 2020))
