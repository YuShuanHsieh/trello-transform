package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	transform "github.com/YuShuanHsieh/trello-transform"
)

func main() {
	var filePath string
	var data []byte
	var err error
	flag.StringVar(&filePath, "file", "", "import local json file")
	list := flag.String("lists", "", "filter list names. Separate each list name by `,`")
	types := flag.String("types", "headers", "display data types. Separate each list name by `,`")
	flag.Parse()

	if filePath != "" {
		if data, err = ioutil.ReadFile(filePath); err != nil {
			log.Fatal(err)
		}
	}
	if data != nil {
		trans := transform.New(data)
		if *list != "" {
			names := strings.Split(*list, ",")
			trans.Select(transform.ByListNames(names...))
		}

		tps := strings.Split(*types, ",")
		for _, tp := range tps {
			switch tp {
			case "headers":
				trans.Use(tp, transform.CardHeaderTransformer)
			case "labels":
				trans.Use(tp, transform.CountLabelsTransformer)
			case "urls":
				trans.Use(tp, transform.MarkdownURLTransformer)
			default:
				log.Printf("The type %s is unknown \n", tp)
			}
		}
		trans.Exec()

		for k, v := range trans.GetAllResult() {
			fmt.Printf("%s: \r\n", strings.ToUpper(k))
			fmt.Println(v)
		}

		fmt.Println("== DONE ==")
	} else {
		fmt.Println("== NO DATA. PLEASE TRY IT AGAIN ==")
	}
}
