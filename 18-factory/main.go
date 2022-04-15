package main

import "fmt"

// interface
type iLang interface {
	getName() string
	setName(name string)
	compile()
	runtime()
}

// abstract class
type lang struct {
	name string
}

func (l *lang) setName(name string) { l.name = name }
func (l *lang) getName() string     { return l.name }
func (l *lang) compile()            { fmt.Printf("compiling by %s compiler", l.getName()) }
func (l *lang) runtime()            {}

// concrete class golang
type golang struct {
	lang
}

func newGolang() iLang {
	return &golang{
		lang: lang{name: "golang"},
	}
}

// concrete class cpp
type cpp struct {
	lang
}

func newCpp() iLang {
	return &cpp{
		lang: lang{name: "cpp"},
	}
}

// concrete class java
type java struct {
	lang
}

func newJava() iLang {
	return &java{
		lang: lang{name: "java"},
	}
}

// factory design pattern
func getLang(langType string) (iLang, error) {
	if langType == "golang" {
		return newGolang(), nil
	} else if langType == "cpp" {
		return newCpp(), nil
	} else if langType == "java" {
		return newJava(), nil
	}

	return nil, fmt.Errorf("please specify correct lang type")
}

func main() {
	langIns, err := getLang("cpp")
	if err != nil {
		panic("can get lang")
	}
	langIns.compile()
}
