version := 1.0.0

build:
	go build -o ./release/css-botprofiles-$(version).out

clean:
	rm -rf ./release/css-botprofiles-$(version).out
