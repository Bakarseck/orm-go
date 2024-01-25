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

start: rm-migrates
	@cd $(app) && go run .

rm-migrates:
	cd test_orm && rm -rf migrates		
