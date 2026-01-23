package main

import (
	"fmt"
	"math/rand/v2"

	statsig "github.com/statsig-io/go-sdk"
)

func main() {
	statsig.InitializeWithOptions(

		"supply-secret-here",
		&statsig.Options{
			Environment: statsig.Environment{Tier: "production"},
		},
	)

	countActive := 0
	countInactive := 0
	// check 10 random users
	for i := 0; i < 100; i++ {
		userID := rand.Int64()
		stat := IsActive("payflow_apple_pay_via_pug", userID)
		fmt.Printf("User %d: %t\n", userID, stat)
		if stat {
			countActive++
		} else {
			countInactive++
		}
	}
	fmt.Printf("Count Active: %d, Count Inactive: %d\n", countActive, countInactive)
}

func IsActive(featureName string, regi int64) bool {
	user := statsig.User{CustomIDs: map[string]string{"regi_id": fmt.Sprint(regi)}}
	return statsig.CheckGate(user, featureName)
	experiment := statsig.GetExperiment(user, featureName)
	return experiment.GetBool(featureName, true)
}
