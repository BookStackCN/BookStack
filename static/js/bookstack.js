$(function () {
    //回滚到顶部
    $(".bookstack-gotop").click(function () {
        $('html,body').animate({scrollTop:0}, 200);
    });
    //TODO:根据summary.md，生成上一篇和下一篇的链接地址

    //以https或者http开头的url链接，加上target='_blank'
    $(document).on("click","a",function (e) {
        e.preventDefault();
        var href=$(this).attr("href").toLowerCase();
        if($(this).attr("target")=="_blank" ||href.indexOf("http://")==0 || href.indexOf("https://")==0 ){
            window.open(href);
        }else{
            location.href=href;
        }
    });


    $("#show-summary").click(function () {
        var summary=$("#bookstack-render-summary");
        summary.hasClass("show")?summary.removeClass("show"):summary.addClass("show");
    });

});