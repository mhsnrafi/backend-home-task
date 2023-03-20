# Backend Task




![](https://img.shields.io/badge/GO-1.19-blue.svg)


The goal of the task is to print the players of the given teams. If there is a team overlap (e.g. a player plays in Germany and Chelsea), then print both teams.

## Solution Implementation
The goal of the task is to discover the valid team id we get the desired teams players information.
Following are the steps:
1. First to get the valid team id's we took worker pool approach using goroutines and channels.
   1. To hit the api concurrently untill and unless we get the desired team and the team and player information, We launch the workers which will receive work on the jobs channel and send the corresponding results. We make 2 channels for this.
   2. We launch 10 number of jobs/workers
   3. If we find the desired team result we removed the team names from the map list, if length of the team names map is empty so we save the response in channel and return
2. Once we get find all the valid id's of the listed teams and their player information. We start iteration on the records:
   1. First we create a map of player information and add the records in a desired manner.
   2. Player id is the key of map, so key is unique if next time same player id is coming and we check if this key is already in the map we go to else check and just append the new team name and compiled the record in the existing map we created in if condition
   3. Using this logic no duplicate records can be stored in the map
3. Once all the teams player information is complied we apply the sorting on the player id
4. Once we get the sorted teams player information, we just STDOUT the records
5. During the api hits we just get the 404 error which is handled while making the api requests
6. Logging is properly while we check any error condition, anything we just have to print the information
7. Using modular approach like created a proper utils methods


## First, we need $GOPATH/bin

```bash
export PATH=$PATH:$GOPATH/bin
```

## Clone

```bash
git clone https://github.com/mhsnrafi/backend-home-task.git
```

###### Run

```bash
go run main.go
```
###### Terminal
![](https://i.postimg.cc/P5p9N22V/Screenshot-2022-11-24-at-11-35-35-PM.png)

###### Run Test Case
```bash
go test
```
###### Terminal
![](https://i.postimg.cc/RCYTq9y9/Screenshot-2022-11-24-at-11-34-32-PM.png)
