package body

// language=sql
const (
	SQLSample     = "SELECT * FROM users"
	SQLAppVersion = `SELECT MAX(version_id) as version_id FROM goose_db_version WHERE is_applied = TRUE`
)

func GetQueries() []string {
	return []string{
		SQLSample,
		SQLAppVersion,
	}
}
