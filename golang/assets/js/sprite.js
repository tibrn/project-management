
var $head = $(document.head)
function randomCSS(start){
    start = start * 200
    return ``

}


function SpriteInit(){
    var $style = $('#sprite-css')
    if($style.length === 0){
     $("<style>")
    .prop("type", "text/css")
    .html(randomCSS(0))
    .appendTo('head')
    }
}
$(() => {
    $(document).ready(function() {

        // SpriteInit()
    })
});
