package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "print",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start",
					Value: 1,
				},
				cli.IntFlag{
					Name:     "end",
					Required: true,
				},
				cli.StringFlag{
					Name:  "ignore",
					Value: "",
				},
				cli.BoolFlag{
					Name: "back",
				},
			},
			Action: calculatePagesOrderToPrint,
		},
	}

	return app
}

func calculatePagesOrderToPrint(c *cli.Context) {
	start := c.Int("start")
	end := c.Int("end")
	ignore := c.String("ignore")

	pagesToIgnore := []int{}
	if len(ignore) > 0 {
		for _, page := range strings.Split(ignore, ",") {
			pageConvertedToNumber := stringToInt(page)
			pagesToIgnore = append(pagesToIgnore, pageConvertedToNumber)
		}
	}

	isFrontAndBackPrint := c.Bool("back")
	if isFrontAndBackPrint {
		backPages := []string{}
		frontPages := []string{}

		addToBackPages := false
		for i := end; i >= start; i-- {
			if !elementExist(pagesToIgnore, i) {
				if addToBackPages {
					backPages = append(backPages, strconv.Itoa(i))
				} else {
					frontPages = append(frontPages, strconv.Itoa(i))
				}
				addToBackPages = !addToBackPages
			}
		}

		backPagesString := strings.Join(backPages, ",")
		frontPagesString := strings.Join(frontPages, ",")

		totalPagesToPrint := len(backPages) + len(frontPages)
		var necessaryPapers int
		if totalPagesToPrint%2 != 0 {
			necessaryPapers = (totalPagesToPrint-1)/2 + 1
		} else {
			necessaryPapers = totalPagesToPrint / 2
		}

		fmt.Println("necessary papers:", necessaryPapers)
		fmt.Println("(first)  back:", backPagesString)
		fmt.Println("(second) front:", frontPagesString)
		if len(backPages) < len(frontPages) {
			fmt.Println("\nObs: Add a blank paper below the back pages print result")
		}
	} else {

		pagesInReverseOrder := []string{}
		for i := end; i >= start; i-- {
			if !elementExist(pagesToIgnore, i) {
				pagesInReverseOrder = append(pagesInReverseOrder, strconv.Itoa(i))
			}
		}

		pagesInReverseOrderString := strings.Join(pagesInReverseOrder, ",")

		fmt.Println("necessary papers:", (end-start)+1)
		fmt.Println("copy and paste:", pagesInReverseOrderString)
	}

}

func stringToInt(s string) int {
	stringConverted, e := strconv.Atoi(s)

	if e != nil {
		log.Fatalf("Error converting %s to int", s)
	}

	return stringConverted
}

func elementExist(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
