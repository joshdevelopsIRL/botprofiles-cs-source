version := 1.0.0

build:
	go build -o ./css-botprofiles-$(version).out

clean:
	rm -rf ./css-botprofiles-$(version).out
