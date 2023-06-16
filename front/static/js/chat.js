function range() {
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
}

function newPost() {
    let form = document.createElement('form');
    form.classList.add('gossip');
    document.body.append(form);

    let txt = document.createElement('textarea');
    txt.setAttribute("name","test")
    form.appendChild(txt);

    let sub = document.createElement("button");
    sub.type = "submit";
    sub.name = "post";
    sub.innerHTML = "Share";
    form.appendChild(sub);

    form.addEventListener("submit", (event) => {
        event.preventDefault();
        let newDiv = document.createElement("div");
        newDiv.innerHTML = txt.value;
        newDiv.classList.add("gossip");
        let form_comment = document.createElement("form");
        let txt_comment = document.createElement('textarea');
        txt_comment.className = 'form_comment';
        form_comment.appendChild(txt_comment);
        let sub_comment = document.createElement("button");
        sub_comment.type = "submit";
        sub_comment.name = "post";
        sub_comment.innerHTML = "Comment";
        form_comment.appendChild(sub_comment);
        newDiv.appendChild(form_comment);
        form.insertAdjacentElement('afterend', newDiv);
        form.reset();

        form_comment.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDiv_comment = document.createElement("div");
            newDiv_comment.innerHTML = txt_comment.value;
            newDiv_comment.classList.add('comment');
            form_comment.insertAdjacentElement('beforebegin', newDiv_comment);
            form_comment.reset();
        });
    });
}

function oldPost() {
    let post = document.getElementsByClassName("gossip");
    for (let i = 0; i < post.length; i++) {
        let form = document.createElement("form");
        let txt_comment = document.createElement('textarea');
        txt_comment.className = 'form_comment';
        form.appendChild(txt_comment);
        let sub_comment = document.createElement("button");
        sub_comment.type = "submit";
        sub_comment.name = "post";
        sub_comment.innerHTML = "Comment";
        form.appendChild(sub_comment);
        post[i].appendChild(form);
        form.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDiv_comment = document.createElement("div");
            newDiv_comment.innerHTML = txt_comment.value;
            newDiv_comment.classList.add('comment');
            form.insertAdjacentElement('beforebegin', newDiv_comment);
            form.reset();
        });
    }
}

export function chat() {
    oldPost();
    range();
    newPost();
}