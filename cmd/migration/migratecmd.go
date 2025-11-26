package migration
import (
	"metabeat/internal/db"
)
func MigrateDatabase(database *db.DB) {
	database.Migrate()
}
