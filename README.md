# Championship

Build RESTful API in Go and MongoDB The REST API service will expose endpoints to manage a store of leagues and teams. The API will allow to see data on leagues, teams and matchs.The API will also allow to receive alerts when new data are added using Slack and webhooks via a registration.

We worked ~20h each on the project, we learned about mongodb because we had difficulties with it during the 2nd assignement. What was hard was to find json files that we can use for football leagues and using mongodb with golang because their not so much exemple on internet. For this type of project, it it probably easier to choose sql because we have relatonal data structure. We can be limited if we decide to improve our project with MongoDB.

- Heroku
- MongoDB
- Slack Webhooks

The app url is https://championshipleague.herokuapp.com

The home page is https://championshipleague.herokuapp.com/champ

## Leagues

To have information about the leagues use https://championshipleague.herokuapp.com/champ/leagues/ using GET method

To put information about a league use https://championshipleague.herokuapp.com/champ/leagues/ using POST method with the proper .json file

To have information about a league use https://championshipleague.herokuapp.com/champ/leagues/{#id}/{country|name|leagueID|teams} using GET method or just https://championshipleague.herokuapp.com/champ/leagues/{#id} 
ex : id0/country

To Delete a league  use https://championshipleague.herokuapp.com/champ/leagues/delete/{#id} with DELETE method



## Leagues and Matches

To have information about the matches of all leagues just use https://championshipleague.herokuapp.com/champ/matchs/ using GET method

To put information about a league and the matches use https://championshipleague.herokuapp.com/champ/matchs/ using POST method with the proper .json file

For a league https://championshipleague.herokuapp.com/champ/matchs/{#id} using GET method or 
https://championshipleague.herokuapp.com/champ/matchs/{#id}/{name|leagueID|rounds}

To have information about a league matchday use https://championshipleague.herokuapp.com/champ/matchs/{#id}/{#matchday} using GET method (ex : id0/matchday1)

To have information about a matchday use https://championshipleague.herokuapp.com/champ/matchs/{#id}/{#matchday}/{date|team1|team2|score1|score2}

To Delete a league and the list a matches use https://championshipleague.herokuapp.com/champ/matchs/delete/{#id} with DELETE method


## Webhhooks

Webhook League URL : https://championshipleague.herokuapp.com/champ/webhookLeague/ using PUSH method with the proper URL link to the webhook for registration

Same for Match URL : https://championshipleague.herokuapp.com/champ/webhookMatch/

https://championshipleague.herokuapp.com/champ/webhookMatch/{#id} with GET method to see the webhook URL

### With curl

To create a league : curl -d "@file.json" -X POST https://championshipleague.herokuapp.com/champ/leagues

To create matches associate to a league : curl -d "@file.matches.json" -X POST https://championshipleague.herokuapp.com/champ/matchs

(replace "file" by the three first letters of a country)