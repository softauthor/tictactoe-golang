/*
  Game: Tic Tac Toe
  Author: Rajarajan Thangarasu
*/

package main

import (
  "fmt"
  "os"
  "math/rand"
  "time"
  "strconv"
)


// ============================== Global Vars ==================================

var ttt2D  = [3][3]string{
                          {"1", "2", "3"},
                          {"4", "5", "6"},
                          {"7", "8", "9"},
                          }
var symbol = "🍎"
var gameChance = 0
var gameInputsArray = []int{1,2,3,4,5,6,7,8,9}


//=============== Draw 2D Matrix on the system console ========================+

func drawGameBoard() {

  fmt.Println("---------")

   for i := 0; i < 3; i++ {
       for j := 0; j < 3; j++ {
          fmt.Print(" ")
          fmt.Print( ttt2D[i][j])
          fmt.Print(" ")
       } //for
        fmt.Println("\n---------")
   } //  for

}

//===================== setting and get PlayerName =============================

func getPlayerName() string {

    var playerName string;

    if(symbol == "🍎"){
      playerName ="Player A"
    } else {
      playerName ="Player B"
    }
    return playerName
}


//=======Takes Random Numer and Sets an appropriate symbol on the matrix =======

func userInput(randomValue int) {

    fmt.Println("\n\n"+ symbol +" "+ getPlayerName() +" Turn. (" + strconv.Itoa(gameChance) + ")"  )
    fmt.Println("Random Number = " + strconv.Itoa(randomValue));

     switch randomValue {
      case 1:
       ttt2D[0][0] = symbol
      case 2:
       ttt2D[0][1] = symbol
      case 3:
      ttt2D[0][2] = symbol
      case 4:
       ttt2D[1][0] = symbol
      case 5:
        ttt2D[1][1] = symbol
      case 6:
       ttt2D[1][2] = symbol
      case 7:
        ttt2D[2][0] = symbol
      case 8:
         ttt2D[2][1] = symbol
      case 9:
        ttt2D[2][2] = symbol
      default:

    } //switch
}


//==================Toggle symbol for Player A and Player B ===================
func toggleSymbol() {
  if(symbol == "🍎"){
    symbol = "🍇"
  } else {
    symbol = "🍎"
  }
}

//======================= Check row, colum and diagonal ========================

func winCheck() string {

  // I know this could be done in a loop but I did it this way for time constraint
  if(ttt2D[0][0] == "🍎" && ttt2D[0][1] == "🍎" && ttt2D[0][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[1][0] == "🍎" && ttt2D[1][1] == "🍎" && ttt2D[1][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[2][0] == "🍎" && ttt2D[2][1] == "🍎" && ttt2D[2][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[0][0] == "🍎" && ttt2D[1][0] == "🍎" && ttt2D[2][0] == "🍎") {
    return "🍎"
  } else if(ttt2D[0][1] == "🍎" && ttt2D[1][1] == "🍎" && ttt2D[2][1] == "🍎") {
    return "🍎"
  } else if(ttt2D[0][2] == "🍎" && ttt2D[1][2] == "🍎" && ttt2D[2][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[0][0] == "🍇" && ttt2D[0][1] == "🍇" && ttt2D[0][2] == "🍇") {
    return "🍇"
  } else if(ttt2D[1][0] == "🍇" && ttt2D[1][1] == "🍇" && ttt2D[1][2] == "🍇") {
    return "🍇"
  } else if(ttt2D[2][0] == "🍇" && ttt2D[2][1] == "🍇" && ttt2D[2][2] == "🍇") {
    return "🍇"
  } else if(ttt2D[0][0] == "🍇" && ttt2D[1][0] == "🍇" && ttt2D[2][0] == "🍇") {
    return "🍇"
  } else if(ttt2D[0][1] == "🍇" && ttt2D[1][1] == "🍇" && ttt2D[2][1] == "🍇") {
    return "🍇"
  } else if(ttt2D[0][2] == "🍇" && ttt2D[1][2] == "🍇" && ttt2D[2][2] == "🍇") {
    return "🍇"
  } else if(ttt2D[0][0] == "🍎" && ttt2D[1][1] == "🍎" && ttt2D[2][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[2][0] == "🍎" && ttt2D[1][1] == "🍎" && ttt2D[0][2] == "🍎") {
    return "🍎"
  } else if(ttt2D[0][0] == "🍇" && ttt2D[1][1] == "🍇" && ttt2D[2][2] == "🍇") {
    return "🍇"
  } else if(ttt2D[2][0] == "🍇" && ttt2D[1][1] == "🍇" && ttt2D[0][2] == "🍇") {
    return "🍇"
  }

return "@"

}
//win


//============================= Display Results ================================

func displayResutls() {
  if(winCheck()== "🍎") {

    fmt.Println("\n🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆");
    fmt.Println("🍎 Player A Wins 😁");
    fmt.Println("🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆\n");
    os.Exit(3)

  } else if (winCheck() == "🍇") {
    fmt.Println("\n🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆");
    fmt.Println("🍇 PlayerB Wins 😁");
    fmt.Println("🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆🏆\n");
    os.Exit(3)

  } else if (winCheck() == "@" && gameChance == 9) {
    fmt.Println("It's Draw 🍎🍇🤗!");
    os.Exit(3)
  }
}



//============================= removeRandomNumberFromAnArray ==================

func removeRandomNumberFromAnArray(getRandomNumber int) {
  i := 0

  for i < len(gameInputsArray) {
    if(gameInputsArray[i] == getRandomNumber) {
      gameInputsArray = append(gameInputsArray[:i], gameInputsArray[i+1:]...)
      break;
    }
    i++
  }
}


//=================================== Game Play ================================
func gamePlay() {

  // Timer for every 2 seconds
  ticker := time.NewTicker(time.Millisecond * 2000)

  go func() {
      for _= range ticker.C {

        rand.Seed(time.Now().UTC().UnixNano())

        //  get the random number from a gameInputsArray
        var getRandomNumber = gameInputsArray[rand.Intn(len(gameInputsArray))];

        // tracking the count of the timer
        gameChance++

         userInput(getRandomNumber)

         drawGameBoard()

         displayResutls();

         toggleSymbol()

         removeRandomNumberFromAnArray(getRandomNumber)

        } //for

  }()

  time.Sleep(time.Millisecond * 20000)
  ticker.Stop()
}


//=================================== APP BEGINS ===============================
func main(){

   fmt.Println("\n****************************************")
   fmt.Println("*********** Tic Tac Toe Game ***********")
   fmt.Println("****************************************\n")

    drawGameBoard()
    gamePlay()
}
