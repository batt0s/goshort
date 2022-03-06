function showCustomBox() {
    var checkbox = document.getElementById("customcheck");
    var customboxdiv = document.getElementById("customdiv");
    var shortenbutton = document.getElementById("shortenbtn");
    if (checkbox.checked == true) {
        customboxdiv.style.display = "block";
        shortenbutton.style.display = "none";

    } else {
        customboxdiv.style.display = "none";
        shortenbutton.style.display = "initial";
    }
}

function shortenUrl() {
    let urlToShort = document.querySelector("#url");
    let shortResult = document.querySelector("#short");

    let xhr = new XMLHttpRequest();
    let url = "shorten";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            shortResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            shortResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":urlToShort.value});

    xhr.send(data);

}

function customShorten() {
    let urlToShort = document.querySelector("#url");
    let customShortUrl = document.querySelector("#customurl");
    let shortResult = document.querySelector("#short");

    let xhr = new XMLHttpRequest();
    let url = "shorten/custom";
    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            shortResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            shortResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":urlToShort.value, "custom":customShortUrl.value});

    xhr.send(data);
}

function getOrigin() {
    let shortUrl = document.querySelector("#shorturl");
    let originResult = document.querySelector("#origin");

    let xhr = new XMLHttpRequest();
    let url = "getOrigin";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            originResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            originResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":shortUrl.value});

    xhr.send(data);
}