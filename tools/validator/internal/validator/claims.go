package validator

import "github.com/open-rta/genesis/tools/validator/internal/manifest"

func validateClaims(m manifest.Manifest) []string {
	errors := []string{}
	has := func(k string) bool { _, ok := m.Evidence[k]; return ok }
	claims := map[string]bool{}
	for _, c := range m.Compliance.LawsClaimed {
		claims[c] = true
	}

	if claims["no-rogue-autonomy"] && (!has("objective") || !has("authority")) {
		errors = append(errors, "claim 'no-rogue-autonomy' requires objective and authority evidence")
	}
	if claims["observable-execution"] && !has("trace") {
		errors = append(errors, "claim 'observable-execution' requires trace evidence")
	}
	if claims["interruptible-control"] && !has("control") {
		errors = append(errors, "claim 'interruptible-control' requires control evidence")
	}
	if claims["replayable-behavior"] && !has("replay") {
		errors = append(errors, "claim 'replayable-behavior' requires replay evidence")
	}
	return errors
}
