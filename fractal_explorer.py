# Simple interactive fractal explorer using Tkinter

import colorsys
import tkinter as tk


class FractalExplorer:
    def __init__(self):
        self.width = 600
        self.height = 400
        self.max_iter = tk.IntVar(value=100)
        self.fractal_type = tk.StringVar(value="Mandelbrot")
        self.center = 0 + 0j
        self.scale = 3.0
        self.julia_c = -0.4 + 0.6j

        self.root = tk.Tk()
        self.root.title("Fractalizer")
        self.root.configure(bg="#ddf")

        control = tk.Frame(self.root, bg="#ddf")
        control.pack(side=tk.TOP, fill=tk.X)

        tk.Label(control, text="Fractal:", bg="#ddf").pack(side=tk.LEFT, padx=5)
        options = ["Mandelbrot", "Julia", "Burning Ship"]
        tk.OptionMenu(control, self.fractal_type, *options, command=self.redraw).pack(
            side=tk.LEFT, padx=5
        )

        tk.Label(control, text="Iterations:", bg="#ddf").pack(side=tk.LEFT, padx=5)
        tk.Scale(
            control,
            variable=self.max_iter,
            from_=20,
            to=300,
            orient=tk.HORIZONTAL,
            command=lambda e: self.redraw(),
            bg="#ddf",
        ).pack(side=tk.LEFT, padx=5)

        tk.Button(control, text="Reset", command=self.reset, bg="#aaf").pack(
            side=tk.RIGHT, padx=5
        )

        self.canvas = tk.Canvas(self.root, width=self.width, height=self.height)
        self.canvas.pack()
        self.canvas.bind("<Button-1>", self.zoom_in)
        self.canvas.bind("<Button-3>", self.zoom_out)

        self.image = tk.PhotoImage(width=self.width, height=self.height)
        self.canvas.create_image((0, 0), image=self.image, state="normal", anchor=tk.NW)
        self.redraw()

    def run(self):
        self.root.mainloop()

    def reset(self):
        self.center = 0 + 0j
        self.scale = 3.0
        self.redraw()

    def zoom_in(self, event):
        self.center = self.pixel_to_complex(event.x, event.y)
        self.scale /= 2
        self.redraw()

    def zoom_out(self, event):
        self.center = self.pixel_to_complex(event.x, event.y)
        self.scale *= 2
        self.redraw()

    def pixel_to_complex(self, px, py):
        x = (px - self.width / 2) * self.scale / self.width + self.center.real
        y = (py - self.height / 2) * self.scale / self.width + self.center.imag
        return complex(x, y)

    def get_color(self, i):
        if i == self.max_iter.get():
            return "#000000"
        hue = i / self.max_iter.get()
        r, g, b = [int(255 * c) for c in colorsys.hsv_to_rgb(hue, 0.8, 1.0)]
        return f"#{r:02x}{g:02x}{b:02x}"

    def draw_mandelbrot(self):
        for px in range(self.width):
            for py in range(self.height):
                c = self.pixel_to_complex(px, py)
                z = 0 + 0j
                i = 0
                while abs(z) <= 2 and i < self.max_iter.get():
                    z = z * z + c
                    i += 1
                color = self.get_color(i)
                self.image.put(color, (px, py))

    def draw_julia(self):
        for px in range(self.width):
            for py in range(self.height):
                z = self.pixel_to_complex(px, py)
                i = 0
                while abs(z) <= 2 and i < self.max_iter.get():
                    z = z * z + self.julia_c
                    i += 1
                color = self.get_color(i)
                self.image.put(color, (px, py))

    def draw_burning_ship(self):
        for px in range(self.width):
            for py in range(self.height):
                c = self.pixel_to_complex(px, py)
                z = 0 + 0j
                i = 0
                while abs(z) <= 2 and i < self.max_iter.get():
                    z = complex(abs(z.real), abs(z.imag))
                    z = z * z + c
                    i += 1
                color = self.get_color(i)
                self.image.put(color, (px, py))

    def redraw(self, *_):
        fractal = self.fractal_type.get()
        if fractal == "Mandelbrot":
            self.draw_mandelbrot()
        elif fractal == "Julia":
            self.draw_julia()
        elif fractal == "Burning Ship":
            self.draw_burning_ship()
        self.canvas.create_image((0, 0), image=self.image, state="normal", anchor=tk.NW)
        self.canvas.update()


if __name__ == "__main__":
    explorer = FractalExplorer()
    explorer.run()
