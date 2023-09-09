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

	availableCustomerSuccess := customerSuccess
	if customerSuccessAway != nil {
		availableCustomerSuccess = getAvailableCustomerSuccess(customerSuccess, customerSuccessAway)
	}

	countPerCustomerSuccess := countAssignedCustomers(customers, availableCustomerSuccess)

	busiestCustomerSuccessID := getBusiestCustomerSuccess(countPerCustomerSuccess)
	return busiestCustomerSuccessID
}

// sortEntitiesByScore sorts a slice of entities by their scores in ascending order.
func sortEntitiesByScore(entities []Entity) {
	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Score < entities[j].Score
	})
}

// getAvailableCustomerSuccess filters out customer success who are away.
func getAvailableCustomerSuccess(customerSuccess []Entity, customerSuccessAway []int) []Entity {
	var availableCustomerSuccess []Entity

	for _, cs := range customerSuccess {
		isAway := false
		for _, csAway := range customerSuccessAway {
			if csAway == cs.ID {
				isAway = true
				break
			}
		}
		if !isAway {
			availableCustomerSuccess = append(availableCustomerSuccess, cs)
		}
	}

	return availableCustomerSuccess
}

// countAssignedCustomers counts the number of customers assigned to each available customer success.
func countAssignedCustomers(customers []Entity, availableCustomerSuccess []Entity) map[int]int {
	countPerCustomerSuccess := make(map[int]int)

	for _, c := range customers {
		for _, cs := range availableCustomerSuccess {
			if c.Score <= cs.Score {
				countPerCustomerSuccess[cs.ID]++
				break
			}
		}
	}

	return countPerCustomerSuccess
}

// getBusiestCustomerSuccess gets the busiest customer success.
// If there is a tie for the busiest representative, it will return 0.
func getBusiestCustomerSuccess(countPerCustomerSuccess map[int]int) int {
	var busiestCustomerSuccessScore, busiestCustomerSuccessID int
	var isTie bool

	for id, score := range countPerCustomerSuccess {
		if score > busiestCustomerSuccessScore {
			busiestCustomerSuccessScore = score
			busiestCustomerSuccessID = id
			isTie = false
		} else if score == busiestCustomerSuccessScore {
			isTie = true
		}
	}

	if isTie {
		return 0
	}

	return busiestCustomerSuccessID
}
