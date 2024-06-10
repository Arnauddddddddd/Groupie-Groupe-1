var prevScrollpos = window.pageYOffset;
window.onscroll = function() {
    var currentScrollPos = window.pageYOffset;
    if (prevScrollpos > currentScrollPos) {
        document.getElementById("navigation").style.top = "0";
    } else {
        document.getElementById("navigation").style.top = "-50px";
    }
    prevScrollpos = currentScrollPos;
}

// Fonction pour afficher les boutons de pays
function filterCountries() {
    let input = document.getElementById('searchInput');
    let filter = input.value.toUpperCase();

    let form = document.getElementById('countriesForm');
    let buttons = form.getElementsByClassName('country-button');

    for (let i = 0; i < buttons.length; i++) {
        let txtValue = buttons[i].textContent || buttons[i].innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
            buttons[i].style.display = "";
        } else {
            buttons[i].style.display = "none";
        }
    }
}
