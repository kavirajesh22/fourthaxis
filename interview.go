package main
import "fmt"

//tic tac toe game practice

func drawBoard(array [3][3] string) {
	for i :=0; i<3; i++{
		for j :=0; j<3; j++{
			fmt.Print(i)
			fmt.Print(j, " ")
			fmt.Print(array[i][j], " ")
		}
		fmt.Println()
	}
}

func checkWinsRow(array [3][3] string) bool{
	var check bool = false
	var wins int = 0
	for i:=0; i<3; i++{
		if array[i][0]=="X"{
			for j:=0; j<3; j++{
				if array[i][j]=="X"{
					check = true
				}else{
					check = false
					wins++
				}
			}
		}else{
			for j:=0; j<3; j++{
				if array[i][j]=="O"{
					check = true
				}else{
					check = false
					wins++
				}
			}
		}
		if (wins ==0 && check==true){
			return check
		}
		wins = 0
	}
	check = false
	return check
}

func checkWinsCol(array [3][3] string) bool{
	var check bool = false
	var wins int = 0
	for j:=0; j<3; j++{
		if array[0][j]=="X"{
			for i:=0; i<3; i++{
				if array[i][j]=="X"{
					check = true
				}else{
					check = false
					wins++
				}
			}
		}else{
			for i:=0; i<3; i++{
				if array[i][j]=="O"{
					check = true
				}else{
					check = false
					wins++
				}
			}
		}
		if (wins ==0 && check == true){
			return check
		}
		wins = 0
	}
	check = false
	return check
}

func playGame(array [3][3] string) [3][3]string{
	//var win bool = false
	//var count int = 0
	var coin string = "X"

	for i:=0; i<3; i++{
		for j:=0; j<3; j++{
			if array[i][j]=="*"{
				array[i][j]=coin
				return array
			}
		}
		if i==2{
			return array
		}
	}
	return array
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")
	var board= [3][3]string{ {"*","X", "X"},{"X","O", "X"},{"O","X", "X"}}
	
	/*
	drawBoard(board)
	if checkWinsRow(board){
		fmt.Println("There's a win in the row.")
	}
	if checkWinsCol(board){
		fmt.Println("There's a win in the column.")
	}*/
	drawBoard(playGame(board))

	
}
