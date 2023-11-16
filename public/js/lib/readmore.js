(() => {
  const cards = document.querySelectorAll('.feature-card')
  const inputs = document.querySelectorAll('.read-more-state')

  function onChecked (input, card) {
      input.checked
          ? card.classList.add('shadow-lg')
          : card.classList.remove('shadow-lg')
      console.log('checked')
  }

  inputs.forEach((input, index)=> {
      input.addEventListener('change', () => { onChecked(input, cards[index]) })
  })
})()