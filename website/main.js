function disableBtn() {
    document.getElementById("upload-btn").disabled = true;
}

function undisableBtn() {
    document.getElementById("upload-btn").disabled = false;
}

function reset() {
    document.getElementById("upload-btn").disabled = true;
    document.getElementById("filename-input").value = "";
    document.getElementById("file-input").value = "";
}