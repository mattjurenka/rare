// +build experimental

package cmd

import (
	"fmt"
	"os"
	"rare/cmd/helpers"
	"rare/pkg/aggregation"
	"rare/pkg/color"
	"rare/pkg/extractor/readState"
	"rare/pkg/multiterm"
	"rare/pkg/multiterm/termrenderers"

	"github.com/urfave/cli"
)

func fuzzyFunction(c *cli.Context) error {
	var (
		topItems    = c.Int("n")
		reverseSort = c.Bool("reverse")
		atLeast     = c.Int64("atleast")
		sortByKey   = c.Bool("sk")
		extra       = c.Bool("extra")
		similarity  = float32(c.Float64("similarity"))
		simOffset   = c.Int("similarity-offset")
		simSize     = c.Int("similarity-size")
		all         = c.Bool("all")
	)

	counter := aggregation.NewFuzzyAggregator(similarity, simOffset, simSize)
	writer := termrenderers.NewHistogram(multiterm.New(), topItems)
	writer.ShowBar = c.Bool("bars") || extra
	writer.ShowPercentage = c.Bool("percentage") || extra

	ext := helpers.BuildExtractorFromArguments(c)

	progressString := func() string {
		return helpers.FWriteExtractorSummary(ext,
			counter.ParseErrors(),
			fmt.Sprintf("(Groups: %s) (Fuzzy: %s)", color.Wrapi(color.BrightBlue, counter.Histo.GroupCount()), color.Wrapi(color.BrightMagenta, counter.FuzzyTableSize())))
	}

	helpers.RunAggregationLoop(ext, counter, func() {
		writeHistoOutput(writer, counter.Histo, topItems, reverseSort, sortByKey, atLeast)
		writer.InnerWriter().WriteForLine(topItems, progressString())
		writer.InnerWriter().WriteForLine(topItems+1, readState.GetReadFileString())
	})

	writer.InnerWriter().Close()

	if all {
		fmt.Println("Full Table:")
		vterm := multiterm.NewVirtualTerm()
		vWriter := termrenderers.NewHistogram(vterm, counter.Histo.GroupCount())
		writeHistoOutput(vWriter, counter.Histo, counter.Histo.GroupCount(), reverseSort, sortByKey, atLeast)

		vterm.WriteToOutput(os.Stdout)
		fmt.Println(progressString())
	}

	return nil
}

func fuzzyCommand() *cli.Command {
	return helpers.AdaptCommandForExtractor(cli.Command{
		Name:      "fuzzy",
		ShortName: "z",
		Aliases:   []string{"fuz"},
		Usage:     "(EXPERIMENTAL) Look for similar matches by using a fuzzy search algorithm",
		Description: `Generates a live-updating histogram of the input data, looking
		for a relative distance between various results.  This is useful to find
		similar log messages that may have slight differences to them (eg ids)
		and aggregating and search for these messages`,
		Action: fuzzyFunction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "After summarization is complete, print all histogram buckets",
			},
			cli.BoolFlag{
				Name:  "bars,b",
				Usage: "Display bars as part of histogram",
			},
			cli.BoolFlag{
				Name:  "percentage",
				Usage: "Display percentage of total next to the value",
			},
			cli.BoolFlag{
				Name:  "extra,x",
				Usage: "Alias for -b --percentage",
			},
			cli.IntFlag{
				Name:  "num,n",
				Usage: "Number of elements to display",
				Value: 5,
			},
			cli.Int64Flag{
				Name:  "atleast",
				Usage: "Only show results if there are at least this many samples",
				Value: 0,
			},
			cli.BoolFlag{
				Name:  "reverse",
				Usage: "Reverses the display sort-order",
			},
			cli.BoolFlag{
				Name:  "sortkey,sk",
				Usage: "Sort by key, rather than value",
			},
			cli.Float64Flag{
				Name:  "similarity,s",
				Usage: "The expression string has to be at least this percent similar to qualify as a fuzzy match",
				Value: 0.75,
			},
			cli.Int64Flag{
				Name:  "similarity-offset,so",
				Usage: "The max offset to examine in the string to look for a similarity",
				Value: 10,
			},
			cli.Int64Flag{
				Name:  "similarity-size,ss",
				Usage: "The maximum size a similarity table can grow to.  Keeps the top most-likely keys at all times",
				Value: 100,
			},
		},
	})
}

func init() {
	commands = append(commands, *fuzzyCommand())
}
