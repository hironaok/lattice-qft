const mainSlides =
    Array.from(document.querySelectorAll(".main-slide"));

const backupSlides =
    Array.from(document.querySelectorAll(".backup-slide"));

const pageNumber =
    document.getElementById("page-number");

const modeButton =
    document.getElementById("mode-button");


let mode = "main";

let mainIndex = 0;
let backupIndex = 0;


function getSlides() {
    return mode === "main"
        ? mainSlides
        : backupSlides;
}


function getIndex() {
    return mode === "main"
        ? mainIndex
        : backupIndex;
}


function setIndex(index) {
    if (mode === "main") {
        mainIndex = index;
    } else {
        backupIndex = index;
    }
}


function showSlide() {
    const slides = getSlides();

    document
        .querySelectorAll(".slide")
        .forEach(slide =>
            slide.classList.remove("active")
        );

    let index = getIndex();

    index =
        (index + slides.length) % slides.length;

    setIndex(index);

    slides[index].classList.add("active");

    pageNumber.textContent =
        `${index + 1} / ${slides.length}`;

    history.replaceState(
        null,
        "",
        `#${mode}-${index + 1}`
    );
}


function nextSlide() {
    setIndex(getIndex() + 1);
    showSlide();
}


function previousSlide() {
    setIndex(getIndex() - 1);
    showSlide();
}


function toggleMode() {
    if (mode === "main") {
        mode = "backup";
        modeButton.textContent = "backup";
    } else {
        mode = "main";
        modeButton.textContent = "main";
    }

    showSlide();
}


/* Keyboard */

document.addEventListener("keydown", event => {

    if (
        event.key === "ArrowRight" ||
        event.key === " "
    ) {
        event.preventDefault();
        nextSlide();
    }

    if (event.key === "ArrowLeft") {
        event.preventDefault();
        previousSlide();
    }
});


/* Restore slide after F5 */

function restoreSlide() {
    const match =
        location.hash.match(
            /^#(main|backup)-(\d+)$/
        );

    if (!match) {
        return;
    }

    mode = match[1];

    const index = Number(match[2]) - 1;

    if (mode === "main") {
        mainIndex = index;
        modeButton.textContent = "main";
    } else {
        backupIndex = index;
        modeButton.textContent = "backup";
    }
}


restoreSlide();
showSlide();

