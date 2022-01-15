package user

// language=sql
const (
	SQLSample = "SELECT * FROM users"
)

func GetQueries() []string {
	return []string{
		SQLSample,
	}
}
