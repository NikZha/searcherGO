let objLinks = {
    strrequest: '',
    links: []
}
function loadLinksFromFile(input) {
    const file = input.files[0];
    const reader = new FileReader();

    reader.onload = function (e) {
        const html = e.target.result;

        const parser = new DOMParser();
        const doc = parser.parseFromString(html, 'text/html');

        let links = doc.querySelectorAll('a');
        let reqs = doc.querySelectorAll('input')

        if (links) {
            getHref(links)
        } else {
            alert('В файле не найдена таблица');
        }
        if (reqs) {
            getRequest(reqs)
        } else {
            alert('Не нашёл input')
        }
    };

    reader.onerror = () => alert('Ошибка чтения файла');
    reader.readAsText(file);
}
function getHref(links) {
    let arrayLinks = []
    links.forEach(element => {
        if (element.href && !element.href.includes('yandex')
            && !element.href.includes('google')
            && !element.href.includes('ya.ru')
            && !element.href.includes('127.0.0.1')
            && !element.href.includes('localhost')) {
            arrayLinks.push(element.href)
        }
    });
    objLinks.links = arrayLinks
}
function getRequest(inputs) {
    let strRequest = ''
    inputs.forEach(input => {
        if (input.value) {
            strRequest += input.value
        }
    });
    objLinks.strrequest = strRequest
}
function makeAndSendJSON() {
    let strJson = JSON.stringify(objLinks)
    sendJSONgetResult(strJson)
}
function resetFormAndOblLinksAndSetH3() {
    let reqMessage = document.getElementById('reqMessage')
    if (objLinks.strrequest != '') reqMessage.innerHTML = objLinks.strrequest;
    const fileInput = document.getElementById('inputKagent');
    fileInput.value = '';
    setTimeout(() => {
        objLinks.links = []
        objLinks.strrequest = ''
    }, 1000);
}
async function sendJSONgetResult(json) {
    let resp = await fetch('/postlinks', {
        method: 'Post',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: json
    });
    if (!resp.ok) {
        alert('Ошибка: ' + resp.status);
        return;
    }
    let result = await resp.json();
    makeTableOfResults(result)
}

function makeTableOfResults(arrayResults) {
    let tbody = document.getElementById('resultsOfSearching')
    for (let i = 0; i < arrayResults.length; i++) {
        let tr = document.createElement('tr')
        tbody.append(tr)
        let tdEmails = document.createElement('td')
        let tdUrl = document.createElement('td')
        const element = arrayResults[i];
        makeEmailsLinks(tdEmails, element.emails)
        tdEmails.classList.add('emails')
        tr.append(tdEmails)
        makeUrlLink(tdUrl, element.url)
        tr.append(tdUrl)
    }
}
function makeEmailsLinks(td, emails) {
    if (emails === null) {
        td.append(emails)
    } else {
        td.innerHTML = emails
            .map(email => `<a href="mailto:${email}?subject=Заявка">${email}</a>`)
            .join(', ');
    }
}
function makeUrlLink(td, url) {
    td.innerHTML = `<a href="${url}" target="_blank">${url}</a>`
}