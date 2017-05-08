var toggleSpeed = 1000;

$(".hidectrl-commitratio").click(function(){
  $(".report-commitratio").toggleClass("hidden", toggleSpeed);
  $(this).toggleClass("fa-plus-square fa-minus-square");
})

$(".hidectrl-commits").click(function(){
  $(".report-commits").toggleClass("hidden", toggleSpeed);
  $(this).toggleClass("fa-plus-square fa-minus-square");
})

$(".hidectrl-mergerequests").click(function(){
  $(".report-mergerequests").toggleClass("hidden", toggleSpeed);
  $(this).toggleClass("fa-plus-square fa-minus-square");
})

$('.hidectrl-li').click(function() {
  $(this).parent().parent().parent().find(".report-li-div").toggleClass("hidden");
  $(this).toggleClass("fa-plus-square fa-minus-square");
})

$('.btn-printview').click(function() {
  var sidebar = $('.sidebar-container');

  if(sidebar.is(":visible")) {
    $('.sidebar-container').hide();
    $('.content-container').toggleClass("col-xl-9 col-lg-9 col-md-9 col-sm-9");
    $('.content-container').toggleClass("col-xl-12 col-lg-12 col-md-12 col-sm-12");
  } else {
    $('.sidebar-container').show();
    $('.content-container').toggleClass("col-xl-9 col-lg-9 col-md-9 col-sm-9");
    $('.content-container').toggleClass("col-xl-12 col-lg-12 col-md-12 col-sm-12");
  }
})

$(document).ready(function(){
    $('[data-spy="scroll"]').each(function(){
        var $spy = $(this).scrollspy('refresh');
    });

    $("#myNavbar").on("activate.bs.scrollspy", function(){
        var currentItem = $(".nav li.active > a").text();

    });
});
