package app

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

type printCommandArgs struct {
	Start               int
	End                 int
	PagesToIgnore       string
	isFrontAndBackPrint bool
}

type pagesInfo struct {
	start         int
	end           int
	pagesToIgnore []int
}

type twoSidedPages struct {
	BackPages  []string
	FrontPages []string
}

func calculatePagesOrderToPrint(c *cli.Context) {
	args := getNecessaryArgs(c)

	isFrontAndBackPrint := args.isFrontAndBackPrint
	pagesToIgnore := convertIgnoreArgToArray(args.PagesToIgnore)

	pageDetails := pagesInfo{
		start:         args.Start,
		end:           args.End,
		pagesToIgnore: pagesToIgnore,
	}

	if isFrontAndBackPrint {
		calculatePagesForFrontAndBackPrint(pageDetails)
	} else {
		calculatePagesForOneSidedPrint(pageDetails)
	}
}

func getNecessaryArgs(c *cli.Context) printCommandArgs {
	start := c.Int("start")
	end := c.Int("end")
	ignore := c.String("ignore")
	isFrontAndBackPrint := c.Bool("back")

	necessaryArgs := printCommandArgs{
		Start:               start,
		End:                 end,
		PagesToIgnore:       ignore,
		isFrontAndBackPrint: isFrontAndBackPrint,
	}

	return necessaryArgs
}

func convertIgnoreArgToArray(ignoreArgContent string) []int {
	pagesToIgnore := []int{}

	ignoreArgHasProvided := len(ignoreArgContent) > 0

	if ignoreArgHasProvided {
		for _, page := range strings.Split(ignoreArgContent, ",") {
			pageConvertedToNumber := stringToInt(page)
			pagesToIgnore = append(pagesToIgnore, pageConvertedToNumber)
		}
	}

	return pagesToIgnore
}

func calculatePagesForFrontAndBackPrint(pagesDetails pagesInfo) {
	pagesSequence := getTwoSidedPagesSequence(pagesDetails)
	necessaryPapers := calculatePapersForTwoSidedPrint(pagesSequence)

	printTwoSidedSequence(pagesSequence, necessaryPapers)
}

func calculatePagesForOneSidedPrint(pagesDetails pagesInfo) {
	pagesInReverseOrder := getOneSidedPagesSequence(pagesDetails)
	necessaryPapers := calculatePapersForOneSidedPrint(pagesDetails.start, pagesDetails.end)

	printOneSidedSequence(pagesInReverseOrder, necessaryPapers)
}

func getOneSidedPagesSequence(pagesDetails pagesInfo) []string {
	start := pagesDetails.start
	end := pagesDetails.end
	pagesToIgnore := pagesDetails.pagesToIgnore

	pagesInReverseOrder := []string{}

	for i := end; i >= start; i-- {
		if !elementExistInSlice(pagesToIgnore, i) {
			pagesInReverseOrder = append(pagesInReverseOrder, intToString(i))
		}
	}

	return pagesInReverseOrder
}

func getTwoSidedPagesSequence(pagesDetails pagesInfo) twoSidedPages {
	start := pagesDetails.start
	end := pagesDetails.end
	pagesToIgnore := pagesDetails.pagesToIgnore

	backPages := []string{}
	frontPages := []string{}

	addToBackPages := false

	for currentPage := start; currentPage <= end; currentPage++ {
		isNotIgnoredPage := !elementExistInSlice(pagesToIgnore, currentPage)

		if isNotIgnoredPage {
			pageConvertedToString := intToString(currentPage)
			if addToBackPages {
				backPages = append(backPages, pageConvertedToString)
			} else {
				frontPages = append(frontPages, pageConvertedToString)
			}
			addToBackPages = !addToBackPages
		}
	}

	reversePagesSequenceSlice(backPages)
	reversePagesSequenceSlice(frontPages)

	pagesSequence := twoSidedPages{
		BackPages:  backPages,
		FrontPages: frontPages,
	}

	return pagesSequence
}

func reversePagesSequenceSlice(pagesSlice []string) {
	for i, j := 0, len(pagesSlice)-1; i < j; i, j = i+1, j-1 {
		pagesSlice[i], pagesSlice[j] = pagesSlice[j], pagesSlice[i]
	}
}

func calculatePapersForTwoSidedPrint(pagesSequence twoSidedPages) int {
	backPages := pagesSequence.BackPages
	frontPages := pagesSequence.FrontPages
	totalPagesToPrint := len(backPages) + len(frontPages)

	var necessaryPapers int
	if totalPagesToPrint%2 != 0 {
		necessaryPapers = (totalPagesToPrint-1)/2 + 1
	} else {
		necessaryPapers = totalPagesToPrint / 2
	}

	return necessaryPapers
}

func calculatePapersForOneSidedPrint(start int, end int) int {
	totalNecessaryPapers := (end - start) + 1
	return totalNecessaryPapers
}

func printTwoSidedSequence(pagesSequence twoSidedPages, necessaryPapers int) {
	backPages := pagesSequence.BackPages
	frontPages := pagesSequence.FrontPages

	backPagesString := strings.Join(backPages, ",")
	frontPagesString := strings.Join(frontPages, ",")

	fmt.Println("necessary papers:", necessaryPapers)
	fmt.Println("(first)  back:", backPagesString)
	fmt.Println("(second) front:", frontPagesString)

	if len(backPages) < len(frontPages) {
		fmt.Println("\nObs: Add a blank paper below the back pages print result")
	}
}

func printOneSidedSequence(pageSequence []string, necessaryPapers int) {
	pagesInReverseOrderString := strings.Join(pageSequence, ",")

	fmt.Println("necessary papers:", necessaryPapers)
	fmt.Println("copy and paste:", pagesInReverseOrderString)
}
