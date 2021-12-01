package auction

import (
	"fmt"
	"sync"
	"time"
	"utils"
)

// Auction is a local bidding auction on which bids can be published and compared.
// It keeps track of the highest bidder / bid and the remaining time of the auction.
type Auction struct {
	HighestBid      int
	HighestBidderId string
	OnGoing         bool
	TimeRemaining   int
	logger          *utils.Logger
	mu              sync.RWMutex
}

// Start the auction and its timer.
func (a *Auction) Start() {
	a.OnGoing = true
	go func() {
		for a.TimeRemaining != 0 {
			time.Sleep(time.Duration(1) * time.Second)
			a.TimeRemaining--
		}
		a.OnGoing = false
	}()
}

// CheckBid checks if a bid is valid or not.
// If the auction has ended it will also return an error.
func (a *Auction) CheckBid(bidderId string, bid int) (bool, error) {
	a.logger.InfoPrintf("Checking bid %v against highest bid %v", bid, a.HighestBid)
	if !a.OnGoing {
		a.logger.InfoPrintln("The auction has ended. Bid rejected.")
		a.logger.InfoPrintf("The winner was %v with highest bid %v", a.HighestBidderId, a.HighestBid)
		return false, fmt.Errorf("the auction is done")
	}

	if bid <= a.HighestBid {
		a.logger.InfoPrintf("The bid %v, is less than highest bid %v", bid, a.HighestBid)
		return false, nil
	}

	defer a.mu.Unlock()

	a.mu.Lock()
	a.HighestBidderId = bidderId
	a.HighestBid = bid
	return true, nil
}

// NewAuction creates and returns a new Auction with specified duration in seconds.
func NewAuction(duration int, logger *utils.Logger) *Auction {
	return &Auction{
		HighestBid:      0,
		HighestBidderId: "no bidder",
		TimeRemaining:   duration,
		logger:          logger,
	}
}
