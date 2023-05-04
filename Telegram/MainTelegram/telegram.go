package maintelegram

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	db "programm/DataBase"
	tok "programm/Telegram/Token"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)
func waitForReply(bot *tgbotapi.BotAPI, chatID int64, userID int) <-chan string {
    ch := make(chan string)

    go func() {
        for {
            updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
                Offset:  0,
                Timeout: 60,
            })
            if err != nil {
                log.Printf("Error getting updates: %s", err.Error())
                continue
            }

            for _, update := range updates {
                if update.Message == nil {
                    continue
                }

                if update.Message.Chat.ID == chatID && update.Message.From.ID == userID {
                    ch <- update.Message.Text
                    return
                }
            }
        }
    }()

    return ch
}

func Telegram() {
	bot, err := tgbotapi.NewBotAPI(tok.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // enable debug logging

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Check if the user is already in the database, and add them if they're not
		
		telegramID := update.Message.From.ID
		user := &db.User{
			TelegramID: telegramID,
		}
		
		_, err := db.GetUser(telegramID)
		if err != nil {
			if err == sql.ErrNoRows {
				err = db.AddUser(user)
				if err != nil {
					log.Printf("Failed to add user %d to database: %v", telegramID, err)
				}
			} else {
				log.Printf("Error checking if user %d is in database: %v", telegramID, err)
			}
		}

		switch update.Message.Text {
		case "/add_balance":
			// Ask the user to enter the amount to add to their balance
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter amount to add to your balance:")
			bot.Send(msg)
		
			// Wait for the user's response
			response := <-waitForReply(bot, update.Message.Chat.ID, update.Message.From.ID)
		
			// Parse the user's response as a float
			amount, err := strconv.ParseFloat(response, 64)
			if err != nil {
				// Handle error
				return
			}
		
			// Add the balance to the database
			err = db.AddUserBalance(update.Message.From.ID, amount)
			if err != nil {
				// Handle error
				return
			}
		
			// Send a confirmation message to the user
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Added %f to your balance.", amount))
			bot.Send(msg)
		

		case "/show_balance":
			// Handle show balance button click
			// Retrieve the balance from the database for the user and display it
			// For this example, let's just display a message with the current balance

			// Get the user's Telegram ID
			telegramID := update.Message.From.ID

			// Retrieve the user's balance from the database
			balance, err := db.GetUserBalance(telegramID)
			if err != nil {
				// Handle error
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error retrieving balance")
				bot.Send(msg)
				continue
			}

			// Display the user's balance
			balanceMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your current balance is "+strconv.FormatFloat(balance, 'f', 2, 64))
			bot.Send(balanceMsg)

		default:
			// Handle unknown command
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
			bot.Send(msg)
		}
	}
}
