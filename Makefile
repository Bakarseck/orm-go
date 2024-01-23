stash:
	git stash

pull: stash
	git pull

apply: pull
	git stash apply	

push:
	git add .
	git commit -m "$(message=)"
	git push
