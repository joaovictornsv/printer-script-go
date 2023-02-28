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
					Value: 0,
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

		for i := end; i > start; i-- {
			if !elementExist(pagesToIgnore, i) {
				if i%2 == 0 {
					backPages = append(backPages, strconv.Itoa(i))
				} else {
					frontPages = append(frontPages, strconv.Itoa(i))
				}
			}
		}

		backPagesString := strings.Join(backPages, ",")
		frontPagesString := strings.Join(frontPages, ",")

		necessaryPapers := end
		if end%2 != 0 {
			necessaryPapers = end + 1
		}

		fmt.Println("necessary papers:", necessaryPapers)
		fmt.Println("back:", backPagesString)
		fmt.Println("front:", frontPagesString)
	} else {

		pagesInReverseOrder := []string{}
		for i := end; i > start-1; i-- {
			if !elementExist(pagesToIgnore, i) {
				pagesInReverseOrder = append(pagesInReverseOrder, strconv.Itoa(i))
			}
		}

		pagesInReverseOrderString := strings.Join(pagesInReverseOrder, ",")

		fmt.Println("necessary papers:", end)
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
