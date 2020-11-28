package query

// language=sql
const (
	SQLSample     = "SELECT * FROM Users"
	SQLAppVersion = `SELECT MAX(version_id) as version_id FROM goose_db_version WHERE is_applied = TRUE`
)

func GetQueries() []string {
	return []string{
		SQLSample,
		SQLAppVersion,
	}
}
