(function () {
    var scenarios = document.querySelectorAll('.scenario');
    var hint = document.getElementById('scenario-hint');
    scenarios.forEach(function (el) {
        el.addEventListener('click', function () {
            scenarios.forEach(function (s) { s.classList.remove('active'); });
            el.classList.add('active');
            hint.textContent = el.getAttribute('data-desc');
        });
    });
})();
