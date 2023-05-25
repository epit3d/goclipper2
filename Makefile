rootdir = $(realpath .)

build-vendor:
	mkdir -p $(rootdir)/third_party/clipper2c/build
	cd $(rootdir)/third_party/clipper2c/build && cmake .. -DGO_BINDINGS=ON -DCLIPPER2_TESTS=OFF -DCLIPPER2_EXAMPLES=OFF && sudo make install

generate:
	cd goclipper2 && python3 generator.py
	gofmt -w .

# -a flag forces rebuild of all packages
run:
	go build -a main.go && ./main