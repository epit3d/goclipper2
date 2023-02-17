rootdir = $(realpath .)

build-vendor:
	mkdir -p $(rootdir)/vendor/clipper2c/build
	cd $(rootdir)/vendor/clipper2c/build &&	cmake .. -DCLIPPER2_TESTS=OFF -DCLIPPER2_EXAMPLES=OFF && sudo make install
