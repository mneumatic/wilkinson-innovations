(() => {
  // Get URL params. Trigger notie flash. pushState to clear params to prevent flash on refresh.
  const queryString = window.location.search
  let urlParams = new URLSearchParams(queryString)
  if (urlParams.get('contact-successful') === "true") {
    notie.alert({ type: 'success', text: 'Success! Your Message has been sent.', time: 3 })
    window.history.replaceState('', document.title, '/')
  }

  if (urlParams.get('mailling-list-successful') === "true") {
    notie.alert({ type: 'success', text: 'Success! Thank you for signing up.', time: 3 })
    window.history.replaceState('', document.title, '/')
  }
})()