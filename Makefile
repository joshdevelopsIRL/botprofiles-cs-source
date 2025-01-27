version := 1.0.1

build:
	go build -o ./css-botprofiles-$(version).out

run:
	./css-botprofiles-$(version).out

clean:
	rm -rf ./css-botprofiles-$(version).out
