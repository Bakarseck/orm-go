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

start:
	cd $(app) && go run .	
