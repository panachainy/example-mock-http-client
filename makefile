doc.generate:
	cd doc && ./generate-doc.sh

test:
	go test -tags=test_all -cover ./...
