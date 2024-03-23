
clean:
	rm -rf internal/models

regenerate: metadata.db
	sqlboiler --add-global-variants --no-tests -c sqlboiler.toml --output internal/models sqlite3
