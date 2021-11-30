package auction

import (
	"fmt"
	"sync"
	"time"
	"utils"
)

type Auction struct {
	highestBid      int
	highestBidderId string
	onGoing         bool
	timeRemaining   int
	log             *utils.Logger
	mu              sync.RWMutex
}

// CheckBid checks if a bid is valid or not. Valid = TRUE Invalid = FALSE
func (a *Auction) CheckBid(bidderId string, bid int) (bool, error) {
	fmt.Printf("Checking bid %v against highest bid %v\n", bid, a.highestBid)
	if !a.onGoing {
		fmt.Println("The auction has ended. Bid rejected.")
		fmt.Printf("The winner was %v with highest bid %v\n", a.highestBidderId, a.highestBid)
		return false, fmt.Errorf("the auction is done")
	}

	if bid <= a.highestBid {
		fmt.Printf("The bid %v, is less than highest bid %v\n", bid, a.highestBid)
		return false, nil
	}

	defer a.mu.Unlock()

	a.mu.Lock()
	a.highestBidderId = bidderId
	a.highestBid = bid
	return true, nil
}

func (a *Auction) Start() {
	a.onGoing = true
	go func() {
		t := 1
		time.Sleep(time.Duration(t) * time.Minute)
		a.onGoing = false
	}()
}

func (a *Auction) GetHighestBid() int {
	return a.highestBid
}

func (a *Auction) GetHighestBidderId() string {
	return a.highestBidderId
}

func (a *Auction) IsOnGoing() bool {
	return a.onGoing
}

func (a *Auction) GetTimeRemaining() int {
	return a.timeRemaining
}

func (a *Auction) SetHighestBidderId(id string) {
	a.highestBidderId = id
}

func (a *Auction) SetHighestBid(amount int) {
	a.highestBid = amount
}

func (a *Auction) SetRemainingTime(timeLeft int) {
	a.timeRemaining = timeLeft
}

func (a *Auction) SetIsOnGoing(onGoing bool) {
	a.onGoing = onGoing
}

func NewAuction() *Auction {
	return &Auction{
		highestBid:      0,
		highestBidderId: "no bidder",
	}
}
