# goclipper2

### golang bindings for [Clipper2](https://github.com/AngusJohnson/Clipper2) - a polygon <a href="https://en.wikipedia.org/wiki/Clipping_(computer_graphics)">Clipping</a> and <a href="https://en.wikipedia.org/wiki/Parallel_curve">Offsetting</a> library (originally in C++, C# &amp; Delphi)<br>

### Documentation:

I would gladly refer you to original [documentation](http://www.angusj.com/clipper2/Docs/Overview.htm) by **AngusJohnson** and ask to search for similar names.

### Getting started:

These instructions are for development (was tested on Ubuntu 22.04). Start with cloning this repository.

```bash
git clone --recursive git@github.com:epit3d/goclipper2.git
cd bar
```

I expect you build `clipper2c` which provides C bindings for original library. This should create shared library for `clipper2c` and `clipper2` in `/usr/local/lib` together with their headers.

```bash
mkdir third_party/clipper2c/build
cd third_party/clipper2c/build
cmake .. -DCLIPPER2_TESTS=OFF -DCLIPPER2_EXAMPLES=OFF
sudo make install
```

You may run existing tests in `goclipper2` directory:

```bash
cd goclipper2
go test . -v
```

### Generator notes:

This project would not be possible without bindings of **geoffder** [repository](https://github.com/geoffder/clipper2c)

I decided to use [C AST parser for python](https://github.com/eliben/pycparser) to generate bindings for go because was too lazy to manually type many functions.

Functions have similar names and I split them into categories:

- constructors
- methods
- delete methods
- destruct methods

code in `goclipper2/generator.py` is organized to match C functions in functions `is_<category>` and has custom generator for each of them.

### Contribution

Feel free to create an issue or PR with your thoughts.
