
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
    let form = document.querySelector('[name="newPost"]')
    console.log(form)
    form.classList.add('posts');
    const header = document.querySelector("header")
    header.insertAdjacentElement('afterend', form);

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
        newDiv.innerHTML = txt.value;
        newDiv.classList.add("posts");
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
        sendDataPost(txt.value).then(res => res.json()).catch(res => Promise.fail({error:res}));
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
    let posts = document.getElementsByClassName("posts");

    for (let i = 0; i < posts.length; i++) {
        // let form = document.createElement("form");
        let form = posts[i].querySelector("[name='newComment']");

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
    newPost();
}