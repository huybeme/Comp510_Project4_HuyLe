**Objective**
(Project 2)
Write a an API driven site that you host yourself (you can’t use something like
rapidapi to host it). Your program needs to read the “games-features.xlsx” file and store all the
data into a sqlite3 database. (you can set this up as a separate mini-program that you usually
only run once). Once the data is in the database, design a simple API which will allow users to
search on a game and will return the game’s database entry as a json response.
Hint – check out: https://github.com/avelino/awesome-go#microsoft-excel

**Instructions**
go build the project to start the program. The following pages will do the following in your preferred web browser

http://localhost:8080/gameapi                   -> The homepage, contains a search bar to make an entry - does not work
http://localhost:8080/gameapi/all-entries/      -> Will get all database entries as a JSON response
http://localhost:8080/gameapi/single-entry/{id} -> replace {id} with integer to get that key's element
http://localhost:8080/gameapi/query/{search}    -> replace {search} with a search entry to get any related entries

if the database is empty, in the main.go file, under the main function, uncomment the fillDatabase(GameDataBase) line
and rebuild the project. Filling in the database will take approx. one hour


Name:       Huy Le
Course:     Comp510
Email :     hle@student.bridgew.edu