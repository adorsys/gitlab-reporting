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

$(document).ready(function(){
    $('[data-spy="scroll"]').each(function(){
        var $spy = $(this).scrollspy('refresh');
    });

    $("#myNavbar").on("activate.bs.scrollspy", function(){
        var currentItem = $(".nav li.active > a").text();
        
    });
});
