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

## Installing on Windows

1. Install [Python 3 for Windows](https://www.python.org/downloads/windows/) and make sure to check **Add Python to PATH** during installation.
2. Install [Git for Windows](https://git-scm.com/download/win) using the default options.
3. Open **Git Bash** or **Command Prompt** and clone this repository:

   ```bash
   git clone <repository-url>
   cd fractalizer
   ```

   You can also download the repository as a ZIP file and extract it.
4. Start the explorer:

   ```bash
   python fractal_explorer.py
   ```

### Optional tools

- To run the Go server, install [Go](https://go.dev/dl/) and execute `go run ./go-fractalizer`.
- To try the Pulumi deployment, install [Pulumi](https://www.pulumi.com/docs/get-started/install/) and Docker Desktop, then run the commands in the `infra` directory.

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
