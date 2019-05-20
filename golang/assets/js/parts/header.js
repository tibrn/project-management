var $window = $(window),
  $header = $(".Header"),
  $header_btn = $(".Header__account a.btn"),
  $menu_toggle = $(".Header__toggle"),
  header_is_set = false,
  header_is_sticky = false,
  scrollTop = 0,
  svg_close = `<svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px"
  width="50" height="50"
  viewBox="0 0 192 192"
  style=" fill:#ffffff;"><g fill="none" fill-rule="nonzero" stroke="none" stroke-width="1" stroke-linecap="butt" stroke-linejoin="miter" stroke-miterlimit="10" stroke-dasharray="" stroke-dashoffset="0" font-family="none" font-weight="none" font-size="none" text-anchor="none" style="mix-blend-mode: normal"><path d="M0,192v-192h192v192z" fill="none"></path><g id="original-icon" fill="#ffffff"><g id="surface1"><path d="M58.08,47.16l-10.92,10.92l38.16,37.92l-38.16,37.92l10.92,10.92l38.16,-37.92l38.16,37.92l10.8,-10.92l-38.04,-37.92l38.04,-37.92l-10.8,-10.92l-38.16,37.92z"></path></g></g></g></svg>`,
  svg_open = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 30" width="30" height="30" focusable="false"><title>Menu</title><path stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-miterlimit="10" d="M4 7h22M4 15h22M4 23h22"></path></svg>`;

function set_scrollHeader() {

  scrollTop = $window.scrollTop();

  if (scrollTop > 0) {
    if (!header_is_sticky) {
      $header.addClass("is_sticky");
      header_is_sticky = true;
    }

    $header_btn.removeClass("btn-outline-light").addClass("btn-outline-primary");

  } else if (scrollTop === 0) {
    if (header_is_sticky) {
      $header.removeClass("is_sticky");
      header_is_sticky = false;
    }

    if ($window.width() > 991) {
      $header_btn.addClass("btn-outline-light").removeClass("btn-outline-primary");
    } else {
      $header_btn.removeClass("btn-outline-light").addClass("btn-outline-primary");
    }
  }

  if (header_is_set) return false;

  setTimeout(function() {
    $header.addClass("is_set");
    header_is_set = true;
  }, 20);
}

function set__resizeHeader() {
  if ($window.width() > 991 && $header.hasClass("is_open")) {
    $header.removeClass("is_open");
    $menu_toggle.find("i").toggleClass("icon-close icon-menu");
  }

  if ($window.width() > 991) {
    $header_btn.addClass("btn-outline-light").removeClass("btn-outline-primary");
  } else {
    $header_btn.removeClass("btn-outline-light").addClass("btn-outline-primary");
  }
}

export function HeaderInit() {
  set_scrollHeader();

  $window.on("scroll", function() {
    set_scrollHeader();
  });

  $window.on("resize", function() {
    set__resizeHeader();
  });
  var open = false
  $menu_toggle.on("click", function() {
    // set mobile login button color
    if ($window.width() < 700 && $header_btn.hasClass("btn-outline-light")) {
      $header_btn.toggleClass("btn-outline-light btn-outline-primary");
    }

    // set header open class
    $header.toggleClass("is_open");

    // change menu toggle icon
    $menu_toggle.find("button")[0].innerHTML = open ? svg_open : svg_close

    open = !open
  });

}
