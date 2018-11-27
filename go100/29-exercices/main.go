package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Logger interface {
	Log(s string)
}

//TODO: créer un type FmtLogger qui implémente Logger
type FmtLogger struct {
	name string
}

func (l *FmtLogger) Log(s string) {
	fmt.Printf("%v %v \n", l.name, s)
}

type Err struct {
	message string
}

func recoverPanic(logger Logger) {
	r := recover() // permet de rattraper une panique. r est de type interface{}
	if r != nil {
		switch s := r.(type) {
		case string:
			logger.Log(s)
		case int:
			logger.Log(strconv.Itoa(s))
		case error:
			logger.Log(s.Error())
		default:
			logger.Log("unknown" + fmt.Sprint(s))
		}
		//TODO: utiliser le logger pour écrire ce que contient la panique (r)
		//Vous devrez gérer le cas où r est un string, où r est une error, et un cas par défaut où on loggera "PANIC: unknown error".
	}
}

func main() {
	//TODO instancier le FmtLogger
	//TODO appeler la fonction recoverPanic (de façon à ce qu'elle soit appelée même en cas de panique) en passant le logger en paramètre
	defer recoverPanic(&FmtLogger{ "main" })

	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(4)
	switch random {
	case 0:
		panic("Oups!")
	case 1:
		panic(errors.New("Zut"))
	case 2:
		panic(Err{"ERR type"})
	default:
		panic(42)
	}
}
