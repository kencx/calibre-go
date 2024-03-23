
clean:
	rm -rf models

regenerate: metadata.db
	sqlboiler --add-global-variants --no-tests -c sqlboiler.toml sqlite3
