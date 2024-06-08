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

function filterByYear() {
    var minYear = document.getElementById("minYear").value;
    var maxYear = document.getElementById("maxYear").value;

    if (!minYear || !maxYear) {
        alert("Veuillez sélectionner les deux années.");
        return;
    }

    var url = new URL(window.location.href);
    url.searchParams.set("minYear", minYear);
    url.searchParams.set("maxYear", maxYear);

    window.location.href = url.toString();
}

function filterFirstAlbum() {
    var minYear = document.querySelector('[name="minFirstAlbumYear"]').value;
    var maxYear = document.querySelector('[name="maxFirstAlbumYear"]').value;

    var url = new URL(window.location.href);
    url.searchParams.set("minFirstAlbumYear", minYear);
    url.searchParams.set("maxFirstAlbumYear", maxYear);

    window.location.href = url.toString();
}

function filterNumberMembers() {
    var minMembers = document.querySelector('[name="minNumberMembers"]').value;
    var maxMembers = document.querySelector('[name="maxNumberMembers"]').value;

    var url = new URL(window.location.href);
    url.searchParams.set("minNumberMembers", minMembers);
    url.searchParams.set("maxNumberMembers", maxMembers);

    window.location.href = url.toString();
}
