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


**Running Tests**
In Database_test.go, there are three simple tests that tests some functions in Database.go. When we go test the project
and in database.go with lines 101 to 106 active, I get an error. In attempt to omit that code through an if statement, I wrote
isGamedataEmpty to skip that code if the code is false. Unfortunately I get an error similar as I would without it. The only
to get the tests running without the panic runtime error below, I just have to remove all the code overall.

=== RUN   TestCreateTables
--- FAIL: TestCreateTables (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x20 pc=0x23b9d0]

goroutine 9 [running]:
testing.tRunner.func1.1(0x522420, 0x4c8310)
	C:/Go/src/testing/testing.go:1072 +0x310
testing.tRunner.func1(0xc000514300)
	C:/Go/src/testing/testing.go:1075 +0x43a
panic(0x522420, 0x4c8310)
	C:/Go/src/runtime/panic.go:969 +0x1c7
database/sql.(*DB).conn(0x0, 0x5eb960, 0xc000106080, 0x701, 0xc003e96000, 0x74f, 0xc003e96000)
	C:/Go/src/database/sql/sql.go:1190 +0x50
database/sql.(*DB).exec(0x0, 0x5eb960, 0xc000106080, 0xc003e96000, 0x74f, 0x0, 0x0, 0x0, 0xc000089d01, 0x33d9e5, ...)
	C:/Go/src/database/sql/sql.go:1546 +0x6d
database/sql.(*DB).ExecContext(0x0, 0x5eb960, 0xc000106080, 0xc003e96000, 0x74f, 0x0, 0x0, 0x0, 0x4, 0x57aca7, ...)
	C:/Go/src/database/sql/sql.go:1528 +0xe5
database/sql.(*DB).Exec(...)
	C:/Go/src/database/sql/sql.go:1542
Comp510_Project4_HuyLe.createTables(0x0, 0xc03361ea88, 0x1e5726, 0x6b0e60)
	C:/Users/HuyLe/go/src/Comp510_Project4_HuyLe/Database.go:102 +0x432
Comp510_Project4_HuyLe.TestCreateTables(0xc000514300)
	C:/Users/HuyLe/go/src/Comp510_Project4_HuyLe/Database_test.go:10 +0x38
testing.tRunner(0xc000514300, 0x599640)
	C:/Go/src/testing/testing.go:1123 +0xef
created by testing.(*T).Run
	C:/Go/src/testing/testing.go:1168 +0x2b3

FAIL	Comp510_Project4_HuyLe	9.060s

Process finished with exit code 1





Name:       Huy Le
Course:     Comp510
Email :     hle@student.bridgew.edu
