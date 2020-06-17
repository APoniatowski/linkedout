/* Toggle between adding and removing the "responsive" class to topnav when the user clicks on the icon */
window.onscroll = function () {
    stickyNavBar()
};

function menuButtonFunction() {
    var burger = document.getElementById("navBar");
    if (burger.className === "topnav") {
        burger.className += " responsive";
    } else {
        burger.className = "topnav";
    }
}

function stickyNavBar() {
var navbar = document.getElementById("navBar");
var sticky = navbar.offsetTop;
    if (window.pageYOffset >= sticky) {
        navbar.classList.add("sticky")
    } else {
        navbar.classList.remove("sticky");
    }
}
