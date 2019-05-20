require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
import { HeaderInit } from "./parts/header";
import { FooterInit } from "./parts/footer";
$(() => {
    $(document).ready(function() {
        HeaderInit()
        FooterInit()
    })
});
