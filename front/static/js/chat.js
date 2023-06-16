export function grid() {
    let content = document.createElement('div');
    content.classList.add('ranges')
    document.body.append(content)

    let range1 = document.createElement('input');
    range1.setAttribute('type','range')
    range1.setAttribute('min','200')
    range1.setAttribute('max','800')
    range1.id ='width'
    range1.addEventListener("input", (event) => {
        let el = document.querySelectorAll('.gossip')
        for (let i = 0 ; i< el.length; i++) {
            el[i].style.width = range1.value + 'px'
        }
    })
    content.appendChild(range1)


    let range2 = document.createElement('input');
    range2.setAttribute('type','range')
    range2.setAttribute('min','20')
    range2.setAttribute('max','40')
    range2.id ='fontSize'
    range2.addEventListener("input", (event) => {
        let el = document.querySelectorAll('.gossip')
        for (let i = 0 ; i< el.length; i++) {
            el[i].style.fontSize = range2.value + 'px'
        }
    })
    content.appendChild(range2)


    let range3 = document.createElement('input');
    range3.id ='background'
    range3.setAttribute('type','range')
    range3.setAttribute('min','20')
    range3.setAttribute('max','75')
    range3.addEventListener("input", (event) => {
        let el = document.querySelectorAll('.gossip')
        for (let i = 0 ; i< el.length; i++) {
            el[i].style.background = 'hsl(280,50%,'+ range3.value + '%)'
        }
    })
    content.appendChild(range3)

    let form = document.createElement('form');
    form.classList.add('gossip');
    document.body.append(form);

    let txt = document.createElement('textarea');
    txt.name = "test"
    form.appendChild(txt);

    let sub = document.createElement("button");
    sub.type = "submit";
    sub.name = "post";
    sub.innerHTML = "Share gossip!";
    form.appendChild(sub);



    form.addEventListener("submit", (event) => {
        event.preventDefault();
        let newDiv = document.createElement("div");
        newDiv.innerHTML = txt.value;
        newDiv.classList.add("gossip");
        form.insertAdjacentElement('afterend', newDiv);
        form.reset();
    });
}