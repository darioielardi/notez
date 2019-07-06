
.PHONY: test

dev:
	export GOOGLE_APPLICATION_CREDENTIALS="/Users/dario/Projects/experiments/notez/config/firebase-dev-service-account.json" && reflex -c ./reflex.conf

test:
	export GOOGLE_APPLICATION_CREDENTIALS="/Users/dario/Projects/experiments/notez/config/firebase-dev-service-account.json" && go test ./test -count=1