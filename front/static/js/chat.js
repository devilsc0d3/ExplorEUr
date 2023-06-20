function range() {
    let content = document.createElement('div');
    content.classList.add('ranges')
    document.body.append(content)

    let range1 = document.createElement('input');
    range1.setAttribute('type','range')
    range1.setAttribute('min','200')
    range1.setAttribute('max','800')
    range1.id ='width'
    range1.addEventListener("input", () => {
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
    range2.addEventListener("input", () => {
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
    range3.addEventListener("input", () => {
        let el = document.querySelectorAll('.gossip')
        for (let i = 0 ; i< el.length; i++) {
            el[i].style.background = 'hsl(280,50%,'+ range3.value + '%)'
        }
    })
    content.appendChild(range3)
}

const sendDataPost = async (txt) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            postContent: txt,
        })
    });
};

const sendDataComment = async (txt, postId) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            comment: txt,
            postID: postId,
        })
    });
};

function newPost() {
    let form = document.createElement('form');
    form.classList.add('gossip');
    document.body.insertAdjacentElement('afterbegin', form);

    let txt = document.createElement('textarea');
    txt.setAttribute("name","postContent")
    form.appendChild(txt);

    let sub = document.createElement("button");
    sub.type = "submit";
    sub.name = "post";
    sub.innerHTML = "Share";
    form.appendChild(sub);

    form.addEventListener("submit", (event) => {
        event.preventDefault();
        let newDiv = document.createElement("div");

        const txtContent = txt.value;
        const insults = ["fuck", "fuck off", "mother fucker", "bitch", "bastar", "bastard", "suck", "suck my dick", "fuck you", "fuck your mother", "shit"];
        const regex = new RegExp("\\b(" + insults.join("|") + ")\\b", "gi");

        const filteredText = txtContent.replace(regex, function (match) {
            return '*'.repeat(match.length);
        });

        newDiv.innerHTML = filteredText;



        newDiv.classList.add("gossip");
        let formComment = document.createElement("form");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name","comment")
        txtComment.className = 'form_comment';
        formComment.appendChild(txtComment);
        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        formComment.appendChild(subComment);
        newDiv.appendChild(formComment);
        form.insertAdjacentElement('afterend', newDiv);
        sendDataPost(txt.value).then(res => res.json());
        form.reset();

        formComment.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment = document.createElement("div");
            newDivComment.innerHTML = txtComment.value;
            newDivComment.classList.add('comment');
            const postId = formComment.parentNode.dataset.id;
            formComment.insertAdjacentElement('beforebegin', newDivComment);
            sendDataComment(txtComment.value, postId).then(res => res.json());
            formComment.reset();
        });
    });
}

function oldPost() {
    let posts = document.getElementsByClassName("gossip");
    for (let i = 0; i < posts.length; i++) {
        let form = document.createElement("form");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name","comment")
        txtComment.className = 'form_comment';
        form.appendChild(txtComment);
        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        form.appendChild(subComment);
        posts[i].appendChild(form);
        form.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment = document.createElement("div");
            newDivComment.innerHTML = txtComment.value;
            newDivComment.classList.add('comment');
            const postId = form.parentNode.dataset.id;
            form.insertAdjacentElement('beforebegin', newDivComment);

            sendDataComment(txtComment.value, postId).then(res => res.json());

            form.reset();
        });
    }
}

export function chat() {
    oldPost();
    range();
    newPost();
}