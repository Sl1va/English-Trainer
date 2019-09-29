package main

import (
	"encoding/json"
	"fmt"
	. "github.com/Sl1va/English-Trainer/vocabulary"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	vocabulary Vocabulary // unlearned words
	regular    Vocabulary // all word
	learned    Vocabulary // learned words
)

var (
	need = 5
	add  = 1
	sub  = 2
)

func training() {
	run := true
	for run {
		// choose random unlearned word index
		uRand := rand.Intn(len(vocabulary))

		// choose random position in question
		randPos := rand.Intn(4)

		// 0 - eng-rus, 1 - rus-eng
		mode := rand.Intn(2)

		if mode == 0 {
			fmt.Println(vocabulary[uRand].Eng)

			question := make([]string, 4)
			question[randPos] = vocabulary[uRand].Rus

			set := NewSet()
			set.Add(vocabulary[uRand].Rus)

			// fill others
			for i := 0; i < 4; i++ {
				if question[i] != "" {
					continue
				}

				// get random value from all words
				randWord := rand.Intn(len(regular))

				// if we already had it
				for set.Contains(regular[randWord].Rus) {
					randWord = rand.Intn(len(regular))
				}

				set.Add(regular[randWord].Rus)

				question[i] = regular[randWord].Rus
			}

			//write questions
			for i := 0; i < 4; i++ {
				fmt.Println(strconv.Itoa(i) + ")" + question[i])
			}

			var command string
			fmt.Scan(&command)

			if checkCommand(command) {
				continue
			}

			answer, _ := strconv.Atoi(command)

			if answer == randPos {
				fmt.Println("[+]")
				vocabulary[uRand].Trained += add
				if vocabulary[uRand].Trained >= need {
					learned.Add(vocabulary[uRand])
					vocabulary.Remove(uRand)
				}
			} else {
				fmt.Println("[-] " + vocabulary[uRand].Rus)
				vocabulary[uRand].Trained -= sub
			}

		} else {
			fmt.Println(vocabulary[uRand].Rus)

			question := make([]string, 4)
			question[randPos] = vocabulary[uRand].Eng

			set := NewSet()
			set.Add(vocabulary[uRand].Eng)

			// fill others
			for i := 0; i < 4; i++ {
				if question[i] != "" {
					continue
				}

				// get random value from all words
				randWord := rand.Intn(len(regular))

				// if we already had it
				for set.Contains(regular[randWord].Eng) {
					randWord = rand.Intn(len(regular))
				}

				set.Add(regular[randWord].Eng)

				question[i] = regular[randWord].Eng
			}

			//write questions
			for i := 0; i < 4; i++ {
				fmt.Println(strconv.Itoa(i) + ")" + question[i])
			}

			var command string
			fmt.Scan(&command)

			if checkCommand(command) {
				continue
			}

			answer, _ := strconv.Atoi(command)

			if answer == randPos {
				fmt.Println("[+]")
				vocabulary[uRand].Trained += add
				if vocabulary[uRand].Trained >= need {
					learned.Add(vocabulary[uRand])
					vocabulary.Remove(uRand)
				}
			} else {
				fmt.Println("[-] " + vocabulary[uRand].Eng)
				vocabulary[uRand].Trained -= sub
			}
		}
		if vocabulary.Size() == 0 {
			fmt.Println("Training finished!")
			run = false
		}
	}
}

func addNewWord() {
	var eng, rus string
	fmt.Println("english: ")
	fmt.Scan(&eng)
	fmt.Println("russian: ")
	fmt.Scan(&rus)
	word := Word{Eng: eng, Rus: rus}
	vocabulary.Add(word)
	regular.Add(word)
}

func checkCommand(com string) bool {
	_, err := strconv.Atoi(com)
	isDigit := err != nil
	if com == "stat" {
		fmt.Println(vocabulary)
	} else if com == "learned" {
		fmt.Println(learned)
	} else if com == "add" {
		addNewWord()
	} else if com == "conf" {
		fmt.Printf("need=%d\nadd=%d\nsub=%d\n", need, add, sub)
	} else if com == "save" {
		fmt.Println("filename: ")
		var name string
		fmt.Scan(&name)
		err := regular.SaveVocabulary(name)
		if err != nil {
			fmt.Println("Couldn't save vocabulary")
		} else {
			fmt.Println("Saved!")
		}
	}
	return isDigit
}

type Conf struct {
	Need int `json:"need"`
	Add  int `json:"add"`
	Sub  int `json:"sub"`
}

func readConf() {
	file, _ := os.Open("conf.txt")
	data, _ := ioutil.ReadAll(file)
	var conf Conf
	json.Unmarshal(data, &conf)
	need = conf.Need
	add = conf.Add
	sub = conf.Sub
}

func main() {
	readConf()

	// read vocabulary and regular
	vocabulary, _ = ReadVocabulary("vocabulary.txt")
	regular = make(Vocabulary, len(vocabulary))
	copy(regular, vocabulary)

	learned = NewVocabulary()

	// init random
	rand.Seed(time.Now().UnixNano())

	//start training
	training()
}
