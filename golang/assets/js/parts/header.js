var $window = $(window),
  $header = $(".Header"),
  $header_btn = $(".Header__account a.btn"),
  $menu_toggle = $(".Header__toggle"),
  header_is_set = false,
  header_is_sticky = false,
  scrollTop = 0;

// Header initial set & scroll event
function set_scrollHeader() {
  // get scroll amount
  scrollTop = $window.scrollTop();

  // make Header white
  if (scrollTop > 0) {
    if (!header_is_sticky) {
      $header.addClass("is_sticky");
      header_is_sticky = true;
    }

    $header_btn.removeClass("btn-outline-light").addClass("btn-outline-primary");

    // make Header transparent
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

  // add header animations
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
}
