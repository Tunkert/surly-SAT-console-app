package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("   Welcome to the Surly SAT Practice Test!")
	fmt.Println("=============================================")
	fmt.Println("")
	fmt.Println("Are you ready to begin? (press enter for yes)")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	score := 0

	if userInput != "" {
		score += problemOne()
	}

	fmt.Println("Are you ready for the next problem? (press enter for yes)")
	userInput, _ = reader.ReadString('\n')
	if userInput != "" {
		score += problemTwo()
	}

	fmt.Println("Are you ready for the next problem? (press enter for yes)")
	userInput, _ = reader.ReadString('\n')
	if userInput != "" {
		score += problemThree()
	}
}

func insult() string {
	var insultSlice = []string{
		"You fool, that's totally wrong!",
		"You idiot!",
		"Nice try, but no, that's totally wrong, chuckles",
	}

}

func problemOne() int {
	fmt.Println("y = x^2 + 5x - 4")
	fmt.Println("y = x - 4")
	fmt.Println("If the ordered pair (x, y) satisfies the system of equations above, what is one possible value of x?")
	fmt.Println("a) -4")
	fmt.Println("b) 4")
	fmt.Println("c) 1")
	fmt.Println("d) 5")
	fmt.Println("Enter your solution by typing 'a', 'b', 'c', or 'd', Haus. Then press enter.")
	myProblemOneReader := bufio.NewReader(os.Stdin)
	problemOneUserInput, _ := myProblemOneReader.ReadString('\n')
	problemOneUserInput = strings.Replace(problemOneUserInput, "\r\n", "", -1)
	problemOneUserInput = strings.Replace(problemOneUserInput, "\n", "", -1)

	if problemOneUserInput == "a" {
		fmt.Println("Good job, chuckles, you got this one right.")
		return 1
	} else {
		fmt.Println("You fool! You're wrong!!! Hahhaahahahhahahahahahahah ...")
		return 0
	}
}

func problemTwo() int {
	fmt.Printf("\nProblem Two:\n\n")
	fmt.Println("B|\\")
	fmt.Println(" | \\")
	fmt.Println(" |  \\")
	fmt.Println(" |   \\")
	fmt.Println(" |    \\")
	fmt.Println("D|_____\\")
	fmt.Println(" |      \\")
	fmt.Println(" |       \\")
	fmt.Println("A|________\\C")
	fmt.Println("")
	fmt.Println("In the figure above, tan B = 9/40. If BC = 41 and")
	fmt.Println("DA = 4, what is the length of DE?")
	fmt.Println("")
	problemTwoReader := bufio.NewReader(os.Stdin)
	problemTwoUserInput, _ := problemTwoReader.ReadString('\n')
	problemTwoUserInput = strings.Replace(problemTwoUserInput, "\r\n", "", -1)
	problemTwoUserInput = strings.Replace(problemTwoUserInput, "\n", "", -1)

	if problemTwoUserInput == "36" {
		fmt.Println("Nice job, boss.")
		return 1
	} else {
		fmt.Println("Life fail, chuckles.")
		return 0
	}
}

func problemThree() int {
	fmt.Println("")
	fmt.Println("Problem 3")
	fmt.Println("")
	fmt.Println("If r - s = 10 and s/3 = 5, what is r + s. Enter your solution and press enter.")
	problemThreeReader := bufio.NewReader(os.Stdin)
	problemThreeUserInput, _ := problemThreeReader.ReadString('\n')
	problemThreeUserInput = strings.Replace(problemThreeUserInput, "\r\n", "", -1)
	problemThreeUserInput = strings.Replace(problemThreeUserInput, "\n", "", -1)

	if problemThreeUserInput == "40" {
		fmt.Println("Nice job, Haus.")
		return 1
	} else {
		fmt.Println("You clown! That's totally wrong, lol.")
		return 0
	}
}
