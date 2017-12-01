package various

import (
	"github.com/go-chat-bot/bot"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "log"
    "fmt"
)

//var db *sql.DB

const (
    dbfile = "./seen.db"
)

func seen(cmd *bot.Cmd) (msg string, err error) {


    if len( cmd.Args) == 0 {
        msg = "\u0001ACTION didn't see no-one.\u0001"
    } else if cmd.Args[0] == "Natasha" {
        msg = "I last saw Natasha 0 seconds ago, saying: 'I last saw Natasha 0 seconds ago, saying: 'I last saw Natasha 0 seconds ago, saying: 'I last saw Natasha 0 seconds ago, saying: 'I last saw Natasha 0 seconds ago, saying: 'I last saw Natasha 0 seconds ago, say..."
    } else {
        username := cmd.Args[0]

        db := openDB()
        stmt, err := db.Prepare("select username, msg, julianday('now')-julianday(seen.timestamp) from seen where username=? collate nocase")
        checkError( err)

        var nick string
        var line string
        var timediff float32

        err = stmt.QueryRow( username).Scan( &nick, &line, &timediff)

        switch {
        case err == sql.ErrNoRows:
            msg = "I have not yet seen "+username+"."
        case err != nil:
            log.Fatal(err)
        default:
            totalseconds := int(timediff*86400)
            minutes := int(totalseconds/60)%60
            hours := int(totalseconds/3600)%24
            days := int(timediff)
            weeks := int(days/7)
            years := int(weeks/52)

            days = days - weeks*7 - years*365
            //seconds := int(totalseconds)%60

            msg = "I last saw "+nick+" "

            if years > 1  { msg = msg + fmt.Sprintf("%d years, ", years) }
            if years == 1 { msg = msg + fmt.Sprintf("%d year, ", years) }
            if weeks > 1  { msg = msg + fmt.Sprintf("%d weeks, ", weeks) }
            if weeks == 1 { msg = msg + fmt.Sprintf("%d week, ", weeks) }
            if days > 1   { msg = msg + fmt.Sprintf("%d days, ", days) }
            if days == 1  { msg = msg + fmt.Sprintf("%d day, ", days) }
            if hours > 1  { msg = msg + fmt.Sprintf("%d hours, ", hours) }
            if hours == 1 { msg = msg + fmt.Sprintf("%d hour, ", hours) }
            if minutes == 1 {
                msg = msg + fmt.Sprintf("%d minute ", minutes)
            } else {
                msg = msg + fmt.Sprintf("%d minutes ", minutes)
            }

            msg = msg + fmt.Sprintf( "ago, saying: '%s'.", line)
        }

        stmt.Close()
        db.Close()
    }

    return
}

func checkError(err error) {
    if err != nil {
        log.Fatal( err)
    }
}

func seenupdate(cmd *bot.PassiveCmd) (msg string, err error) {

    db := openDB()
    tx, err := db.Begin()
    checkError( err)

    stmt, err := tx.Prepare("insert or replace into seen(username, msg) values(?, ?)")
    checkError( err)

    stmt.Exec( cmd.User.Nick, cmd.Raw)
    tx.Commit()
    stmt.Close()

    defer db.Close()

    return
}

func dbDo(query string) {
    db := openDB()
    _, err := db.Exec( query)
    if err != nil {
        log.Printf("%q: %s\n", err, query)
        return
    }
    defer db.Close()
}

func createDB() {

    os.Remove(dbfile)

    db, err := sql.Open("sqlite3", dbfile)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sqlQ := `create table seen (
                id integer not null primary key autoincrement,
                username text,
                msg text,
                timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
            );`

    db.Exec( sqlQ)

    db.Exec( "create unique index 'username' on seen(username collate nocase)")
}

func openDB() (db *sql.DB) {
    if _, err := os.Stat(dbfile); os.IsNotExist(err) {
        createDB()
    }

    db, err := sql.Open("sqlite3", dbfile)
    if err != nil {
        log.Fatal(err)
    }

    return db
}

func init() {

	bot.RegisterCommand(
		"seen",
        "Where are they?!",
		"",
		seen)

    bot.RegisterPassiveCommand(
        "seenupdate",
        seenupdate)
}
