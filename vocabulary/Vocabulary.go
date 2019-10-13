package vocabulary

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	Eng     string
	Rus     string
	Trained int
}

func (w Word) String() string {
	return "[" + w.Eng + "-" + w.Rus + ": " + strconv.Itoa(w.Trained) + "]\n"
}

type Vocabulary []Word

func NewVocabulary() Vocabulary {
	return make(Vocabulary, 0)
}

func ReadVocabulary(name string) (Vocabulary, error) {
	vocBytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	t := strings.Replace(string(vocBytes), "\n", "", -1)
	vocStrings := strings.Split(t, ";")
	vocabulary := NewVocabulary()
	for _, v := range vocStrings {
		word := strings.Split(v, "$")
		vocabulary.Add(Word{Eng: word[0], Rus: word[1]})
	}
	return vocabulary, nil
}

func (v Vocabulary) SaveVocabulary(name string) error {
	var data string
	for _, i := range v {
		data += i.Eng + "$" + i.Rus + ";"
	}

	//remove last ";"
	data = data[:len(data)-1]

	// saving
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	file.Write([]byte(data))
	file.Close()
	return nil
}

func (v *Vocabulary) Add(w Word) {
	*v = append(*v, w)
}

func (v *Vocabulary) Size() int {
	return len(*v)
}

func (v *Vocabulary) Remove(index int) {
	if index == v.Size()-1 {
		*v = (*v)[:index]
		return
	}
	l := (*v)[:index]
	r := (*v)[index+1:]
	*v = append(l, r...)
}
