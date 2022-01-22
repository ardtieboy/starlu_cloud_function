async function dropHandler(ev) {
    console.log('File(s) dropped');

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();

    console.log(ev.dataTransfer)

    if (ev.dataTransfer.items) {
        // Use DataTransferItemList interface to access the file(s)
        for (var i = 0; i < ev.dataTransfer.items.length; i++) {
            // If dropped items aren't files, reject them
            if (ev.dataTransfer.items[i].kind === 'file') {
                var file = ev.dataTransfer.items[i].getAsFile();
                console.log(await file.arrayBuffer())
                console.log('... file[' + i + '].name = ' + file.name);
            }
        }
    } else {
        // Use DataTransfer interface to access the file(s)
        for (var i = 0; i < ev.dataTransfer.files.length; i++) {
            console.log('... file[' + i + '].name = ' + ev.dataTransfer.files[i].name);
        }
    }
}

function dragOverHandler(ev) {
    console.log("OVER HOVER")

    document.getElementById("drop_zone").style.border = "thick dotted #0000FF";

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
}

function dragLeave(ev) {
    console.log("LEAVING")
    document.getElementById("drop_zone").style.border = "thick solid #0000FF";

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
}

function clicker() {
    console.log('CLICK');
}

function myFunction() {
    console.log("BLAAAAAT");
}