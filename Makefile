all: revive
revive:
	revive -config ./revive.toml -formatter friendly -exclude ./vendor/... ./...
release:
	./semver.sh -p
release-major:
	./semver.sh -M
release-minor:
	./semver.sh -m
