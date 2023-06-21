const sendData = async (txt) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            categoryName: txt,
        })
    });
};

function addCategory(){
    let form = document.createElement('form');

    let txt = document.createElement('textarea');
    txt.setAttribute("name","postContent")
    form.appendChild(txt);
}