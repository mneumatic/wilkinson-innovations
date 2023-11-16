(() => {
  const navbar = document.getElementById('site-navigation-navbar')
  const toggle = document.getElementById('site-navigation-navbar-toggle')
  const links = document.querySelectorAll('#site-navigation-navbar-nav-links a')
  const sections = getSections()
  let last

  function getSections () {
    const tmp = document.querySelectorAll('section')
    const arr = []

    tmp.forEach(item => {
        arr.push(item)
    })

    return arr //.slice(1)
  }

  function activeOnScroll () {
    let current = ''
    const OFFSET = 80

    sections.forEach((section) => {
        const sectionTop = section.offsetTop
        if (scrollY >= (sectionTop - OFFSET)) current = section.getAttribute('id')
    })

    links.forEach((li) => {
        li.classList.remove('active')
        if (li.classList.contains(current)) li.classList.add('active')
    })
  }

  function closeMenuOnScroll () {
    const top = window.scrollY

    toggle.checked === true && top < last
      ? toggle.checked = true
      : toggle.checked = false

    last = top
  }

  function closeMenuOnResize () {
    if (toggle.checked === true && window.innerWidth >= 768) toggle.checked = false
  }


  function onScroll() {
    window.scrollY > 0
      ? navbar.classList.add('shadow-lg')
      : navbar.classList.remove('shadow-lg')
  }

  window.addEventListener('scroll', onScroll)
  window.addEventListener('scroll', activeOnScroll)
  window.addEventListener('scroll', closeMenuOnScroll)
  window.addEventListener('resize', closeMenuOnResize)
})()
