package levels

import "github.com/open-rta/genesis/tools/validator/internal/manifest"

func ValidateClaimedLevel(m manifest.Manifest) (manifest.Level, []string, []string) {
	errs := []string{}
	notes := []string{}
	has := func(k string) bool { _, ok := m.Evidence[k]; return ok }
	hasAttestation := m.Attestation.By != "" && m.Attestation.Timestamp != "" && m.Attestation.Statement != ""

	validated := manifest.LevelL0
	if hasAttestation && has("objective") && has("authority") && has("trace") {
		validated = manifest.LevelL1
	}
	if validated == manifest.LevelL1 && has("control") {
		validated = manifest.LevelL2
	}
	if validated == manifest.LevelL2 && has("replay") {
		validated = manifest.LevelL3
	}

	if m.Compliance.LevelClaimed == manifest.LevelL4 {
		notes = append(notes, "L4 requires manual/foundation review and cannot be granted automatically")
		return validated, errs, notes
	}

	rank := map[manifest.Level]int{manifest.LevelL0: 0, manifest.LevelL1: 1, manifest.LevelL2: 2, manifest.LevelL3: 3}
	if rank[validated] < rank[m.Compliance.LevelClaimed] {
		errs = append(errs, "claimed level is not supported by provided prerequisite evidence")
	}
	return validated, errs, notes
}
