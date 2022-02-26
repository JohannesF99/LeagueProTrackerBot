# League Pro Tracker Bot

This is a Twitter-Bot, who tweets regular updates about the SoloQ of League of Legends Pro-Players. 

Currently, it's only implemented for Germany's <a href="https://twitter.com/PL_Pro_Tracker">PrimeLeague</a>, but it can be used to spectate any region or players.

## How to build the application:

To run the App on your device, run the following command in the working directory.
    
    go get .

This will download any needed remote dependencies. After the download, you may build the application.

    go build .

This should create an executable file depending on your OS in the working directory.

##  Configuration

If you want to use this app for your own Twitter bot, there are a few files which need to be configured.
To use this app, you'll need access to the <a href="https://developer.riotgames.com/">Riot Games<a/> and <a href="https://developer.twitter.com/en">Twitter</a> API.

Once you have all accounts set up, create a configuration file in the config folder.

    touch config/api_details.properties

In this file, save your API credentials using the following template.

    RIOT_API_KEY=<Your Riot API Key>

    TWITTER_API_CONSUMER_KEY=<Your Twitter API Key>
    TWITTER_API_CONSUMER_SECRET=<Your Twitter API Secret>
    TWITTER_API_TOKEN=<Your Twitter Token>
    TWITTER_API_TOKEN_SECRET=<Your API Twitter Token Secret>

Since your App can now fetch data from the Riot Games API 
and send Tweets to the Twitter API, the only thing left is to select
which players you want to track.

In the `config` Folder, you'll see a file named `player_watchlist.json`.
You could tell by the name, that in this file, you declare which players should be listed.

To understand the structure of the file, we need to understand how the Riot Games API works under the hood.
The important keys are `Summoner Name` and `PUUID`. Important to note is that the PUUID never changes and is unique for all players across all riot games.
The problem regarding the PUUID is the fact, that it's usually hidden for the average player.
Thanks to websites like <a href="https://lolpros.gg/">LolPros.gg</a> we can get the Summoner Name for most Pro Players.

To add players to track, you need to add the Summoner Name for every player.
After that, you can update all PUUIDs in the file using

    loader.UpdatePuuidForAllTeams(&teamList)

Now the app is ready to run.

#   Documentation

After the introduction, this section is used to explain the packages and their respective methods.

## API

### LeagueApi.go

### TwitterApi.go

## Loader

### WatchListLoader.go

### RankedDataLoader.go

##  Model

### Player

### Team

### LeagueEntryDTO

### SummonerDTO