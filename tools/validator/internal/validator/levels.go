package validator

func validateLevel(m Manifest) (Level, []string, []string) {
	errs := []string{}
	notes := []string{}
	has := func(k string) bool { _, ok := m.Evidence[k]; return ok }
	hasAttestation := m.Attestation.By != "" && m.Attestation.Timestamp != "" && m.Attestation.Statement != ""

	validated := LevelL0
	if hasAttestation && has("objective") && has("authority") && has("trace") {
		validated = LevelL1
	}
	if validated == LevelL1 && has("control") {
		validated = LevelL2
	}
	if validated == LevelL2 && has("replay") {
		validated = LevelL3
	}

	claimed := m.Compliance.LevelClaimed
	if claimed == LevelL4 {
		notes = append(notes, "L4 requires manual/foundation review and cannot be granted automatically")
		return validated, errs, notes
	}

	rank := map[Level]int{LevelL0: 0, LevelL1: 1, LevelL2: 2, LevelL3: 3}
	if rank[validated] < rank[claimed] {
		errs = append(errs, "claimed level is not supported by provided prerequisite evidence")
	}
	return validated, errs, notes
}
