# filename = elo_model.py
# author = John Brichetto
# purpose = Simulate a game between two teams (or individuals) comprised of players of random and unknown skill levels.
#           Accurately determine players relative skill level in as few games as possible.

import random

# Global Constants - Alter these variables to see how they affect an ELO estimation of a player's skill.

TOTAL_GAMES = 10000  # TOTAL number of games to be simulated - NOT the number of games each player will play.
TOTAL_PLAYERS = 1000  # Total number of players to simulate. TOTAL_PLAYERS MUST BE >= 2*TEAM_SIZE
TEAM_SIZE = 5  # The number of players on each team.
ELO_SPREAD = 400  # The Elo difference required for a player to be an order of magnitude more likely to win.
ELO_WEIGHT = 32  # Sets the "weight" of each match. Winning an even match causes a gain of 0.5 * ELO_WEIGHT
UN_RANKED_ELO = 1500.00  # The starting elo of all players.
DISPLAY_DATA_REPORT = False  # True will cause a printout of all players skill levels, elo score, and games played
DISPLAY_DATA_ANALYSIS = True  # True will cause a printout of the analysis


class Player:
    def __init__(self, name):
        self.name = name
        self.skill = random.randint(0, 10)
        self.elo = UN_RANKED_ELO
        self.games_played = 0

    def elo_change(self, change):
        self.elo += change
        self.games_played += 1

    def get_skill(self): return self.skill

    def get_elo(self): return self.elo

    def get_name(self): return self.name

    def get_games_played(self): return self.games_played


# Create Players
player_pool = []
for i in range(0, TOTAL_PLAYERS):
    player_pool.append(Player("p"+str(i)))

# Create two empty teams
team1 = []
team2 = []

for j in range(0, TOTAL_GAMES):
    # Fill the teams with players
    while len(team1) < TEAM_SIZE:
        team1.append(player_pool.pop(random.randint(0, (len(player_pool) - 1))))
        team2.append(player_pool.pop(random.randint(0, (len(player_pool) - 1))))

    # Initialize / Reset elo and skill
    team1_skill, team1_elo = 0, 0
    team2_skill, team2_elo = 0, 0

    # Calculate each team's skill and average elo
    for i in range(0, len(team1)):
        team1_elo += team1[i].get_elo()
        team1_skill += team1[i].get_skill()
        team2_elo += team2[i].get_elo()
        team2_skill += team2[i].get_skill()
    team1_elo = team1_elo / len(team1)
    team2_elo = team2_elo / len(team2)

    # Determine tug of war winner
    if team1_skill > team2_skill:
        team1_win = 1.0
    elif team1_skill == team2_skill:
        team1_win = 0.5
    else:
        team1_win = 0.0

    # Calculate elo change
    n = ELO_SPREAD  # n defines the elo difference needed to be an order of magnitude more likely to win
    k = ELO_WEIGHT  # k defines the weight of each match
    r1 = 10**(team1_elo / n)
    r2 = 10**(team2_elo / n)
    t1_win_prob = r1 / (r1+r2)
    t2_win_prob = r2 / (r1+r2)
    t1_elo_change = k*(team1_win - t1_win_prob)  # if team 1 wins, this creates a positive number (elo gain)
    t2_elo_change = k*(abs(team1_win - 1.0) - t2_win_prob)  # if team 1 wins, this creates a negative number (elo loss)

    # Apply elo change to each player and update games played
    for i in range(0, len(team1)):
        team1[i].elo_change(t1_elo_change)
        team2[i].elo_change(t2_elo_change)

    # Return Players to the player_pool
    for i in range(0, len(team1)):
        player_pool.append(team1.pop())
        player_pool.append(team2.pop())

# Organize the Player Pool
if DISPLAY_DATA_REPORT:
    organization_counter = 0
    for i in range(0, 10):
        for j in range(0, len(player_pool) - 1):
            if player_pool[j].get_skill() == (10 - i):
                player_pool.insert(organization_counter, player_pool.pop(j))
                organization_counter += 1
            else:
                pass

# Report Data
if DISPLAY_DATA_REPORT:
    for i in range(0, len(player_pool) - 1):
        print("Name: ", player_pool[i].get_name(), "\t\tSkill: ", player_pool[i].get_skill(),
              "\t\tElo : %.2f" % player_pool[i].get_elo(), "\t\tGames Played:  ", player_pool[i].get_games_played())
else:
    pass

# Analyze Data
if DISPLAY_DATA_ANALYSIS:
    for i in range(0, 11):
        analysis_elo_list = []
        analysis_games_played_list = []
        for j in range(0, len(player_pool) - 1):
            if player_pool[j].get_skill() == (10 - i):
                analysis_elo_list.append(int(player_pool[j].get_elo()))
                analysis_games_played_list.append(player_pool[j].get_games_played())
                analysis_elo_list.sort()
            else:
                pass

        print("Skill Level: ", (10 - i),
              "\tHighest Elo: ", analysis_elo_list[(len(analysis_elo_list)-1)],
              "\tLowest Elo: ", analysis_elo_list[0],
              "\tAverage Elo: %.2f" % (sum(analysis_elo_list) / len(analysis_elo_list)),
              "\tPlayers this skill level: ", len(analysis_elo_list),
              "\tAverage Number of Games Played: %.2f"
              % (sum(analysis_games_played_list) / len(analysis_games_played_list))
              )
