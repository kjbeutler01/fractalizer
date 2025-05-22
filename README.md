# Fractalizer

Fractalizer is a simple and colorful fractal explorer built with Python's Tkinter module. It allows you to play with a few classic fractals and zoom into them with a couple of mouse clicks.

## Features

- Explore Mandelbrot, Julia and Burning Ship fractals
- Adjustable iteration depth with a slider
- Click to zoom in, right click to zoom out
- Reset button to return to the default view

Run the explorer with:

```bash
python3 fractal_explorer.py
```

Have fun navigating the bubbly interface and discovering new patterns!

## Go & Pulumi version

A minimal HTTP server that renders the Mandelbrot set is available in `go-fractalizer`. You can start it with:

```bash
go run ./go-fractalizer
```

Infrastructure to containerize and run the server using Pulumi is defined in `infra`.
A sample stack configuration `Pulumi.dev.yaml` sets the port used by the container.
You can deploy the stack with:

```bash
cd infra
pulumi up
```

The port can be customized by editing the stack file or via `pulumi config set port <num>`.
