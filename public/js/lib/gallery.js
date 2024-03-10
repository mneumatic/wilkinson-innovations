(() => {
  const btns = document.querySelectorAll('.gallery-btn')
  const slides = document.querySelectorAll(".slide")
  const dots = document.querySelectorAll(".dot")

  // const testimonials = document.querySelectorAll('.testimonial')
  // const testimonialBtns = document.querySelectorAll('.testimonial-btns button')

  let imgCurrent = 0
  let testimonialCurrent = 0

  // testimonials[testimonialCurrent].classList.remove('hidden')

  slides[imgCurrent].classList.remove('hidden')
  dots[imgCurrent].classList.add('active-dot')

  function move (dir, elements, i, type) {
    i += dir
    if (i < 0) {
        i = elements.length - 1
    }

    if (i === elements.length) {
        i = 0
    }
    elements.forEach(item => {
        item.classList.add('hidden')
    })

    if (type === "slides" || type === 'dot') {
      dots.forEach(item => {
        item.classList.remove('active-dot')
      })
    }

    elements[i].classList.remove('hidden')
    
    switch (type) {
      case "dot":
        dots[i].classList.add('active-dot')
        break
      case "slides":  
        dots[i].classList.add('active-dot')
        imgCurrent = i
        break
      case "testimonial":
        testimonialCurrent = i
        break
    }
  }

  // testimonialBtns[0].addEventListener('click', () => { move(-1, testimonials, testimonialCurrent, "testimonial") })
  // testimonialBtns[1].addEventListener('click', () => { move(1, testimonials, testimonialCurrent, "testimonial") })

  btns[0].addEventListener('click', () => { move(-1, slides, imgCurrent, "slides") }) // prev
  btns[1].addEventListener('click', () => { move(1, slides, imgCurrent, "slides") }) // next

  dots.forEach((dot, index) => { dot.addEventListener("click", () => { move(index, slides, imgCurrent, "dot") })})
})()