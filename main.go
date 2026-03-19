package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/joho/godotenv"
	statsig "github.com/statsig-io/go-sdk"
)

func main() {
	_ = godotenv.Load()

	secretKey := os.Getenv("STATSIG_SECRET_KEY")
	if secretKey == "" {
		panic("STATSIG_SECRET_KEY must be set in .env or environment")
	}

	statsig.InitializeWithOptions(
		secretKey,
		&statsig.Options{
			Environment: statsig.Environment{Tier: "development"},
		},
	)

	// check 10 random users
	total := 10
	passed := 0
	for i := 0; i < total; i++ {
		userID := rand.Int64()
		if Evaluate("plato_use_pug_for_paypal_billing", int64(userID)) {
			passed++
		}
	}
	pct := float64(passed) / float64(total) * 100
	fmt.Printf("Gate checks passed: %d/%d (%.1f%%)\n", passed, total, pct)
}

func Evaluate(featureName string, regi int64) bool {
	//fmt.Println("Evaluating", featureName, "for user", regi)
	user := statsig.User{CustomIDs: map[string]string{"regi_id": fmt.Sprint(regi)}}
	// If this is a gate, we need to check if the user is in the gate
	val := statsig.CheckGate(user, featureName)
	if val {
		fmt.Println("User is in the gate")
	}
	return val

	// If this is an experiment, we need to get the treatment
	//experiment := statsig.GetExperiment(user, featureName)
	//fmt.Println(experiment)
	//fmt.Println(experiment.GetString("payment_methods", ""))
}
