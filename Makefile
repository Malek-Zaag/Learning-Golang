init:
	git init
	git remote add origin https://github.com/Malek-Zaag/golang-simple-microservice.git
git:
	git add . 
	git commit -m "$(msg)"
	git push -u origin main

run:
	go build -o $(build)
	./bin/$(build)