package hwb_to_rgb

import "math"

func HWBToRGB(h, w, bl float64) (r, g, b uint8) {
    h /= 60
    w /= 100
    bl /= 100

    if w + bl >= 1 {
        gray := uint8(w / (w + bl) * 255)
        return gray, gray, gray
    }

    f := h - math.Floor(h)
    v := 1 - bl
    n := w + f*(v-w)
    k := v

    switch int(h) % 6 {
    case 0:
        r, g, b = v, n, w
    case 1:
        r, g, b = n, v, w
    case 2:
        r, g, b = w, v, n
    case 3:
        r, g, b = w, n, v
    case 4:
        r, g, b = n, w, v
    case 5:
        r, g, b = v, w, n
    }

    return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}
