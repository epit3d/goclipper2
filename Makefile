rootdir = $(realpath .)

build-vendor:
	rm -rf $(rootdir)/third_party/clipper2c/build
	mkdir -p $(rootdir)/third_party/clipper2c/build
	cd $(rootdir)/third_party/clipper2c/build && cmake .. -DGO_BINDINGS=ON -DCLIPPER2_TESTS=OFF -DCLIPPER2_EXAMPLES=OFF && sudo make
	cp $(rootdir)/third_party/clipper2c/build/libclipper2c.so lib/libclipper2c.so
	cp $(rootdir)/third_party/clipper2c/build/libclipper2c.so.SOVERSION lib/libclipper2c.so.SOVERSION

generate:
	cd goclipper2 && python3 generator.py
	gofmt -w .

# -a flag forces rebuild of all packages
run:
	go build -a main.go && ./main

build:
	go build -a main.go