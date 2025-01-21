package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	symbolsSource = "symbols.csv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) == 1 || !strings.HasPrefix(os.Args[1], "-") {
		color.New(color.FgHiYellow).Fprintln(os.Stdout, "Usage:")
		fmt.Println("\t main -from <base> -to <target> -amount <amount>")
		return
	}

	base := flag.String("from", "", "the base currency symbol   (1 USD = X TRY, USD is base)")
	target := flag.String("to", "", "the target currency symbol (1 USD = X TRY, TRY is target)")
	amount := flag.Float64("amount", 0, "the amount to be converted")

	flag.Parse()

	//* don't say anything
	color.New(color.FgHiYellow).Fprintln(os.Stdout, "The currency fetching...")

	//* I prefered validating flag variables with a map that i created from "symbosl.csv". There was
	//* a trouble with the API. I mentioned them in `api.go`. Also, i wanted to practice with csv.

	if err := validateFlags(symbolsSource, *base, *target, *amount); err != nil {
		color.New(color.FgRed).Fprintln(os.Stdout, "Error:", err)
		return
	}

	data, err := fetchCurrencyAmount(*base, *target, *amount)
	if err != nil {
		color.New(color.FgRed).Fprintln(os.Stdout, "Error:", err)
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", *amount, *base, data.Result, *target)
	fmt.Printf("%s/%s = %.2f\n", *base, *target, data.Info.Rate)
}
