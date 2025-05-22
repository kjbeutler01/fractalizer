package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "net/http"
    "strconv"
)

func hsvToRGB(h, s, v float64) (uint8, uint8, uint8) {
    if s == 0 {
        c := uint8(v * 255)
        return c, c, c
    }
    h = math.Mod(h*6, 6)
    i := int(h)
    f := h - float64(i)
    p := v * (1 - s)
    q := v * (1 - s*f)
    t := v * (1 - s*(1-f))

    var r, g, b float64
    switch i {
    case 0:
        r, g, b = v, t, p
    case 1:
        r, g, b = q, v, p
    case 2:
        r, g, b = p, v, t
    case 3:
        r, g, b = p, q, v
    case 4:
        r, g, b = t, p, v
    default:
        r, g, b = v, p, q
    }
    return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

func mandelbrot(width, height, maxIter int, center complex128, scale float64) *image.RGBA {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for px := 0; px < width; px++ {
        for py := 0; py < height; py++ {
            x := (float64(px)-float64(width)/2)*scale/float64(width) + real(center)
            y := (float64(py)-float64(height)/2)*scale/float64(width) + imag(center)
            c := complex(x, y)
            z := complex(0, 0)
            i := 0
            for cmplx.Abs(z) <= 2 && i < maxIter {
                z = z*z + c
                i++
            }
            col := color.RGBA{0, 0, 0, 255}
            if i < maxIter {
                hue := float64(i) / float64(maxIter)
                r, g, b := hsvToRGB(hue, 0.8, 1.0)
                col = color.RGBA{r, g, b, 255}
            }
            img.Set(px, py, col)
        }
    }
    return img
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
    width := 600
    height := 400
    maxIter := 100
    if n := r.URL.Query().Get("iter"); n != "" {
        if val, err := strconv.Atoi(n); err == nil {
            maxIter = val
        }
    }
    center := complex(0, 0)
    scale := 3.0
    img := mandelbrot(width, height, maxIter, center, scale)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}

func main() {
    http.HandleFunc("/", fractalHandler)
    http.ListenAndServe(":8080", nil)
}

