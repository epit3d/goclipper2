# Developer Notes

For development I suggest building `third_party/clipper2c` locally and cloning the whole repository:

```bash
git clone --recursive git@github.com:epit3d/goclipper2.git
cd goclipper2
```

For windows use MinGW compiler to have comparable results as usual GCC from UNIX systems.

## Building `clipper2c`:

```bash
mkdir third_party/clipper2c/build
cd third_party/clipper2c/build
cmake ..
sudo make install
```

For windows you might receive errors and should replace `cmake ..` with `cmake -G "Unix Makefiles" ..`

You will find built files in `/usr/local/lib` or `C:\Program Files\clipper2c`.
In either way you have to add this path to `LD_LIBRARY_PATH` or `PATH` environment variables.

## Deployment:

If you want to use this `goclipper2` with any of your modules, you have to satisfy presence of `lib` directory with prebuilt binaries (done only for x64). Otherwise your application might build, but crash without linked libraries.

## Generator notes:

This project would not be possible without bindings of **geoffder** [repository](https://github.com/geoffder/clipper2c)

I decided to use [C AST parser for python](https://github.com/eliben/pycparser) to generate bindings for go because was too lazy to manually type many functions.

Functions have similar names and I split them into categories:

- constructors
- methods
- delete methods
- destruct methods

code in `goclipper2/generator.py` is organized to match C functions in functions `is_<category>` and has custom generator for each of them.