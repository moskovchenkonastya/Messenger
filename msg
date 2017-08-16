package main

import(
	"os"
 	"fmt"
	"bufio"
	"strconv"
)

var user = profile{id: 0, name: "", password: "", pictureURL: ""}

//  Structure for profile of user
type profile struct {

	id		   int
	name	   string
	password   string
	pictureURL string
}

// Structure for list of friends
type listOfFriends struct {
	userID  int
	friendID int
}

// Structure of history of massages
type history struct {
	userID  int
	frendID int
	time 	float64
}

// Structure for database
type authData struct{
	userName     string
	userPassword string
	// passwordHash []byte ?
}

const (
	statusEnter = iota
	statusLogin
	statusError
)


func main() {

	// status = 0 - user try login or register
	// status = 1 - user login
	// status = 2 - error
	status := statusEnter
	process := true
	commandName := ""

	// считыванние данных из терминала / команды ?? начет ошибок
	for process {

		if (status == statusEnter) {
			commandName = getCommandName("Enter the command: login | register | forgivePassword | exit ")
			process, status = doCommand(commandName)
		} else if (status == statusLogin) {
			commandName = getCommandName("Enter the command:  putAvatar | getFrendList | getHistory | logout | exit")
			process, status = doCommand(commandName)
		} else {
			consoleLog("Error of input command")
		}
	}
}

// doing commands
 func doCommand (command string) (bool, int) {
	 status := statusEnter
	 switch command {
	 case "login":
		 user, status = processLogin()
		 break

	 case "register":

		 var authdata authData;
		 authdata, status = processRegister()

		 user, _ = auth(authdata)
		 consoleLog(user.name)
		 consoleLog(user.password)
		 break

	 case "logout":
		 user, status = getEmptyUser(), statusEnter
		 break

	 case "fogivePassword":

		 userName := getCommandName("Enter username")
		 user, status = checkDatabase(authData{userName: userName, userPassword: ""})
		 consoleLog(user.password)

		 break

	 case "exit":
		 consoleLog("userId: " + strconv.Itoa(user.id))
		 return false, statusEnter
		 break

	 case "getFriendList":
		 getFriendList(user)
		 break

	 case "putAvatar":
		 putAvatar(user)
		 break

	 case "getHistory":
		 getHistory(user)
		 break

	 default:
		 consoleLog("command not found")
	 }

	 return true, status
 }

func getHistory(user profile) {
    // userID and userFriendID
	// list of history
	return
}

// get list if user friends
func getFriendList(user profile) listOfFriends {

	consoleLog(user.name + "'s list ")
	// list of friends
	return listOfFriends{userID: user.id, friendID: 11}
}

// put avatar
func putAvatar(user profile) string {

	// download the picture to database
	// save URL picture

	user.pictureURL = "URL"
	return user.pictureURL

}

//  Get command from terminal
func getCommandName(msg string) string {

	consoleLog (msg)
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error of input:", err)
	}
	return in.Text()
}


func consoleLog (msg string)  {
	fmt.Println(msg)
}

// user login, enter login and password
func processLogin () (profile, int) {

	userName := getCommandName("Enter user name")
	userPassword := getCommandName("Enter user passoword")
	return checkDatabase(authData{userName: userName, userPassword: userPassword})
}

func auth(authdata authData) (profile, int) {
	// TODO: logAuth
	return checkDatabase(authdata)
}

// check usarname and password in database
func checkDatabase(authdata authData) (profile, int) {
	//connect to database


	if "vasya" == authdata.userName && "qwe" == authdata.userPassword {
		consoleLog("Data is correct")
		return profile{id: 555, name: authdata.userName, password: authdata.userPassword}, statusLogin
	}

	if "vasya" == authdata.userName  && "" == authdata.userPassword {
		consoleLog("Username is correct")
		return profile{id: 555, name: authdata.userName, password: authdata.userPassword}, statusEnter
	}

	return getEmptyUser(), statusError
}

func processRegister() (authData, int) {

	userName := getCommandName("Enter user name")
	userPassword := getCommandName("Enter user passoword")


	if (userName != "" || userPassword != ""){
		consoleLog("Unable to create user, username or password missing")
		processRegister()
	}
	if (userName != "" && userPassword != "") {
		return authData{userName: userName, userPassword: userPassword}, statusEnter
	}
	return authData{userName: userName, userPassword: userPassword}, statusError

}

func getEmptyUser() profile {
	return profile {id: 0, name: "", password: "", pictureURL: ""}
}

