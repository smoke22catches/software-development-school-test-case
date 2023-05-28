package emails

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"software-development-school-test-case/price"
)

const emailSubscriptionsFileName = "emails.txt"

func AddEmailToSubscriptionList(email string) (emailAdded bool, err error) {
	subscriptions, err := getOrCreateSubscriptionsList()
	if err != nil {
		return false, err
	}

	// check if email already exists in subscriptions list
	if isEmailInSubscriptionsList(email, subscriptions) {
		return false, nil
	}

	subscriptions = append(subscriptions, email)

	// before saving subscriptions must be sorted
	sortSubscriptionsList(subscriptions)
	err = saveSubscriptionsList(subscriptions)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SendBtcPriceToSubscribers() error {
	btcPrice, err := price.GetBtcPriceInUah()

	if err != nil {
		return err
	}

	subscriptions, err := getOrCreateSubscriptionsList()

	if err != nil {
		return err
	}

	for _, email := range subscriptions {
		fmt.Printf("Send email to %s with price %f\n", email, btcPrice)
	}

	return nil
}

// get subscriptions list from file, if it doesn't exist then create it
func getOrCreateSubscriptionsList() ([]string, error) {
	fileExists, err := subscriptionFileExists()
	if err != nil {
		return nil, err
	}
	if fileExists {
		return getSubscriptionsList()
	}

	err = createSubscriptionsFile()
	if err != nil {
		return nil, err
	}
	return []string{}, nil
}

func subscriptionFileExists() (bool, error) {
	if _, err := os.Stat(emailSubscriptionsFileName); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func getSubscriptionsList() ([]string, error) {
	fileContent, err := os.ReadFile(emailSubscriptionsFileName)
	if err != nil {
		return nil, err
	}
	var subscriptions []string = strings.Split(string(fileContent), "\n")
	return subscriptions, nil
}

func createSubscriptionsFile() error {
	_, err := os.Create(emailSubscriptionsFileName)
	return err
}

// subscriptions must be sorted
func isEmailInSubscriptionsList(email string, subscriptions []string) bool {
	if len(subscriptions) == 0 {
		return false
	}

	i := sort.SearchStrings(subscriptions, email)

	if i >= len(subscriptions) {
		return false
	}

	return subscriptions[i] == email
}

func sortSubscriptionsList(subscriptions []string) {
	sort.Strings(subscriptions)
}

func saveSubscriptionsList(subscriptions []string) error {
	fileContent := strings.Join(subscriptions, "\n")
	return os.WriteFile(emailSubscriptionsFileName, []byte(fileContent), 0644)
}
