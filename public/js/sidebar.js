$('#nav').affix({
  offset: {
    bottom: ($('footer').outerHeight(true) + $('.application').outerHeight(true)) + 40
  }
})

$('#nav').affix({
    offset: {
        top: $('#nav').offset().top
    }
})

$('body').scrollspy({
       target: '.thingy'
})
