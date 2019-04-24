require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
import { HeaderInit } from "./parts/header";
$(() => {
    $(document).ready(function() {
        HeaderInit()
    })
});
