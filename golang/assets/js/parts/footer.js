var $window = $(window),
  $body = $(document.body)

const KEY = 'WELCOME_BACK',
      UPDATE_TIME = 5 * 60 * 1000 // in ms

const MODAL = function(text) {
    return `<div class="Modal modal" id="modalWelcome" tabindex="-1" role="dialog" aria-labelledby=""
    aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Welcome back</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <p>${text}</p>
          </div>
          <div class="modal-footer">
            <button id="closeModalWelcome" type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
    </div>`
}


function secondsToHms(d) {
    d = Number(d);
    var h = Math.floor(d / 3600);
    var m = Math.floor(d % 3600 / 60);
    var s = Math.floor(d % 3600 % 60);

    var hDisplay = h > 0 ? h + (h == 1 ? " ora si " : "  ore si ") : "";
    var mDisplay = m > 0 ? m + (m == 1 ? " minut si " : " de minute si ") : "";
    var sDisplay = s > 0 ? s + (s == 1 ? " secunda" : " de secunde") : "";
    return hDisplay + mDisplay + sDisplay;
}

function removeModal(){
    $('#modalWelcome').remove()
    $('.modal-backdrop').remove();
}


function set_WelcomeBack(){
    var value =  localStorage.getItem(KEY)
    var done = sessionStorage.getItem(KEY)
    if (!done || done === '') {
        var text = 'Ultima data cand ne-am vazut a fost acum '
        if (value && value !== ''){
            var currentDate = new Date()

            var diffTime = Math.floor(  (currentDate.getTime() - parseInt(value)) / 1000)
            text += secondsToHms(diffTime)

        }else {
             text  += 'mult timp '
        }

        $body.append(MODAL(text))

        $('#modalWelcome').modal('toggle')
        $('#closeModalWelcome').click(removeModal)

        //TO BE COMMENTED IF YOU WANT TO TEST
       sessionStorage.setItem(KEY,'DONE ALREADY')
    }


    const updateTime = function () {
        localStorage.setItem(KEY, new Date().getTime())
        setTimeout(updateTime, UPDATE_TIME)
    }

    updateTime()
}



export function FooterInit() {
  set_WelcomeBack();
 }
