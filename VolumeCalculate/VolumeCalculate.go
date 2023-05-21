package volumecalculate

import (
	"fmt"
	"math"
	"time"

	prs "programm/CsvPackges/Parse"
	sig "programm/LogicForSignal"
	ma "programm/Ma/MaFirst"
	rsi "programm/RsiModel"
)

func LogicForCoast() {

	var initialCapital float64
	fmt.Println("Enter the initial capital:")
	fmt.Scan(&initialCapital)
	fmt.Println("========================================================================")

	sharePrice := prs.ParseCSVEndClose()
	sharesOwned := 0.0

	balance := initialCapital
	fmt.Printf("Initial balance: %.2f\n", balance)

	for {
		var action string
		fmt.Println("Enter your action (buy, sell, wait, or self-purchase, auto-sell):")
		fmt.Scan(&action)

		switch action {
		case "buy":
			degreeOfConfidence := sig.SignalWithMemberShipDegreeLong()
			numShares := math.Ceil(balance / sharePrice * math.Abs(degreeOfConfidence))
	
			purchaseAmount := sharePrice * numShares
			if purchaseAmount > balance {
				fmt.Printf("Not enough money to make the purchase. Remaining balance: %.2f\n", balance)
				fmt.Println("========================================================================")
				continue
			}
			balance -= purchaseAmount

			sharesOwned += numShares
			fmt.Printf("Purchased %.2f shares for %.2f. New balance: %.2f\n", numShares, purchaseAmount, balance)
	
			fmt.Printf("You now own %.2f shares.\n", sharesOwned)
			fmt.Println("========================================================================")

		case "sell":
			if sharesOwned < 1 {
				fmt.Println("No shares to sell.")
				continue
			}
			
			degreeOfConfidence := sig.SignalWithMemberShipDegreeShort()
			numShares := math.Ceil(sharesOwned * math.Abs(degreeOfConfidence))
			
			if numShares < 1 {
				fmt.Println("Signal is not strong enough to recommend a sale.")
				fmt.Println("========================================================================")
				continue
			}
			
			saleAmount := sharePrice  * numShares
			fmt.Printf("Recommended number of shares to sell: %.2f\n", numShares)
			fmt.Printf("Estimated sale amount: %.2f\n", saleAmount)
			fmt.Println("========================================================================")
			
			var confirm string
			fmt.Printf("Do you want to sell %.2f shares for %.2f? (y/n): ", numShares, saleAmount)
			fmt.Scan(&confirm)
			
			if confirm != "y" {
				fmt.Println("Sale canceled.")
				continue
			}
			
			sharesOwned -= numShares
			balance += saleAmount
			fmt.Printf("Sold %.2f shares for %.2f. New balance: %.2f. Shares in the portfolio: %.2f\n", numShares, saleAmount, balance, sharesOwned)
			fmt.Println("========================================================================")	


		case "wait":
			fmt.Println("Your balance does not change.")
			fmt.Printf("You own %.2f shares.\n", sharesOwned)
			fmt.Println("========================================================================")
			

		case "self-purchase":
			fmt.Printf("Enter the number of shares to purchase: ")
			var numShares float64
			fmt.Scan(&numShares)

			purchaseAmount := sharePrice * numShares
			if purchaseAmount > balance {
				fmt.Printf("Not enough money to make the purchase. Remaining balance: %.2f\n", balance)
				continue
			}
			balance -= purchaseAmount
			sharesOwned += numShares
			fmt.Printf("Purchased %.2f shares for %.2f. New balance: %.2f\n", numShares, purchaseAmount, balance)
			fmt.Println("========================================================================")

		case "auto-sell":
			saleAmount := sharePrice * sharesOwned
			balance += saleAmount
			fmt.Printf("Sold all %.2f shares for %.2f. New balance: %.2f\n", sharesOwned, saleAmount, balance)
			fmt.Println("========================================================================")
			sharesOwned = 0
		default:
			fmt.Println("Invalid action.")
			fmt.Println("========================================================================")

		}

		if balance == 0 {
			fmt.Println("You ran out of money.")
			fmt.Println("========================================================================")
			break
		}

		time.Sleep(15 * time.Second)
		sharePrice = prs.ParseCSVEndClose()
		ma.ShowMA()
		rsi.ShowRsi()
		sig.SignalWithMemberShipDegree()
		fmt.Println(sig.SignalForBuyOrSail())
		

	}
	fmt.Println("Simulation ended.")
	fmt.Println("========================================================================")
}
