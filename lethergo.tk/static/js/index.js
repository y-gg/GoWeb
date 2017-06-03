$(function () {
    $("#play").click(function () {
        var api = $("#api").select().val();
        var url = $("#url").val().trim();
        $("#player").attr("src",api + url);
    });
});