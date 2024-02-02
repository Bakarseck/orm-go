stash:
	git stash

pull: stash
	git pull

apply: pull
	git stash apply	

push: pull apply
	git add .
	git commit -m "$(message)"
	git push

rm-migrates:
	cd test_orm && rm -rf migrates

rm-db:
	cd test_orm && rm -rf test.db

start:
	cd $(app) && go run .	

		
