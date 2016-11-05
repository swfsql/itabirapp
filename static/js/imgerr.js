$(document).ready(function()
{
    $(".err_user").on("error", function(){
        $(this).attr('src', '/static/images/user/err.png');
    });
});