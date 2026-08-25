const slides = document.querySelectorAll(".slide");
const pageNumber = document.getElementById("page-number");

let currentSlide = Number(location.hash.slice(1)) || 0;

function showSlide(index) {
    slides[currentSlide].classList.remove("active");

    currentSlide =
        (index + slides.length) % slides.length;

    slides[currentSlide].classList.add("active");

    location.hash = currentSlide;

    updatePageNumber();
}

function updatePageNumber() {
    pageNumber.textContent =
        `${currentSlide + 1} / ${slides.length}`;
}

document.addEventListener("keydown", function(event) {

    if (
        event.key === "ArrowRight" ||
        event.key === " " ||
        event.key === "Enter"
    ) {
        showSlide(currentSlide + 1);
    }

    if (event.key === "ArrowLeft") {
        showSlide(currentSlide - 1);
    }

    if (event.key === "Home") {
        showSlide(0);
    }

    if (event.key === "End") {
        showSlide(slides.length - 1);
    }
});

/* Show the saved slide when the page loads */
slides.forEach(slide => slide.classList.remove("active"));
slides[currentSlide].classList.add("active");

updatePageNumber();
