(define real-min -2 .5)
(define real-max +1 .5)
(define imag-min -2 .0)
(define imag-max +2 .0)

(define max-color 255)
(define iter-count 40)
(define escape-radius 4 .0)

(define width 1024)
(define scale (/ width (- real-max real-min)))
(define height (inexact->exact (round (* scale (- imag-max imag-min)))))

(define (mandelbrot x y)
  (define (*mandelbrot z0 z rem)
    (if (or (= rem 0) (>= (magnitude z) escape-radius))
      (inexact->exact (round (* (/ rem iter-count) max-color)))
      (*mandelbrot z0 (+ (* z z) z0) (- rem 1))))
  (let ((val (make-rectangular (+ (/ x scale) real-min)
               (+ (/ y scale) imag-min))))
    (*mandelbrot val 0 iter-count)))

(define (plot file)
  (call-with-output-file file
   (λ (fp)
    (display "P2" fp) (newline fp)
    (display width fp) (newline fp)
    (display height fp) (newline fp)
    (display max-color fp) (newline fp)
    (do ((y 0 (+ y 1))) ((>= y height))
      (do ((x 0 (+ x 1))) ((>= x width))
        (display (mandelbrot x y) fp)
        (display #\space fp))
      (newline fp)))))

(plot "mandelbrot.pgm")
