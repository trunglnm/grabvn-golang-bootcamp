package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recoverError() {
	if r := recover(); r != nil {
		fmt.Println("recover from error: ", r)
	}
}

func calFromStr(inputStr string) (result float64) {
	inputArr := strings.Split(inputStr, " ")
	for idx := 0; idx < len(inputArr)-1; idx += 2 {
		if idx == 0 {
			temp, err1 := strconv.ParseFloat(inputArr[idx], 64)
			if err1 != nil {
				errStr := fmt.Sprintf("parameter at [%v] must be numberic", idx)
				panic(errors.New(errStr))
			}
			result = temp
		}
		if idx >= len(inputArr)-2 {
			panic(errors.New("not enough paramaters"))
		}
		para2, err2 := strconv.ParseFloat(inputArr[idx+2], 64)
		if err2 != nil {
			errStr := fmt.Sprintf("parameter at [%v] must be numberic", idx+2)
			panic(errors.New(errStr))
		}

		op := inputArr[idx+1]
		switch op {
		case "+":
			result += para2
		case "-":
			result -= para2
		case "*":
			result *= para2
		case "/":
			if para2 == 0 {
				panic(errors.New("can not divide by zero"))
			}
			result /= para2
		default:
			errStr := fmt.Sprintf("operator at [%v] is not supported", idx+1)
			panic(errors.New(errStr)
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("> ")
	for scanner.Scan() {
		defer recoverError()
		inputStr := scanner.Text()
		output := calFromStr(inputStr)
		fmt.Printf("[%v] = [%v]\n", inputStr, output)
	}
}
