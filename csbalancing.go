package csbalancing

import (
	"sort"
)

// Entity represents a generic entity for customers and customer success
type Entity struct {
	ID    int
	Score int
}

// CustomerSuccessBalancing finds the busiest customer success.
// If there is a tie for the busiest representative, it will return 0.
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	sortEntitiesByScore(customerSuccess)
	sortEntitiesByScore(customers)

	availableCS := customerSuccess
	if customerSuccessAway != nil {
		availableCS = getAvailableCustomerSuccess(customerSuccess, customerSuccessAway)
	}

	countCustomer := countCustomer(customers, availableCS)

	busiestCSID := getBusiestCustomerSuccess(countCustomer)
	return busiestCSID
}

// sortEntitiesByScore sorts a slice of entities by their scores in ascending order.
func sortEntitiesByScore(entities []Entity) {
	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Score < entities[j].Score
	})
}

// getAvailableCustomerSuccess filters out customer success who are away.
func getAvailableCustomerSuccess(customerSuccess []Entity, customerSuccessAway []int) []Entity {
	var availableCS []Entity

	for _, cs := range customerSuccess {
		isAway := false
		for _, csAway := range customerSuccessAway {
			if csAway == cs.ID {
				isAway = true
				break
			}
		}
		if !isAway {
			availableCS = append(availableCS, cs)
		}
	}

	return availableCS
}

// countCustomer counts the number of customers assigned to each available customer success.
func countCustomer(customers []Entity, availableCS []Entity) map[int]int {
	countCustomer := make(map[int]int)

	for _, c := range customers {
		for _, cs := range availableCS {
			if c.Score <= cs.Score {
				countCustomer[cs.ID]++
				break
			}
		}
	}

	return countCustomer
}

// getBusiestCustomerSuccess gets the busiest customer success.
// If there is a tie for the busiest representative, it will return 0.
func getBusiestCustomerSuccess(countCustomer map[int]int) int {
	var busiestCS, busiestCSID int
	var isTie bool

	for k, v := range countCustomer {
		if v > busiestCS {
			busiestCS = v
			busiestCSID = k
			isTie = false
		} else if v == busiestCS {
			isTie = true
		}
	}

	if isTie {
		return 0
	}

	return busiestCSID
}
