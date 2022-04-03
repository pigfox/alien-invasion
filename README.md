# alien-invasion
Instructions
1. $ go run main.go

Assumptions
1. Not every city has a road to every other cities, dead ends do exist.
2. I created the data map given #1 in Findings section below.
3. Go maps are used for cities and aliens, the IDs for respective instances are the map indices.
4. The use of "-1" indicates that the integer value is not set.

Findings
1. The instructions were "You are given a map containing the names of cities in the non-existent world of X." There was no map provided so I created one with 10 cities. The very limited data set opens the possibility of many of unpredictable outcomes.

Sample Run
$ go run main.go 
Alien invasion starting...
Enter number of aliens you wish to create.
1000000
Berlin(1) has been destroyed by alien 457177 and alien 623734!
New-York(4) has been destroyed by alien 75690 and alien 1343!
London(3) has been destroyed by alien 938006 and alien 96791!
San-fransico(6) has been destroyed by alien 621919 and alien 688301!
Paris(2) has been destroyed by alien 636484 and alien 743453!
Copenhagen​(0) has been destroyed by alien 216334 and alien 801650!
Game over! All aliens have died... 

