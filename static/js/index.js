/**
 * Created by tlakatlekutl on 13.07.17.
 */
'use strict';

(function main() {
    const frame = document.getElementById('frame');
    const input = document.getElementById('url-input');
    const submit = document.getElementById('submit');

    const headers = new Headers();

    const init = { method: 'POST',
        headers: headers,
        body: '',
        cache: 'no-store' };


    submit.onclick = function () {
        init.body = JSON.stringify({url: input.value});
        fetch('/go/', init).then(function(response) {
            return response.text();
        }).then(function(string) {
            frame.innerHTML = string;
        });
    };

})();
