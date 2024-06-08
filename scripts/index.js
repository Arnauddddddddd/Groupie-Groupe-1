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


function filterCountries() {
    var input, filter, forms, buttons, i;
    input = document.getElementById("searchInput");
    filter = input.value.toUpperCase();
    forms = document.getElementsByClassName("dropdown-content")[0].getElementsByTagName("form");
    for (i = 0; i < forms.length; i++) {
        buttons = forms[i].getElementsByTagName("button")[0];
        if (buttons.innerHTML.toUpperCase().indexOf(filter) > -1) {
            forms[i].style.display = "";
        } else {
            forms[i].style.display = "none";
        }
    }
}